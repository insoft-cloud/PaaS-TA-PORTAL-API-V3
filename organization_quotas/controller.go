package organization_quotas

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
	quota의 guid 값을 알 수 없고, Unknown request로 인해 전체 테스트 할 수 없음.
*/

var uris = "organization_quotas"

func OrganizationQuotasHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createOrganizationQuota).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getOrganizationQuota).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getOrganizationQuotas).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/organizations", applyOrganizationQuota).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteOrganizationQuota).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateOrganizationQuota).Methods("PATCH")
}

// Permitted Roles "Admin"
// 404 error : Unknown request
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
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Permitted Roles
// "Admin"
// "Admin" Read-Only
// "Global" Auditor
// "Org Manager" Response will only include guids of managed organizations
// "Org Auditor" Response will only include guids of audited organizations
// "Org Billing Manager" Response will only include guids of billing-managed organizations
// "Space Auditor" Response will only include guids of parent organizations
// "Space Developer" Response will only include guids of parent organizations
// "Space Manager" Response will only include guids of parent organizations
// Required parameters의 name에 해당한 quota-guid를 찾을 수 없음
func getOrganizationQuota(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final OrganizationQuota
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Permitted Roles
// "Admin"
// "Admin" Read-Only
// "Global Auditor"
// "Org Manager" Response will only include guids of managed organizations
// "Org Auditor" Response will only include guids of audited organizations
// "Org Billing Manager" Response will only include guids of billing-managed organizations
// "Space Auditor" Response will only include guids of parent organizations
// "Space Developer" Response will only include guids of parent organizations
// "Space Manager" Response will only include guids of parent organizations
// 404 error : Unknown request
func getOrganizationQuotas(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final OrganizationQuotasList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Permitted Roles "Admin"
// quotas에 대한 guid를 찾을 수 없음
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
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Permitted Roles "Admin"
// quotas에 대한 guid를 찾을 수 없음
func deleteOrganizationQuota(w http.ResponseWriter, r *http.Request) {
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
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
