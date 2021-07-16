package service_brokers

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "service_brokers"

func ServiceBrokerHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createServiceBroker).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getServiceBroker).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getServiceBrokers).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateServiceBroker).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteServiceBroker).Methods("DELETE")
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Create a service broker
// @Description This endpoint creates a new service broker and a job to synchronize the service offerings and service plans with those in the broker’s catalog. The Location header refers to the created job which syncs the broker with the catalog. See Service broker jobs for more information and limitations.
// @Tags Service Brokers
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateServiceBroker body CreateServiceBroker true "Create ServiceBroker"
// @Success 202 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_brokers [POST]
func createServiceBroker(w http.ResponseWriter, r *http.Request) {
	var pBody CreateServiceBroker
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
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Space Developer (only space-scoped brokers)'
// @Summary Get a service broker
// @Description This endpoint retrieves the service broker by GUID.
// @Tags Service Brokers
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceBroker Guid"
// @Success 200 {object} ServiceBroker
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_brokers/{guid} [GET]
func getServiceBroker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceBroker
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Space Developer (only space-scoped brokers)'
// @Summary List service brokers
// @Description This endpoint retrieves the service brokers the user has access to.
// @Tags Service Brokers
// @Produce  json
// @Security ApiKeyAuth
// @Param names query []string false "Comma-delimited list of service broker names to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} ServiceBrokerList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_brokers [GET]
func getServiceBrokers(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceBrokerList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer (only space-scoped brokers)'
// @Summary Update a service broker
// @Description This endpoint updates a service broker. Depending on the parameters specified, the endpoint may respond with a background job, and it may synchronize the service offerings and service plans with those in the broker’s catalog.
// @Description When a service broker has a synchronization job in progress, only updates with metadata are permitted until the synchronization job is complete.
// @Tags Service Brokers
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceBroker Guid"
// @Param UpdateServiceBroker body UpdateServiceBroker true "Update ServiceBroker"
// @Success 202 {object} ServiceBroker
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_brokers/{guid} [PATCH]
func updateServiceBroker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateServiceBroker
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
		var final ServiceBroker
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer (only space-scoped brokers)'
// @Summary Delete a service broker
// @Description This endpoint creates a job to delete an existing service broker. The Location header refers to the created job. See Service broker jobs for more information and limitations.
// @Tags Service Brokers
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceBroker Guid"
// @Success 202 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_brokers/{guid} [DELETE]
func deleteServiceBroker(w http.ResponseWriter, r *http.Request) {
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
