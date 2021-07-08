package processes

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "processes"

func ProcessHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getProcess).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris+"/{type}", getAppProcess).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/stats", getStatsProcess).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris+"/{type}/stats", getAppStatsProcess).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getProcesses).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris, getAppProcesses).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateProcess).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/actions/scale", scaleProcess).Methods("POST")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris+"/{type}/actions/scale", scaleAppProcess).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/instances/{index}", terminateProcessInstance).Methods("DELETE")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris+"/{type}/instances/{index}", terminateAppProcessInstance).Methods("DELETE")
}

// Permitted roles "Admin", "Admin" Read-Only, "Global Auditor", "Org Manager", "Space Auditor", "Space Developer", "Space Manager"
func getProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Process
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// Permitted roles "Admin", "Admin" Read-Only, "Global Auditor", "Org Manager", "Space Auditor", "Space Developer", "Space Manager"
func getAppProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	appType := vars["type"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"/"+appType, nil, "GET", w, r)
	if rBodyResult {
		var final Process
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// Permitted roles
// Admin, Admin Read-Only, Global Auditor, Org Manager, Space Auditor, Space Developer, Space Manager
func getStatsProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/stats", nil, "GET", w, r)
	if rBodyResult {
		var final ProcessStats
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// Permitted roles
// Admin, Admin Read-Only, Global Auditor, Org Manager, Space Auditor, Space Developer, Space Manager
func getAppStatsProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	appType := vars["type"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"/"+appType+"/stats", nil, "GET", w, r)
	if rBodyResult {
		var final ProcessStats
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//Permitted roles
//All Roles
func getProcesses(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ProcessList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles
//Admin, Admin Read-Only, Global Auditor, Org Manager, Space Auditor, Space Developer, Space Manager
func getAppProcesses(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ProcessList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles
//Admin, Space Developer
func updateProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdateProcess
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final Process
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles Admin, Space Developer
func scaleProcess(w http.ResponseWriter, r *http.Request) {
	var pBody ScaleProcess
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

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/actions/scale", reqBody, "POST", w, r)
	if rBodyResult {
		var final Process
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles Admin, Space Developer
func scaleAppProcess(w http.ResponseWriter, r *http.Request) {
	var pBody ScaleProcess
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
	appType := vars["type"]

	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"/"+appType+"/actions/scale", reqBody, "POST", w, r)
	if rBodyResult {
		var final Process
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted Roles, Admin, Space Developer
func terminateProcessInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	index := vars["index"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/instances/"+index, nil, "DELETE", w, r)
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

func terminateAppProcessInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	index := vars["index"]
	appType := vars["type"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"/"+appType+"/instances/"+index, nil, "DELETE", w, r)
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
