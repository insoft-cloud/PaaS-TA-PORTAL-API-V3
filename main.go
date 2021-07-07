package main

import (
	"PAAS-TA-PORTAL-V3/admin"
	"PAAS-TA-PORTAL-V3/app_features"
	"PAAS-TA-PORTAL-V3/app_usage_events"
	"PAAS-TA-PORTAL-V3/apps"
	"PAAS-TA-PORTAL-V3/audit_events"
	"PAAS-TA-PORTAL-V3/buildpacks"
	"PAAS-TA-PORTAL-V3/builds"
	"PAAS-TA-PORTAL-V3/config"
	"PAAS-TA-PORTAL-V3/deployments"
	"PAAS-TA-PORTAL-V3/organization_quotas"
	"PAAS-TA-PORTAL-V3/organizations"
	"PAAS-TA-PORTAL-V3/packages"
	"PAAS-TA-PORTAL-V3/service_brokers"
	_ "fmt"
	"github.com/gorilla/mux"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	admin.AdminHandleRequests(myRouter)
	apps.AppHandleRequests(myRouter)
	app_features.AppFeatureHandleRequests(myRouter)
	app_usage_events.AppUsageEventHandleRequests(myRouter)
	audit_events.AuditEventHandleRequests(myRouter)
	builds.BuildPackHandleRequests(myRouter)
	buildpacks.BuildPackHandleRequests(myRouter)
	deployments.DeploymentHandleRequests(myRouter)
	service_brokers.ServiceBrokerHandleRequests(myRouter)
	organizations.OrganizationsRequests(myRouter)
	organization_quotas.OrganizationQuotasHandleRequests(myRouter)
	packages.PackagesHandleRequests(myRouter)
	log.Fatal(http.ListenAndServe(":"+config.Config["port"], myRouter))
}

func main() {
	config.SetConfig()
	config.ClientSetting()
	config.ValidateConfig()
	handleRequests()
}
