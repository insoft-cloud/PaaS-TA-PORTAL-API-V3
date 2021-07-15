package processes

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "processes"

func ProcessHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getProcess).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris+"/{type}", getAppProcess).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/stats", getStatsProcess).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris+"/{type}/stats", getAppStatsProcess).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getProcesses).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris, getAppProcesses).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateProcess).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/actions/scale", scaleProcess).Methods("POST")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris+"/{type}/actions/scale", scaleAppProcess).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/instances/{index}", terminateProcessInstance).Methods("DELETE")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris+"/{type}/instances/{index}", terminateAppProcessInstance).Methods("DELETE")
}

//  @Description Permitted roles "Admin", "Admin" Read-Only, "Global Auditor", "Org Manager", "Space Auditor", "Space Developer", "Space Manager"
// @Summary Get a process
// @Description
// @Tags Processes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "app guid"
// @Success 200 {object} Process
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /processes/{guid} [GET]
func getProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Process
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//  @Description Permitted roles "Admin", "Admin" Read-Only, "Global Auditor", "Org Manager", "Space Auditor", "Space Developer", "Space Manager"
// @Summary Get a process
// @Description
// @Tags Processes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "app guid"
// @Param type path string true "process type"
// @Success 200 {object} Process
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/processes/{type} [GET]
func getAppProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	appType := vars["type"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"/"+appType, nil, "GET", w, r)
	if rBodyResult {
		var final Process
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//  @Description Permitted roles
// Admin, Admin Read-Only, Global Auditor, Org Manager, Space Auditor, Space Developer, Space Manager
// @Summary Get stats for a process
// @Description
// @Tags Processes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "process guid"
// @Success 200 {object} Process
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /processes/{guid}/stats [GET]
func getStatsProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/stats", nil, "GET", w, r)
	if rBodyResult {
		var final ProcessStats
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//  @Description Permitted roles
// Admin, Admin Read-Only, Global Auditor, Org Manager, Space Auditor, Space Developer, Space Manager
// @Summary Get stats for a process
// @Description
// @Tags Processes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "app guid"
// @Param type path string true "process type"
// @Success 200 {object} Process
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/processes/{type}/stats [GET]
func getAppStatsProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	appType := vars["type"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"/"+appType+"/stats", nil, "GET", w, r)
	if rBodyResult {
		var final ProcessStats
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted roles
//All Roles
// @Summary List processes
// @Description Retrieve all processes the user has access to.
// @Tags Processes
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} Process
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /processes [GET]
func getProcesses(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ProcessList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles
//Admin, Admin Read-Only, Global Auditor, Org Manager, Space Auditor, Space Developer, Space Manager
// @Summary List processes for app
// @Description Retrieves all processes belonging to an app.
// @Tags Processes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "app guid"
// @Success 200 {object} Process
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/processes [GET]
func getAppProcesses(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ProcessList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles
//Admin, Space Developer
// TypeError: Failed to fetch
// @Summary Update a process
// @Description
// @Tags Processes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "app guid"
// @Param command body string false "The command used to start the process; use null to revert to the buildpack-detected or procfile-provided start command"
// @Param health_check body string false "The health check to perform on the process"
// @Param metadata.labels body string false "Labels applied to the process"
// @Param metadata.annotations body string false "Annotations applied to the process"
// @Success 200 {object} Process
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /processes/{guid} [PATCH]
func updateProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdateProcess
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final Process
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles Admin, Space Developer
// @Summary Scale a process
// @Description
// @Tags Processes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "app guid"
// @Param instances body integer false "The number of instances to run"
// @Param memory_in_mb body integer false "The memory in mb allocated per instance"
// @Param disk_in_mb body integer false "The disk in mb allocated per instance"
// @Success 200 {object} Process
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /processes/{guid}/actions/scale [POST]
func scaleProcess(w http.ResponseWriter, r *http.Request) {
	var pBody ScaleProcess
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)

	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/actions/scale", reqBody, "POST", w, r)
	if rBodyResult {
		var final Process
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles Admin, Space Developer
// @Summary Scale a process
// @Description
// @Tags Processes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "app guid"
// @Param type path string true "app type"
// @Param instances body integer false "The number of instances to run"
// @Param memory_in_mb body integer false "The memory in mb allocated per instance"
// @Param disk_in_mb body integer false "The disk in mb allocated per instance"
// @Success 200 {object} Process
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/processes/{type}/actions/scale [POST]
func scaleAppProcess(w http.ResponseWriter, r *http.Request) {
	var pBody ScaleProcess
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)

	vars := mux.Vars(r)
	guid := vars["guid"]
	appType := vars["type"]

	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"/"+appType+"/actions/scale", reqBody, "POST", w, r)
	if rBodyResult {
		var final Process
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted Roles, Admin, Space Developer
// index에 대한 값을 모르겠음.
// @Summary Terminate a process instance
// @Description Terminate an instance of a specific process. Health management will eventually restart the instance.
// @Description This allows a user to stop a single misbehaving instance of a process.
// @Tags Processes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "app guid"
// @Param type path string true "app type"
// @Success 200 {object} Process
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /processes/{guid}/instances/{index} [DELETE]
func terminateProcessInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	index := vars["index"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/instances/"+index, nil, "DELETE", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted Roles, Admin, Space Developer
// index에 대한 값을 모르겠음.
// @Summary Terminate a process instance
// @Description Terminate an instance of a specific process. Health management will eventually restart the instance.
// @Description This allows a user to stop a single misbehaving instance of a process.
// @Tags Processes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "app guid"
// @Param type path string true "app type"
// @Success 200 {object} Process
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/processes/{type}/instances/{index} [DELETE]
func terminateAppProcessInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	index := vars["index"]
	appType := vars["type"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"/"+appType+"/instances/"+index, nil, "DELETE", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
