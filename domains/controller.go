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

//Permitted Roles 'Org Manager'
// @Summary Create a domain
// @Description
// @Tags Domains
// @Produce  json
// @Security ApiKeyAuth
// @Param name path string true "name"
// @Param internal path boolean true "false"
// @Success 200 {object} CreateDomain
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

//Permitted Roles Admin Read-Only Admin Global Auditor Org Auditor Org Billing Manager Can only view domains without an organization relationship Org Manager Space Auditor Space Developer Space Manager
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

//Permitted All Roles
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

//Permitted All Roles
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

//Permitted Roles 'Admin, Org Manager'
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

//Permitted Roles Admin Org Manager
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

//Permitted Roles Admin Org Manager
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

//Permitted Roles 'Org Manager'
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
