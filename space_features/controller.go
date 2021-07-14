package space_features

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var uris = "spaces"

func SpaceFeatureHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/features/{name}", getSpaceFeature).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/features", getSpaceFeatures).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/features/{name}", updateSpaceFeature).Methods("PATCH")
}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
func getSpaceFeature(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	name := vars["name"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/features/"+name, nil, "GET", w, r)
	if rBodyResult {
		var final SpaceFeature
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
func getSpaceFeatures(w http.ResponseWriter, r *http.Request) {
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/features", nil, "GET", w, r)
	if rBodyResult {
		var final SpaceFeatureList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
func updateSpaceFeature(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	name := vars["name"]
	var pBody interface{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/features/"+name, nil, "PATCH", w, r)
	if rBodyResult {
		var final SpaceFeature
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
