package main

import (
	"PAAS-TA-PORTAL-V3/admin"
	"PAAS-TA-PORTAL-V3/app"
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
	app.AppHandleRequests(myRouter)
	log.Fatal(http.ListenAndServe(":"+config.Config["port"], myRouter))
}

func main() {
	config.SetConfig()
	config.ClientSetting()
	config.ValidateConfig()
	handleRequests()
}
