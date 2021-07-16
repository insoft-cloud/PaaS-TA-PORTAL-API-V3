package service_plans

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "service_plans"

func ServicePlanHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getServicePlan).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getServicePlans).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateServiceBroker).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteServiceBroker).Methods("DELETE")
}

// @Description Permitted Roles 'All'
// @Summary Get a service plan
// @Description This endpoint retrieves the service plan by GUID.
// @Tags Service Plans
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServicePlan Guid"
// @Param fields[service_offering.service_broker] query string false "string enums" Enums(guid, name)
// @Param include query []string false "Optionally include a list of related resources in the response; valid values are space.organization and service_offering" collectionFormat(multi)
// @Success 200 {object} ServicePlan
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_plans/{guid} [GET]
func getServicePlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServicePlan
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary List service plans
// @Description This endpoint retrieves the service plans the user has access to.
// @Tags Service Plans
// @Produce  json
// @Security ApiKeyAuth
// @Param names query []string false "Comma-delimited list of names to filter by" collectionFormat(csv)
// @Param available query boolean false "Filter by the available property; valid values are true or false"
// @Param broker_catalog_ids query []string false "Comma-delimited list of IDs provided by the service broker for the service plan to filter by" collectionFormat(csv)
// @Param space_guids query []string false "Comma-delimited list of space GUIDs to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization GUIDs to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization GUIDs to filter by" collectionFormat(csv)
// @Param service_broker_guids query []string false "Comma-delimited list of service broker GUIDs to filter by" collectionFormat(csv)
// @Param service_broker_names query []string false "Comma-delimited list of service broker names to filter by" collectionFormat(csv)
// @Param service_offering_guids query []string false "Comma-delimited list of service Offering GUIDs to filter by" collectionFormat(csv)
// @Param service_offering_names query []string false "Comma-delimited list of service Offering names to filter by" collectionFormat(csv)
// @Param service_instance_guids query []string false "Comma-delimited list of service Instance GUIDs to filter by" collectionFormat(csv)
// @Param include query []string false "Optionally include a list of unique related resources in the response; valid values are space.organization and service_offering" collectionFormat(multi)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param fields[service_offering.service_broker] query string false "string enums" Enums(guid, name)
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} ServicePlanList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_plans [GET]
func getServicePlans(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServicePlanList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer (only space-scoped brokers)'
// @Summary Update a service plan
// @Description This endpoint updates a service plan with labels and annotations.
// @Tags Service Plans
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServicePlan Guid"
// @Param UpdatePlan body UpdatePlan false "Update Plan"
// @Success 200 {object} ServicePlan
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_plans/{guid} [PATCH]
func updateServiceBroker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdatePlan
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
		var final ServicePlan
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer (only space-scoped brokers)'
// @Summary Delete a service plan
// @Description This endpoint deletes a service plan. This is used to remove service plans from the Cloud Foundry database when they are no longer provided by the service broker.
// @Tags Service Plans
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServicePlan Guid"
// @Success 204 {object} string "ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_plans/{guid} [DELETE]
func deleteServiceBroker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdatePlan
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
		var final ServicePlan
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
