package app

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func AppHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/apps", createApp).Methods("POST")
	myRouter.HandleFunc("/v3/apps", getApps).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}", getApp).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}", updateApp).Methods("PATCH")
	myRouter.HandleFunc("/v3/apps/{guid}", deleteApp).Methods("DELETE")
	myRouter.HandleFunc("/v3/apps/{guid}/droplets/current", getDroplet).Methods("get")
}

//Permitted roles 'Admin, SpaceDeveloper'
func createApp(w http.ResponseWriter, r *http.Request) {
	var pBody CreateApp
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.Curl("/v3/apps", reqBody, "POST", w, r)
	if rBodyResult {
		var final App
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
func getApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	//req, _ := http.NewRequest("GET", config.GetDomainConfig() +"/v3/apps/" + guid, nil)
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final App
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted All Roles
func getApps(w http.ResponseWriter, r *http.Request) {
	rBody, rBodyResult := config.Curl("/v3/apps", nil, "GET", w, r)
	if rBodyResult {
		var final AppList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles Admin Space Developer
func updateApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdateApp
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final App
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles Admin Space Developer
func deleteApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid, nil, "DELETE", w, r)
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

//Permitted roles Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager
func getDroplet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/droplets/current", nil, "GET", w, r)
	if rBodyResult {
		var final AppDroplet
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
