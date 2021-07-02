package apps

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

func AppHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/apps", createApp).Methods("POST")
	myRouter.HandleFunc("/v3/apps", getApps).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}", getApp).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}", updateApp).Methods("PATCH")
	myRouter.HandleFunc("/v3/apps/{guid}", deleteApp).Methods("DELETE")
	myRouter.HandleFunc("/v3/apps/{guid}/droplets/current", getAppDroplet).Methods("get")
	myRouter.HandleFunc("/v3/apps/{guid}/relationships/current_droplet", getAppDropletAssociation).Methods("get")
	myRouter.HandleFunc("/v3/apps/{guid}/env", getAppEnv).Methods("get")
	myRouter.HandleFunc("/v3/apps/{guid}/environment_variables", getAppEnvVariables).Methods("get")
	myRouter.HandleFunc("/v3/apps/{guid}/permissions", getAppPermissions).Methods("get")
	myRouter.HandleFunc("/v3/apps/{guid}/relationships/current_droplet", setAppDroplet).Methods("PATCH")
	myRouter.HandleFunc("/v3/apps/{guid}/ssh_enabled", getAppSSH).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/actions/start", startApp).Methods("POST")
	myRouter.HandleFunc("/v3/apps/{guid}/actions/stop", stopApp).Methods("POST")
	myRouter.HandleFunc("/v3/apps/{guid}/actions/restart", restartApp).Methods("POST")
	myRouter.HandleFunc("/v3/apps/{guid}/environment_variables", setAppEnv).Methods("PATCH")
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
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
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
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"?"+query, nil, "GET", w, r)
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
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/apps?"+query, nil, "GET", w, r)
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
func getAppDroplet(w http.ResponseWriter, r *http.Request) {
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

//Permitted roles Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager
func getAppDropletAssociation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/relationships/current_droplet", nil, "GET", w, r)
	if rBodyResult {
		var final AppDropletAssociation
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles Admin Admin Read-Only Space Developer
func getAppEnv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/env", nil, "GET", w, r)
	if rBodyResult {
		var final AppEnv
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles Admin Admin Read-Only Space Developer
func getAppEnvVariables(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/environment_variables", nil, "GET", w, r)
	if rBodyResult {
		var final AppEnvVariable
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager *
func getAppPermissions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/permissions", nil, "GET", w, r)
	if rBodyResult {
		var final AppPermission
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles Admin Space Developer
func setAppDroplet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody AppSetDroplet
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/relationships/current_droplet", reqBody, "PATCH", w, r)
	if rBodyResult {
		var final AppDropletAssociation
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager *
func getAppSSH(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/ssh_enabled", nil, "GET", w, r)
	if rBodyResult {
		var final AppSSH
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles Admin Space Developer
func startApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/actions/start", nil, "POST", w, r)
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
func stopApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/actions/stop", nil, "POST", w, r)
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
func restartApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/actions/restart", nil, "POST", w, r)
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
func setAppEnv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody AppEnvVar
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/environment_variables", reqBody, "PATCH", w, r)
	if rBodyResult {
		var final AppEnvVariable
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
