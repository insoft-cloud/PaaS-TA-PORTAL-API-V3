package organizationQuotas

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
	myRouter.HandleFunc("/v3/"+uris, createOrganizationQuotas).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getOrganizationQuota).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getOrganizationQuotas).Methods("GET")

}

// Permitted roles "Admin"
// 404 error : Unknown request
func createOrganizationQuotas(w http.ResponseWriter, r *http.Request) {
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
		var final OrganizationQuotas
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Required parameters의 name에 해당한 quota-guid를 찾을 수 없음
func getOrganizationQuota(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final OrganizationQuotas
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// Permitted roles
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
