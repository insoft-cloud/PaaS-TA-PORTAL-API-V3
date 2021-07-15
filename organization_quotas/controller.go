package organization_quotas

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "organization_quotas"

func OrganizationQuotasHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createOrganizationQuota).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getOrganizationQuota).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getOrganizationQuotas).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/organizations", applyOrganizationQuota).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteOrganizationQuota).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateOrganizationQuota).Methods("PATCH")
}

// @Description Permitted Roles 'Admin'
// @Summary Create an organization quota
// @Description This endpoint creates a new organization quota, but does not assign it to a specific organization unless an organization GUID is provided in the relationships.organizations parameter.
// @Description To create an organization quota you must be an admin.
// @Tags Organization Quotas
// @Produce json
// @Security ApiKeyAuth
// @Param CreateOrganizationQuotas body CreateOrganizationQuotas true "Create OrganizationQuotas"
// @Success 201 {object} OrganizationQuota
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organization_quotas [POST]
func createOrganizationQuota(w http.ResponseWriter, r *http.Request) {
	var pBody CreateOrganizationQuotas
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
		var final OrganizationQuota
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All Roles'
// @Summary Get an organization quota
// @Description This endpoint gets an individual organization quota resource.
// @Tags Organization Quotas
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "organizations_quotas guid"
// @Success 200 {object} OrganizationQuota
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organization_quotas/{guid} [GET]
func getOrganizationQuota(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final OrganizationQuota
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All Roles'
// @Summary List organization quotas
// @Description This endpoint lists all organization quota resources.
// @Tags Organization Quotas
// @Produce json
// @Security ApiKeyAuth
// @Param guids query []string false "Comma-delimited list of organization quota guids to filter by" collectionFormat(csv)
// @Param names query []string false "Comma-delimited list of organization quota names to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} OrganizationQuota
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organization_quotas [GET]
func getOrganizationQuotas(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final OrganizationQuotasList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin'
// @Summary Apply an organization quota to an organization
// @Description This endpoint applies an organization quota to one or more organizations.
// @Description Only admin users can apply an organization quota to an organization.
// @Tags Organization Quotas
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "organizations_quotas guid"
// @Param ApplyOrganizationQuotas body ApplyOrganizationQuotas true "Apply Organization Quotas"
// @Success 201 {object} OrganizationQuota
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organization_quotas/{guid}/relationships/organizations [POST]
func applyOrganizationQuota(w http.ResponseWriter, r *http.Request) {
	var pBody ApplyOrganizationQuotas
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)

	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/organizations", reqBody, "POST", w, r)
	if rBodyResult {
		var final OrganizationQuota
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin'
// @Summary Delete an organization quota
// @Description Organization quotas cannot be deleted when applied to any organizations.
// @Tags Organization Quotas
// @Produce json
// @Security ApiKeyAuth
// @Success 202 {object} string "ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organization_quotas/{guid} [DELETE]
func deleteOrganizationQuota(w http.ResponseWriter, r *http.Request) {
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

// @Description Permitted roles 'Admin'
// @Summary Update an organization quota
// @Description This endpoint will only update the parameters specified in the request body. Any unspecified parameters will retain their existing values.
// @Tags Organization Quotas
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "organizations_quotas guid"
// @Param UpdateOrganizationQuota body UpdateOrganizationQuota true "Update Organization Quota"
// @Success 200 {object} OrganizationQuota
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /organization_quotas/{guid} [PATCH]
func updateOrganizationQuota(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateOrganizationQuota
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final OrganizationQuota
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
