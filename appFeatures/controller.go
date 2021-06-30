package appFeatures

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func AppFeatureHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/apps/{guid}/features/{name}", getAppFeature).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/features", getAppFeatures).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/features/{name}", updateAppFeature).Methods("PATCH")
}

//Permitted roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
func getAppFeature(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	name := vars["name"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/features/"+name, nil, "GET", w, r)
	if rBodyResult {
		var final AppFeature
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
func getAppFeatures(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/features", nil, "GET", w, r)
	if rBodyResult {
		var final AppFeatureList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin, Space Developer'
func updateAppFeature(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	name := vars["name"]
	var pBody UpdateAppFeature
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/features/"+name, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final AppFeature
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
