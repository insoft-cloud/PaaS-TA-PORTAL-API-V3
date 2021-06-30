package appUsageEvents

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

func AppFeatureHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/audit_events/{guid}", getAuditEvent).Methods("GET")
	myRouter.HandleFunc("/v3/audit_events", getAuditEvents).Methods("GET")
}

//Permitted roles 'Admin Admin Read-Only Global Auditor Space Auditor Space Developer Org Auditor'
func getAuditEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/audit_events/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final AuditEvent
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin Read-Only Admin Global Auditor Org Auditor Org Manager Space Auditor Space Developer Space Manager'
func getAuditEvents(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/app_usage_events?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final AuditEventList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
