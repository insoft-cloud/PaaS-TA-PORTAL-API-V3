package service_credential_bindings

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "service_credential_bindings"

func ServiceCredentialBindingHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createServiceCredentialBinding).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getServiceCredentialBinding).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getServiceCredentialBindings).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateServiceCredentialBinding).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteServiceCredentialBinding).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/details", getServiceCredentialBindingDetail).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/parameters", getServiceCredentialBindingParameter).Methods("GET")
}

// @Description Permitted Roles 'Admin SpaceDeveloper'
// @Summary Create a service credential binding
// @Description This endpoint creates a new service credential binding. Service credential bindings can be of type app or key; key is only valid for managed service instances.
// @Description If failures occur when creating a service credential binding for a managed service instances, the API might execute orphan mitigation steps accordingly to cases outlined in the OSBAPI specification
// @Tags Service Credential Bindings
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateServiceCredentialBinding body CreateServiceCredentialBinding true "Create ServiceCredentialBinding"
// @Success 202 {object} ServiceCredentialBinding
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_credential_bindings [POST]
func createServiceCredentialBinding(w http.ResponseWriter, r *http.Request) {
	var pBody CreateServiceCredentialBinding
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
		var final ServiceCredentialBinding
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get a service credential binding
// @Description This endpoint retrieves the service credential binding by GUID.
// @Tags Service Credential Bindings
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceCredentialBinding Guid"
// @Param include query []string false "Optionally include a list of unique related resources in the response. Valid values are: app, service_instance" collectionFormat(multi)
// @Success 200 {object} ServiceCredentialBinding
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_credential_bindings/{guid} [GET]
func getServiceCredentialBinding(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceCredentialBinding
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary List service credential bindings
// @Description This endpoint retrieves the service credential bindings the user has access to.
// @Tags Service Credential Bindings
// @Produce  json
// @Security ApiKeyAuth
// @Param guids query []string false "Comma-delimited list of service credential binding guids to filter by" collectionFormat(csv)
// @Param service_instance_guids query []string false "Comma-delimited list of service instance guids to filter by" collectionFormat(csv)
// @Param service_instance_names query []string false "Comma-delimited list of service instance names to filter by" collectionFormat(csv)
// @Param app_guids query []string false "Comma-delimited list of app guids to filter by" collectionFormat(csv)
// @Param app_names query []string false "Comma-delimited list of app names to filter by" collectionFormat(csv)
// @Param service_plan_guids query []string false "Comma-delimited list of service plan guids to filter by" collectionFormat(csv)
// @Param service_plan_names query []string false "Comma-delimited list of service plan names to filter by" collectionFormat(csv)
// @Param service_offering_guids query []string false "Comma-delimited list of service offering guids to filter by" collectionFormat(csv)
// @Param service_offering_names query []string false "Comma-delimited list of service offering names to filter by" collectionFormat(csv)
// @Param type query []string false "Type of credential binding to filter by. Valid values are: app or key" collectionFormat(csv)
// @Param guids query []string false "omma-delimited list of service route binding guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param include query []string false "Optionally include a list of unique related resources in the response. Valid values are: app, service_instance" collectionFormat(multi)
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} ServiceCredentialBindingList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_credential_bindings [GET]
func getServiceCredentialBindings(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceCredentialBindingList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Update a service credential binding
// @Description This endpoint updates a service credential binding with labels and annotations.
// @Tags Service Credential Bindings
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceCredentialBinding Guid"
// @Param UpdateServiceCredentialBinding body UpdateServiceCredentialBinding true "Update ServiceCredentialBinding"
// @Success 200 {object} ServiceCredentialBinding
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_credential_bindings/{guid} [PATCH]
func updateServiceCredentialBinding(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdateServiceCredentialBinding
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
		var final ServiceCredentialBinding
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Delete a service credential binding
// @Description This endpoint deletes a service credential binding. When deleting credential bindings originated from user provided service instances, the delete operation does not require interactions with service brokers, therefore the API will respond synchronously to the delete request.
// @Tags Service Credential Bindings
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceCredentialBinding Guid"
// @Success 202 {object} string "ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_credential_bindings/{guid} [DELETE]
func deleteServiceCredentialBinding(w http.ResponseWriter, r *http.Request) {
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

// @Description Permitted Roles 'Admin, Admin Read-Only Space Developer'
// @Summary Get a service credential binding details
// @Description This endpoint retrieves the service credential binding details.
// @Tags Service Credential Bindings
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceCredentialBinding Guid"
// @Success 200 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_credential_bindings/{guid}/details [GET]
func getServiceCredentialBindingDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/details", nil, "GET", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Admin Read-Only Space Developer'
// @Summary Get parameters for a service credential binding
// @Description Queries the Service Broker for the parameters associated with this service credential binding. The broker catalog must have enabled the bindings_retrievable feature for the Service Offering. Check the Service Offering object for the value of this feature flag. This endpoint is not available for User-Provided Service Instances.
// @Tags Service Credential Bindings
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceCredentialBinding Guid"
// @Success 200 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_credential_bindings/{guid}/parameters [GET]
func getServiceCredentialBindingParameter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/parameters", nil, "GET", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
