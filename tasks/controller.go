package tasks

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "tasks"

func TaskHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createTask).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getTask).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getTasks).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris, getAppTask).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateTask).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/actions/cancel", cancelTask).Methods("POST")

}

// @Description Permitted Roles Admin Space Developer
// @Summary Create a task
// @Description
// @Tags Tasks
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateTask body CreateTask true "Create Task"
// @Success 200 {object} Tasks
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /tasks [POST]
func createTask(w http.ResponseWriter, r *http.Request) {
	var pBody CreateTask
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
		var final Tasks
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Admin Read-Only Admin Global Auditor Org Manager Space Auditor Space Developer Space Manage
// @Summary Get a task
// @Description
// @Tags Tasks
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Task Guid"
// @Success 200 {object} Tasks
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /tasks/{guid} [GET]
func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Tasks
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary List tasks
// @Description
// @Tags Tasks
// @Produce  json
// @Security ApiKeyAuth
// @Param query query string false "query"
// @Success 200 {object} AppTask
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /tasks [GET]
func getTasks(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final AppTask
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles All Roles
// @Summary List tasks for an app
// @Description
// @Tags Tasks
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Task Guid"
// @Success 200 {object} AppTask
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/tasks [GET]
func getAppTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris, nil, "GET", w, r)
	if rBodyResult {
		var final AppTask
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin Space Developer
// @Summary Update a task
// @Description
// @Tags Tasks
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Task Guid"
// @Param UpdateTask body UpdateTask true "Update Task"
// @Success 200 {object} Tasks
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /tasks/{guid} [PATCH]
func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateTask
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
		var final Tasks
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin Space Developer
// @Summary Cancel a task
// @Description Cancels a running task.
// @Tags Tasks
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Tasks Guid"
// @Param Tasks body Tasks true "Cancel Tasks"
// @Success 200 {object} Tasks
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /tasks/{guid}/actions/cancel [POST]
func cancelTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody Tasks
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/actions/cancel", nil, "POST", w, r)
	if rBodyResult {
		var final Tasks
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}

}
