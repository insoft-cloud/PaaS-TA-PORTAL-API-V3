package builds

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "builds"

func BuildPackHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createBuild).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getBuild).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getBuilds).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris, getBuildApps).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateBuild).Methods("PATCH")
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Create a build
// @Description
// @Tags Builds
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateBuild body CreateBuild true "Create Build"
// @Success 200 {object} Build
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /builds [POST]
func createBuild(w http.ResponseWriter, r *http.Request) {
	var pBody CreateBuild
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)

	rBody, rBodyResult := config.Curl("/v3/"+uris, reqBody, "POST", w, r)
	if rBodyResult {
		var final Build
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Space Auditor Space Developer Org Auditor'
// @Summary Get a build
// @Description
// @Tags Builds
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Build Guid"
// @Success 200 {object} Build
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /builds/{guid} [GET]
func getBuild(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Build
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary List builds
// @Description Retrieve all builds the user has access to.
// @Tags Builds
// @Produce  json
// @Security ApiKeyAuth
// @Param states query []string false "Comma-delimited list of build states to filter by" collectionFormat(csv)
// @Param app_guids query []string false "Comma-delimited list of app guids to filter by" collectionFormat(csv)
// @Param package_guids query []string false "Comma-delimited list of package guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} BuildList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /builds [GET]
func getBuilds(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final BuildList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Space Auditor Space Developer Org Auditor'
// @Summary List builds
// @Description Retrieve all builds for the app.
// @Tags Builds
// @Produce  json
// @Security ApiKeyAuth
// @Param states query []string false "Comma-delimited list of build states to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} BuildList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/builds [GET]
func getBuildApps(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final BuildList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer Build State Updater'
// @Summary Update a build
// @Description
// @Tags Builds
// @Produce  json
// @Security ApiKeyAuth
// @Param UpdateBuild body UpdateBuild true "Update Build"
// @Success 200 {object} Build
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /builds/{guid} [PATCH]
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
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final Build
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
