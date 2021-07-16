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

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor'
// @Summary Get an app usage event
// @Description Retrieve an app usage event.
// @Tags App Usage Events
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} AppUsageEvent
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /app_usage_events/{guid} [GET]
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

// @Description Permitted Roles 'All'
// @Summary List app usage events
// @Description Retrieve an app usage event.
// @Tags App Usage Events
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page;\n valid values are 1 through 5000"
// @Param order_by query string false "	Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid value is created_at"
// @Param after_guid query string false "Filters out events before and including the event with the given guid"
// @Param guids query string false "Comma-delimited list of usage event guids to filter by"
// @Param created_ats  query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} AppUsageEvent
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /app_usage_events/{guid} [GET]
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

// @Description Permitted Roles 'Admin'
// @Summary Purge and seed app usage events
// @Description Destroys all existing events. Populates new usage events, one for each started app. All populated events will have a created_at value of current time. There is the potential race condition if apps are currently being started, stopped, or scaled. The seeded usage events will have the same guid as the app.
// @Tags App Usage Events
// @Produce  json
// @Security ApiKeyAuth
// @Success 200
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /app_usage_events/actions/destructively_purge_all_and_reseed [POST]
func purgeSeedAppUsageEvents(w http.ResponseWriter, r *http.Request) {
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/actions/destructively_purge_all_and_reseed", nil, "POST", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
