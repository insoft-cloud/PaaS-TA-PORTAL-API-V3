package main

import (
	"PAAS-TA-PORTAL-V3/admin"
	"PAAS-TA-PORTAL-V3/appFeatures"
	"PAAS-TA-PORTAL-V3/appUsageEvents"
	"PAAS-TA-PORTAL-V3/apps"
	"PAAS-TA-PORTAL-V3/auditEvents"
	"PAAS-TA-PORTAL-V3/buildpacks"
	"PAAS-TA-PORTAL-V3/builds"
	"PAAS-TA-PORTAL-V3/config"
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
	appFeatures.AppFeatureHandleRequests(myRouter)
	appUsageEvents.AppUsageEventHandleRequests(myRouter)
	auditEvents.AuditEventHandleRequests(myRouter)
	builds.BuildPackHandleRequests(myRouter)
	buildpacks.BuildPackHandleRequests(myRouter)
	log.Fatal(http.ListenAndServe(":"+config.Config["port"], myRouter))
}

func main() {
	config.SetConfig()
	config.ClientSetting()
	config.ValidateConfig()
	handleRequests()
}
