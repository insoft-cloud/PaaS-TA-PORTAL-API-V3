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

// Permitted Roles "Admin" If the user_org_creation feature flag is enabled, any user with the cloud_controller.write scope will be able to create organizations.
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

// Permitted All Roles
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

// Permitted All Roles
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

// Permitted Roles "Admin", "Admin Read-Only", "Global Auditor", "Org Auditor", "Org Billing Manager", "Org Manager"
// 404 error: Isolation segment not found
// 진행 오래걸릴것 같은 부분 pass
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

// Permitted Roles "Admin", "Org Manager"
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
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Permitted Roles "Admin"
// Unknown request
func deleteOrganizations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "DELETE", w, r)
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

//Permitted Roles "Admin", "Org Manager"
// 진행 오래걸릴것 같은 부분 pass
// iso-seg guid 확인해야됨.
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
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Permitted All Roles
// 진행안되는 부분 pass
func getDefaultIsolationSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/default_isolation_segment", nil, "GET", w, r)
	if rBodyResult {
		var final DefaultIsolationSegmentOrganizations
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Permitted Roles "Space Developer", "Space Manager", "Space Auditor", "Org Auditor", "Org Manager"
// "Org Billing Manager" Can only view domains without an organization relationship
//  "Admin", "Admin" Read-Only, "Global Auditor"
func getDefaultDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/domains/default", nil, "GET", w, r)
	if rBodyResult {
		var final GetDefaultDomain
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Permitted All Roles
// Unknown request
// 진행안되는 부분 pass
func getUsageSummary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/usage_summary", nil, "GET", w, r)
	if rBodyResult {
		var final GetUsageSummary
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
