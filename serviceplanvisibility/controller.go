package serviceplanvisibility

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var uris = "service_plans"

func ServicePlanVisibilityHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/visibility", getServicePlanVisibility).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/visibility", updateServicePlanVisibility).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/visibility", applyServicePlanVisibility).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/visibility/{organization_guid]", deleteServiceBroker).Methods("DELETE")
}

//Permitted All Roles
func getServicePlanVisibility(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/visibility", nil, "GET", w, r)
	if rBodyResult {
		var final ServicePlanVisibility
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted All Roles
func updateServicePlanVisibility(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateServicePlanVisibility
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/visibility", reqBody, "PATCH", w, r)
	if rBodyResult {
		var final ServicePlanVisibility
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted All Roles
func applyServicePlanVisibility(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateServicePlanVisibility
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/visibility", reqBody, "POST", w, r)
	if rBodyResult {
		var final ServicePlanVisibility
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin Space Developer (only space-scoped brokers)'
func deleteServiceBroker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	org_guid := vars["organization_guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/visibility/"+org_guid, nil, "DELETE", w, r)
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
