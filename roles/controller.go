package roles

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
	Valid role types
	organization_user
	organization_auditor
	organization_manager
	organization_billing_manager
	space_auditor
	space_developer
	space_manager
	space_application_supporter - This role is under active development and is not supported
*/

var uris = "roles"

func RoleHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createRole).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getRole).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getRoles).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, deleteRole).Methods("DELETE")
}

//Permitted roles
//Role	Notes
//Admin
//Org Manager	Can create roles in managed organizations and spaces within those organizations;
//				can also create roles for users outside of managed organizations
//				when set_roles_by_username feature_flag is enabled; this requires identifying users by username and origin
//Space Manager	Can create roles in managed spaces for users in their org
func createRole(w http.ResponseWriter, r *http.Request) {
	var pBody CreateRole
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
		var final Role
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles
//Role	Notes
//Admin
//Admin Read-Only
//Global Auditor
//Org Manager	Can see roles in managed organizations or spaces in those organizations
//Org Auditor	Can only see organization roles in audited organizations
//Org Billing Manager	Can only see organization roles in billing-managed organizations
//Space Auditor	Can see roles in audited spaces or parent organizations
//Space Application Supporter (under development)	Can see roles in supported spaces or parent organizations
//Space Developer	Can see roles in developed spaces or parent organizations
//Space Manager	Can see roles in managed spaces or parent organizations
func getRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Role
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

func getRoles(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final RoleList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

func deleteRole(w http.ResponseWriter, r *http.Request) {
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
