package buildpacks

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "buildpacks"

func BuildPackHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createBuildPack).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getBuildPack).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getBuildPacks).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateBuildPack).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteBuildPack).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/upload", uploadBuildPack).Methods("POST")
}

// @Description Permitted Roles 'Admin'
// @Summary Create a buildpack
// @Description
// @Tags Buildpacks
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateBuildPack body CreateBuildPack true "Create BuildPack"
// @Success 200 {object} BuildPack
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /buildpacks [POST]
func createBuildPack(w http.ResponseWriter, r *http.Request) {
	var pBody CreateBuildPack
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
		var final BuildPack
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary Get a buildpack
// @Description
// @Tags Buildpacks
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Buildpacks Guid"
// @Success 200 {object} BuildPack
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router  /v3/buildpacks/{guid} [GET]
func getBuildPack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final BuildPack
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary List buildpacks
// @Description Retrieve all buildpacks the user has access to.
// @Tags Buildpacks
// @Produce  json
// @Security ApiKeyAuth
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param names query []string false "Comma-delimited list of buildpack names to filter by" collectionFormat(csv)
// @Param stacks query []string false "Comma-delimited list of stack names to filter by" collectionFormat(csv)
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} BuildPackList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /buildpacks [GET]
func getBuildPacks(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final BuildPackList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin'
// @Summary Update a buildpack
// @Description
// @Tags Buildpacks
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "BuildPack Guid"
// @Param UpdateBuildPack body UpdateBuildPack true "Update BuildPack"
// @Success 200 {object} BuildPack
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /buildpacks/{guid} [PATCH]
func updateBuildPack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdateBuildPack
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final BuildPack
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin'
// @Summary Delete a buildpack
// @Description
// @Tags Buildpacks
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "BuildPack Guid"
// @Success 200 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /buildpacks/{guid} [DELETE]
func deleteBuildPack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "PATCH", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin'
// @Summary Update a buildpack
// @Description
// @Tags Buildpacks
// @Accept mpfd
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "BuildPack Guid"
// @Param bits formData file true "A binary zip file containing the buildpack bits"
// @Success 200 {object} BuildPack
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /buildpacks/{guid}/upload [POST]
func uploadBuildPack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.FileCurl("bits", "/v3/"+uris+"/"+guid+"/upload", "POST", w, r)
	if rBodyResult {
		var final BuildPack
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
