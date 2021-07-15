package organizations

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "organizations"

func OrganizationsRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createOrganizations).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getOrganization).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getOrganizations).Methods("GET")
	myRouter.HandleFunc("/v3/isolation_segments/{guid}/"+uris, getOrganizationsIsolationSegment).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateOrganizations).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteOrganizations).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/default_isolation_segment", assignDefaultIsolationSegmentOrganizations).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/default_isolation_segment", getDefaultIsolationSegment).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/domains/default", getDefaultDomain).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/usage_summary", getUsageSummary).Methods("GET")
}

// @Description Permitted Roles 'Admin, If the user_org_creation feature flag is enabled, any user with the cloud_controller.write scope will be able to create organizations.'
// @Summary Create an organization
// @Description
// @Tags Organizations
// @Produce json
// @Security ApiKeyAuth
// @Param CreateOrganizations body CreateOrganizations true "Create Organizations"
// @Success 200 {object} Organizations
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organizations [POST]
func createOrganizations(w http.ResponseWriter, r *http.Request) {
	var pBody CreateOrganizations
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris, reqBody, "POST", w, r)
	if rBodyResult {
		var final Organizations
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted 'All Roles'
// @Summary Get an organization
// @Description Retrieve all organizations the user has access to.
// @Tags Organizations
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "organization Guid"
// @Success 200 {object} Organizations
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organizations/{guid} [GET]
func getOrganization(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Organizations
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted 'All Roles'
// @Summary List organizations
// @Description Retrieve all organizations the user has access to.
// @Tags Organizations
// @Produce  json
// @Security ApiKeyAuth
// @Param guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param names query []string false "Comma-delimited list of organization names to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param lifecycle_type query string false "Lifecycle type to filter by; valid values are buildpack, docker"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} Organizations
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organizations [GET]
func getOrganizations(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final OrganizationsList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Admin Read-Only, Global Auditor, Org Auditor, Org Billing Manager, Org Manager'
// @Summary List organizations for isolation segment
// @Description Retrieve the organizations entitled to the isolation segment. Return only the organizations the user has access to.
// @Tags Organizations
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "isolation_segment Guid"
// @Param guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param names query []string false "Comma-delimited list of organization names to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Success 200 {object} Organizations
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /isolation_segments/{guid}/organizations [GET]
func getOrganizationsIsolationSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/isolation_segments/"+guid+"/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final OrganizationsList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Org Manager'
// @Summary Update an organization
// @Description
// @Tags Organizations
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "organization Guid"
// @Param UpdateOrganizations body UpdateOrganizations true "Update Organizations"
// @Success 200 {object} Organizations
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organizations/{guid} [PATCH]
func updateOrganizations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdateOrganizations
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final Organizations
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin'
// @Summary Delete an organization
// @Description When an organization is deleted, user roles associated with the organization will also be deleted.
// @Tags Organizations
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "organization Guid"
// @Success 202 {object} string "ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organizations/{guid} [DELETE]
func deleteOrganizations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "DELETE", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Org Manager'
// @Summary Assign default isolation segment
// @Description Set the default isolation segment for a given organization. Only isolation segments that are entitled to the organization are eligible to be the default isolation segment.
// @Description Apps will not run in the new default isolation segment until they are restarted.
// @Tags Organizations
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "organization Guid"
// @Param DefaultIsolationSegmentOrganizations body DefaultIsolationSegmentOrganizations true "Isolation segment relationship; apps will run in this isolation segment; set data to null to remove the relationship"
// @Success 200 {object} Organizations
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organizations/{guid}/relationships/default_isolation_segment [PATCH]
func assignDefaultIsolationSegmentOrganizations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody DefaultIsolationSegmentOrganizations
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/default_isolation_segment", reqBody, "PATCH", w, r)
	if rBodyResult {
		var final Organizations
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted 'All Roles'
// @Summary Get default isolation segment
// @Description Retrieve the default isolation segment for a given organization.
// @Tags Organizations
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "organization Guid"
// @Success 200 {object} Organizations
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organizations/{guid}/relationships/default_isolation_segment [GET]
func getDefaultIsolationSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/default_isolation_segment", nil, "GET", w, r)
	if rBodyResult {
		var final DefaultIsolationSegmentOrganizations
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Space Developer Space Manager Space Auditor Org Auditor Org Manager Org Billing Manager Can only view domains without an organization relationship Admin Admin Read-Only Global Auditor'
// @Summary Get default domain
// @Description Retrieve the default domain for a given organization.
// @Tags Organizations
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "organization Guid"
// @Success 200 {object} Organizations
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organizations/{guid}/domains/default [GET]
func getDefaultDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/domains/default", nil, "GET", w, r)
	if rBodyResult {
		var final GetDefaultDomain
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted 'All Roles'
// @Summary Get usage summary
// @Description This endpoint retrieves the specified organization object’s memory and app instance usage summary.
// @Tags Organizations
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "organization Guid"
// @Success 200 {object} Organizations
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organizations/{guid}/usage_summary [GET]
func getUsageSummary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/usage_summary", nil, "GET", w, r)
	if rBodyResult {
		var final GetUsageSummary
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
