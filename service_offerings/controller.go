package service_offerings

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "service_offerings"

func ServiceOfferingHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getServiceOffering).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getServiceOfferings).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateServiceBroker).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteServiceBroker).Methods("DELETE")
}

// @Description Permitted Roles 'All'
// @Summary Get a service offering
// @Description This endpoint retrieves the service offering by GUID.
// @Tags Service Offerings
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceOffering Guid"
// @Param fields[service_broker] query string false "string enums" Enums(guid, name)
// @Success 200 {object} ServiceOffering
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_offerings/{guid} [GET]
func getServiceOffering(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceOffering
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary List service offerings
// @Description This endpoint retrieves the service offerings the user has access to.
// @Tags Service Offerings
// @Produce  json
// @Security ApiKeyAuth
// @Param names query []string false "Comma-delimited list of names to filter by" collectionFormat(csv)
// @Param available query boolean false "Filter by the available property; valid values are true or false"
// @Param service_broker_guids query []string false "Comma-delimited list of service broker GUIDs to filter by" collectionFormat(csv)
// @Param service_broker_names query []string false "Comma-delimited list of service broker names to filter by" collectionFormat(csv)
// @Param space_guids query []string false "Comma-delimited list of space GUIDs to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param fields[service_broker] query string false "string enums" Enums(guid, name)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} ServiceOfferingList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_offerings [GET]
func getServiceOfferings(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceOfferingList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer (only space-scoped brokers)'
// @Summary Update a service offering
// @Description This endpoint updates a service offering with labels and annotations.
// @Tags Service Offerings
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceOffering Guid"
// @Param UpdateServiceOffering body UpdateServiceOffering true "Update ServiceOffering"
// @Success 200 {object} UpdateServiceOffering
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_offerings/{guid} [PATCH]
func updateServiceBroker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateServiceOffering
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
		var final ServiceOffering
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer (only space-scoped brokers)'
// @Summary Delete a service offering
// @Description This endpoint deletes a service offering. This is typically used to remove orphan service offerings from the Cloud Foundry database when they have been removed from the service broker catalog, or when the service broker has been removed.
// @Description Note that this operation only affects the Cloud Foundry database, and no attempt is made to contact the service broker.
// @Tags Service Offerings
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceOffering Guid"
// @Param purge query boolean false "If true, any service plans, instances, and bindings associated with this service offering will also be deleted."
// @Success 204 {object} string "ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_offerings/{guid} [DELETE]
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
