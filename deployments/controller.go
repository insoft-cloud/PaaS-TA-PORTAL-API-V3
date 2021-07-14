package deployments

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "deployments"

func DeploymentHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createDeployment).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getDeployment).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getDeployments).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateDeployment).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", cancelDeployment).Methods("DELETE")
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Create a deployment
// @Description
// @Tags Deployments
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateDeployment body CreateDeployment true "Create Deployment"
// @Success 200 {object} Deployment
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /deployments [POST]
func createDeployment(w http.ResponseWriter, r *http.Request) {
	var pBody CreateDeployment
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
		var final Deployment
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get a deployment
// @Description
// @Tags Deployments
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Deployment Guid"
// @Success 200 {object} Deployment
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /deployments/{guid} [GET]
func getDeployment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Deployment
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Read-Only Admin Global Auditor Org Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary List  deployment
// @Description Retrieve all deployments the user has access to.
// @Tags Deployments
// @Produce  json
// @Security ApiKeyAuth
// @Param app_guids query []string false "Comma-delimited list of app guids to filter by" collectionFormat(csv)
// @Param states query []string false "Comma-delimited list of build states to filter by" collectionFormat(csv)
// @Param status_reasons query []string false "Comma-delimited list of status reasons to filter by; valid values include DEPLOYING, CANCELING, DEPLOYED, CANCELED, SUPERSEDED, and DEGENERATE" collectionFormat(csv)
// @Param status_values query []string false "Comma-delimited list of status values to filter by; valid values include ACTIVE and FINALIZED" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} DeploymentList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /deployments [GET]
func getDeployments(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final DeploymentList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Update a deployment
// @Description
// @Tags Deployments
// @Produce  json
// @Security ApiKeyAuth
// @Param UpdateDeployment body UpdateDeployment true "Update Deployment"
// @Param guid path string true "Deployment Guid"
// @Success 200 {object} Deployment
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /deployments/{guid} [PATCH]
func updateDeployment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdateDeployment
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
		var final Deployment
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Cancel a deployment
// @Description
// @Tags Deployments
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Deployment Guid"
// @Success 200 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /deployments/{guid}/actions/cancel [POST]
func cancelDeployment(w http.ResponseWriter, r *http.Request) {
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
