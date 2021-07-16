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
// @Summary Get a space feature
// @Description
// @Tags Space Features
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Space Guid"
// @Param name path string true "Feature Name"
// @Success 200 {object} SpaceFeature
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces/{guid}/features/{name} [GET]
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
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary List space features
// @Description This endpoint retrieves the list of features for the specified space. Currently, the only feature on spaces is the SSH feature.
// @Tags Space Features
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} SpaceFeatureList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces/{guid}/features [GET]
func getSpaceFeatures(w http.ResponseWriter, r *http.Request) {
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/features", nil, "GET", w, r)
	if rBodyResult {
		var final SpaceFeatureList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Update space features
// @Description
// @Tags Space Features
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Space Guid"
// @Param name path string true "Feature Name"
// @Param SpaceFeature body SpaceFeature true "Update SpaceFeature"
// @Success 200 {object} SpaceFeature
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces/{guid}/features/{name} [PATCH]
func updateSpaceFeature(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	name := vars["name"]
	var pBody SpaceFeature
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/features/"+name, nil, "PATCH", w, r)
	if rBodyResult {
		var final SpaceFeature
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
