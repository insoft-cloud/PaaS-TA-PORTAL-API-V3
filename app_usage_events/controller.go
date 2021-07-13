package app_usage_events

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

var uris = "app_usage_events"

func AppUsageEventHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getAppUsageEvent).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"", getAppUsageEvents).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/actions/destructively_purge_all_and_reseed", purgeSeedAppUsageEvents).Methods("POST")
}

//Permitted Roles 'Admin Admin Read-Only Global Auditor'
// @Summary Get an app
// @Description
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Param include query string false "Optionally include additional related resources in the response; valid values are space and space.organization"
// @Success 200 {object} App
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router  /v3/apps/{guid} [GET]
func getAppUsageEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final AppUsageEvent
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//Permitted All Roles
func getAppUsageEvents(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final AppUsageEventList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//Permitted Roles 'Admin'
func purgeSeedAppUsageEvents(w http.ResponseWriter, r *http.Request) {
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/actions/destructively_purge_all_and_reseed", nil, "POST", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
