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
	"PAAS-TA-PORTAL-V3/domains"
	"PAAS-TA-PORTAL-V3/droplets"
	"PAAS-TA-PORTAL-V3/environment_variable_groups"
	"PAAS-TA-PORTAL-V3/feature_flags"
	"PAAS-TA-PORTAL-V3/info"
	"PAAS-TA-PORTAL-V3/isolation_segment"
	"PAAS-TA-PORTAL-V3/jobs"
	"PAAS-TA-PORTAL-V3/organization_quotas"
	"PAAS-TA-PORTAL-V3/organizations"
	"PAAS-TA-PORTAL-V3/packages"
	"PAAS-TA-PORTAL-V3/processes"
	"PAAS-TA-PORTAL-V3/service_brokers"
	"PAAS-TA-PORTAL-V3/stacks"
	"PAAS-TA-PORTAL-V3/tasks"
	"PAAS-TA-PORTAL-V3/users"
	_ "fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name cf-Authorization
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
	domains.DomainHandleRequests(myRouter)
	droplets.DropletHandleRequests(myRouter)
	organization_quotas.OrganizationQuotasHandleRequests(myRouter)
	packages.PackagesHandleRequests(myRouter)
	environment_variable_groups.EnvironmentVariableGroupsHandleRequests(myRouter)
	feature_flags.FeatureFlagHandleRequests(myRouter)
	info.InforHandleRequests(myRouter)
	isolation_segments.IsolationSegmentsHandleRequests(myRouter)
	jobs.JobsHandleRequests(myRouter)
	stacks.AppHandleRequests(myRouter)
	tasks.TaskHandleRequests(myRouter)
	users.UserHandleRequests(myRouter)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type", "cf-Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodTrace, http.MethodOptions})

	log.Fatal(http.ListenAndServe(":"+config.Config["port"], handlers.CORS(originsOk, headersOk, methodsOk)(myRouter)))

	processes.ProcessHandleRequests(myRouter)
	log.Fatal(http.ListenAndServe(":"+config.Config["port"], myRouter))
}

func main() {
	config.SetConfig()
	config.ClientSetting()
	config.ValidateConfig()
	handleRequests()
}
