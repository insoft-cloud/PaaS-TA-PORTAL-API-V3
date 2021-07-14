package audit_events

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

var uris = "audit_events"

func AuditEventHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getAuditEvent).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getAuditEvents).Methods("GET")
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Space Auditor Space Developer Org Auditor'
// @Summary Get an audit event
// @Description
// @Tags Audit Events
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Audit Event Guid"
// @Success 200 {object} AuditEvent
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /audit_events/{guid} [GET]
func getAuditEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final AuditEvent
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Read-Only Admin Global Auditor Org Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary List audit events
// @Description Retrieve all audit events the user has access to.
// @Tags Audit Events
// @Produce  json
// @Security ApiKeyAuth
// @Param types query []string false "Comma-delimited list of event types to filter by" collectionFormat(csv)
// @Param target_guids query []string false "Comma-delimited list of target guids to filter by. Also supports filtering by exclusion." collectionFormat(csv)
// @Param space_guids query []string false "Comma-delimited list of space guids to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} AuditEventList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /audit_events [GET]
func getAuditEvents(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final AuditEventList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
