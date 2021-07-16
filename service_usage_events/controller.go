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
// @Summary Get a service usage event
// @Description Retrieve a service usage event.
// @Tags Service Usage Events
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceUsageEvent Guid"
// @Success 200 {object} ServiceUsageEvent
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_usage_events/{guid} [GET]
func getServiceUsageEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceUsageEvent
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary List service usage events
// @Description Retrieve all service usage events the user has access to.
// @Tags Service Usage Events
// @Produce  json
// @Security ApiKeyAuth
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param after_guid query string false "Filters out events before and including the event with the given guid"
// @Param guids query []string false "Comma-delimited list of app guids to filter by" collectionFormat(csv)
// @Param service_instance_types query []string false "Comma-delimited list of service instance types to filter by; valid values are managed_service_instance and user_provided_service_instance" collectionFormat(csv)
// @Param service_offering_guids query []string false "Comma-delimited list of service offering guids to filter by" collectionFormat(csv)
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} ServiceUsageEventList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_usage_events [GET]
func getServiceUsageEvents(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceUsageEventList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin'
// @Summary Purge and seed service usage events
// @Description Destroys all existing events. Populates new usage events, one for each existing service instance. All populated events will have a created_at value of current time. There is the potential race condition if service instances are currently being created or deleted. The seeded usage events will have the same guid as the service instance.
// @Tags Service Usage Events
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} string "ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_usage_events/actions/destructively_purge_all_and_reseed [POST]
func purgeAndSeedServiceUsageEvent(w http.ResponseWriter, r *http.Request) {
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/actions/destructively_purge_all_and_reseed", nil, "POST", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
