package stacks

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "stacks"

func AppHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createStack).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getStack).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getStacks).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateStack).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteStack).Methods("DELETE")
}

// @Description Permitted Roles Admin
// @Summary Create a stack
// @Description
// @Tags Stacks
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateStack body CreateStack true "Create Stack"
// @Success 200 {object} Stack
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /stacks [POST]
func createStack(w http.ResponseWriter, r *http.Request) {
	var pBody CreateStack
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
		var final Stack
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles All Roles
// @Summary Get a stack
// @Description
// @Tags Stacks
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Stack Guid"
// @Success 200 {object} Stack
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /stacks/{guid} [GET]
func getStack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Stack
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary List stacks
// @Description Retrieve all stacks.
// @Tags Stacks
// @Produce  json
// @Security ApiKeyAuth
// @Param query query string false "query"
// @Success 200 {object} Stacks
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /stacks [GET]
func getStacks(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final Stacks
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin Space Developer
// @Summary Update a stack
// @Description
// @Tags Stacks
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Stack Guid"
// @Param UpdateStack body UpdateStack true "Update Stack"
// @Success 200 {object} Stack
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /stacks/{guid} [PATCH]
func updateStack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateStack
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
		var final Stack
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin
// @Summary Delete a stack
// @Description
// @Tags Stacks
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Stack Guid"
// @Success 202 {object} string	"ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /stacks/{guid} [DELETE]
func deleteStack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "DELETE", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
