package service_plan_visibility

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var uris = "service_plans"

func ServicePlanVisibilityHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/visibility", getServicePlanVisibility).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/visibility", updateServicePlanVisibility).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/visibility", applyServicePlanVisibility).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/visibility/{organization_guid]", deleteServiceBroker).Methods("DELETE")
}

// @Description Permitted Roles 'All'
// @Summary Get a service plan visibility
// @Description This endpoint retrieves the service plan visibility for a given plan.
// @Tags Service Plan Visibility
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServicePlan Guid"
// @Success 200 {object} ServicePlanVisibility
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_plans/{guid}/visibility [GET]
func getServicePlanVisibility(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/visibility", nil, "GET", w, r)
	if rBodyResult {
		var final ServicePlanVisibility
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary Update a service plan visibility
// @Description This endpoint updates a service plan visibility. It behaves similar to the POST service plan visibility endpoint but this endpoint will replace the existing list of organizations when the service plan is organization visible.
// @Tags Service Plan Visibility
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServicePlan Guid"
// @Param UpdateServicePlanVisibility body UpdateServicePlanVisibility true "Update ServicePlanVisibility"
// @Success 200 {object} ServicePlanVisibility
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_plans/{guid}/visibility [PATCH]
func updateServicePlanVisibility(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateServicePlanVisibility
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/visibility", reqBody, "PATCH", w, r)
	if rBodyResult {
		var final ServicePlanVisibility
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary Apply a service plan visibility
// @Description This endpoint applies a service plan visibility. It behaves similar to the PATCH service plan visibility endpoint but this endpoint will append to the existing list of organizations when the service plan is organization visible.
// @Tags Service Plan Visibility
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServicePlan Guid"
// @Param UpdateServicePlanVisibility body UpdateServicePlanVisibility true "Update ServicePlanVisibility"
// @Success 200 {object} ServicePlanVisibility
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_plans/{guid}/visibility [POST]
func applyServicePlanVisibility(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateServicePlanVisibility
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/visibility", reqBody, "POST", w, r)
	if rBodyResult {
		var final ServicePlanVisibility
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer (only space-scoped brokers)'
// @Summary Remove organization from a service plan visibility
// @Description This endpoint removes an organization from a service plan visibility list of organizations. It is only defined for service plans which are org-restricted. It will fail with a HTTP status code of 422 for any other visibility type (e.g. Public).
// @Tags Service Plan Visibility
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServicePlan Guid"
// @Param guid path string true "Organization Guid"
// @Success 204 {object} string "ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_plans/{guid}/visibility/{organization_guid} [DELETE]
func deleteServiceBroker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	org_guid := vars["organization_guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/visibility/"+org_guid, nil, "DELETE", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
