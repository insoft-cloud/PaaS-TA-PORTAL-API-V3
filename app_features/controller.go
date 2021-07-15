package app_features

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var uris = "apps"

func AppFeatureHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/features/{name}", getAppFeature).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/features", getAppFeatures).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/features/{name}", updateAppFeature).Methods("PATCH")
}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get an app feature
// @Description
// @Tags App Features
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Param name path string true "App Feature Name"
// @Success 200 {object} AppFeature
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/features/{name} [GET]
func getAppFeature(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	name := vars["name"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/features/"+name, nil, "GET", w, r)
	if rBodyResult {
		var final AppFeature
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}

}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary List app features
// @Description This endpoint retrieves the list of features for the specified app.
// @Tags App Features
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} AppFeatureList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/features [GET]
func getAppFeatures(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/features", nil, "GET", w, r)
	if rBodyResult {
		var final AppFeatureList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Space Developer'
// @Summary Update an app feature
// @Description
// @Tags App Features
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Param name path string true "App Name"
// @Param UpdateAppFeature body UpdateAppFeature true "Update App Feature"
// @Success 200 {object} AppFeature
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/features/{name} [PATCH]
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
	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/features/"+name, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final AppFeature
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
