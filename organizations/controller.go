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
}

// Permitted roles "Admin" If the user_org_creation feature flag is enabled, any user with the cloud_controller.write scope will be able to create organizations.
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
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Permitted roles "All Roles"
func getOrganization(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Organizations
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Permitted roles "All Roles"
func getOrganizations(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final OrganizationsList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Permitted roles "Admin", "Admin Read-Only", "Global Auditor", "Org Auditor", "Org Billing Manager", "Org Manager"
// 404 error
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
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
