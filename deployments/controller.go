package deployments

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

func DeploymentHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/deployments", createDeployment).Methods("POST")
	myRouter.HandleFunc("/v3/deployments/{guid}", getDeployment).Methods("GET")
	myRouter.HandleFunc("/v3/deployments", getDeployments).Methods("GET")
	myRouter.HandleFunc("/v3/deployments/{guid}", updateDeployment).Methods("PATCH")
	myRouter.HandleFunc("/v3/deployments/{guid}", cancelDeployment).Methods("DELETE")
}

//Permitted roles 'Admin Space Developer'
func createDeployment(w http.ResponseWriter, r *http.Request) {
	var pBody CreateDeployment
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)

	rBody, rBodyResult := config.Curl("/v3/deployments", reqBody, "POST", w, r)
	if rBodyResult {
		var final Deployment
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
func getDeployment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/deployments/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Deployment
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin Read-Only Admin Global Auditor Org Auditor Org Manager Space Auditor Space Developer Space Manager'
func getDeployments(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/deployments?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final DeploymentList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin Space Developer'
func updateDeployment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateDeployment
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)

	rBody, rBodyResult := config.Curl("/v3/deployments/"+guid, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final Deployment
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin Space Developer'
func cancelDeployment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/deployments/"+guid, nil, "Delete", w, r)
	if rBodyResult {
		var final Deployment
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
