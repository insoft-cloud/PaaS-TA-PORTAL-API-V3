package users

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	_ "net/url"
)

var uris = "users"

func UserHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createUser).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getUser).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getUsers).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, updateUser).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteUser).Methods("DELETE")
}

//Permitted Roles Admin
// @Summary Create a user
// @Description Creating a user requires one value, a GUID. This creates a user in the Cloud Controller database.
// @Tags Users
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateUser body CreateUser true "Create User"
// @Success 200 {object} Users
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /users [POST]
func createUser(w http.ResponseWriter, r *http.Request) {
	var pBody CreateUser
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
		var final Users
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//Permitted Roles All Roles Admin Read-Only Admin Global Auditor Org Auditor Org Billing Manager Org Manager Space Auditor Space Developer Space Manager (Can only view users affiliated with their org)
// @Summary Get a user
// @Description
// @Tags Users
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "User Guid"
// @Success 200 {object} Users
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /users/{guid} [GET]
func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Users
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//Permitted All Roles Admin Read-Only Admin Global Auditor Org Auditor Org Billing Manager Org Manager Space Auditor Space Developer Space Manager (Can only view users affiliated with their org)
// @Summary List users
// @Description
// @Tags Users
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} GetUsers
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /users [GET]
func getUsers(w http.ResponseWriter, r *http.Request) {
	rBody, rBodyResult := config.Curl("/v3/"+uris, nil, "GET", w, r)
	if rBodyResult {
		var final GetUsers
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//Permitted Roles Admin
// @Summary Update a user
// @Description Update a user’s metadata.
// @Tags Users
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "User Guid"
// @Param Updateuser body Updateuser true "Update User"
// @Success 200 {object} Users
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /users/{guid} [PATCH]
func updateUser(w http.ResponseWriter, r *http.Request) {
	var pBody Updateuser
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final Users
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//Permitted Roles Admin
// @Summary Delete a user
// @Description All roles associated with a user will be deleted if the user is deleted.
// @Tags Users
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "User Guid"
// @Success 202 {object} string	"ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /users/{guid} [DELETE]
func deleteUser(w http.ResponseWriter, r *http.Request) {
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
