package domains

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "domains"

func DomainHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createDomain).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getDomain).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getDomains).Methods("GET")
	myRouter.HandleFunc("/v3/organizations/{guid}/"+uris, getDomainsOrganization).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateDomains).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteDomains).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/shared_organizations", shareDomains).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/shared_organizations/{org_guid}", unShareDomains).Methods("DELETE")

}

// @Description Permitted Roles 'Admin Org Manager'
// @Summary Create a domain
// @Description
// @Tags Domains
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateDomain body CreateDomain true "Create Domain"
// @Success 200 {object} Domain
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /domains [POST]
func createDomain(w http.ResponseWriter, r *http.Request) {
	var pBody CreateDomain
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
		var final Domain
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin Read-Only Admin Global Auditor Org Auditor Org Billing Manager Can only view domains without an organization relationship Org Manager Space Auditor Space Developer Space Manager
// @Summary Get a domain
// @Description
// @Tags Domains
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Domain Guid"
// @Success 200 {object} Domain
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /domains/{guid} [GET]
func getDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Domain
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary List domains
// @Description Retrieve all domains the user has access to.
// @Tags Domains
// @Produce  json
// @Security ApiKeyAuth
// @Param guids query []string false "Comma-delimited list of guids to filter by" collectionFormat(csv)
// @Param names query []string false "Comma-delimited list of domain  names to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} DomainList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /domains [GET]
func getDomains(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final DomainList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary List domains for an organization
// @Description Retrieve all domains available in an organization for the current user. This will return unscoped domains (those without an owning organization), domains that are scoped to the given organization (owned by the given organization), and domains that have been shared with the organization.
// @Tags Domains
// @Produce  json
// @Security ApiKeyAuth
// @Param guids query []string false "Comma-delimited list of guids to filter by" collectionFormat(csv)
// @Param names query []string false "Comma-delimited list of domain  names to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param guid path string true "Organization Guid"
// @Success 200 {object} OrganizationDomainsList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organizations/{guid}/domains [GET]
func getDomainsOrganization(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/organizations/"+guid+"/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult { // /v3/organizations/{guid}/domains
		var final OrganizationDomainsList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}

}

// @Description Permitted Roles 'Admin, Org Manager'
// @Summary Update a domain
// @Description
// @Tags Domains
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Domain Guid"
// @Param UpdateDomains body UpdateDomains true "Update Domains"
// @Success 200 {object} Domain
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /domains/{guid} [PATCH]
func updateDomains(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateDomains
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
		var final Domain
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin Org Manager
// @Summary Delete a domain
// @Description
// @Tags Domains
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Domain Guid"
// @Success 202 {object} string	"ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /domains/{guid} [DELETE]
func deleteDomains(w http.ResponseWriter, r *http.Request) {
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

// @Description Permitted Roles Admin Org Manager
// @Summary Share a domain
// @Description This endpoint shares an organization-scoped domain to other organizations specified by a list of organization guids. This will allow any of the other organizations to use the organization-scoped domain.
// @Tags Domains
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Domain Guid"
// @Param ShareDomains body ShareDomains true "Share Domains"
// @Success 200 {object} Domain
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /domains/{guid}/relationships/shared_organizations [POST]
func shareDomains(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody ShareDomains
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/shared_organizations", reqBody, "POST", w, r)
	if rBodyResult {
		var final Domain
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin Org Manager
// @Summary Unshare a domain
// @Description This endpoint removes an organization from the list of organizations an organization-scoped domain is shared with. This prevents the organization from using the organization-scoped domain.
// @Tags Domains
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Domain Guid"
// @Param org_guid path string true "Organization Guid"
// @Success 200 {object} Domain
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /domains/{guid}/relationships/shared_organizations/{org_guid} [DELETE]
func unShareDomains(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	org_guid := vars["org_guid"]

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/shared_organizations/"+org_guid, nil, "DELETE", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
