package builds

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

func BuildPackHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/builds", createBuild).Methods("POST")
	myRouter.HandleFunc("/v3/builds/{guid}", getBuild).Methods("GET")
	myRouter.HandleFunc("/v3/builds", getBuilds).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/builds", getBuildApps).Methods("GET")
	myRouter.HandleFunc("/v3/builds/{guid}", updateBuild).Methods("PATCH")
}

//Permitted roles 'Admin Space Developer'
func createBuild(w http.ResponseWriter, r *http.Request) {
	var pBody CreateBuild
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.Curl("/v3/builds", reqBody, "POST", w, r)
	if rBodyResult {
		var final Build
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin Admin Read-Only Global Auditor Space Auditor Space Developer Org Auditor'
func getBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/builds/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Build
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted All Roles
func getBuilds(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/builds?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final BuildList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin Admin Read-Only Global Auditor Space Auditor Space Developer Org Auditor'
func getBuildApps(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/builds?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final BuildList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin Space Developer Build State Updater'
func updateBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdateBuild
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.Curl("/v3/builds/"+guid, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final Build
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
