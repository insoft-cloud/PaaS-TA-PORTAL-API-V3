package service_usage_events

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

var uris = "service_usage_events"

func ServiceUsageEventHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getServiceUsageEvent).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getServiceUsageEvents).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/actions/destructively_purge_all_and_reseed", purgeAndSeedServiceUsageEvent).Methods("POST")
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor'
func getServiceUsageEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceUsageEvent
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted Roles 'All'
func getServiceUsageEvents(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceUsageEventList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted Roles 'Admin'
func purgeAndSeedServiceUsageEvent(w http.ResponseWriter, r *http.Request) {
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
