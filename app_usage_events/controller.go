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

//Permitted roles 'Admin Admin Read-Only Global Auditor'
func getAppUsageEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final AppUsageEvent
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'All Roles'
func getAppUsageEvents(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final AppUsageEventList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin'
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
