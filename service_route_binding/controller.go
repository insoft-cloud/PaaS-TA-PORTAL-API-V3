package service_route_binding

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "service_route_bindings"

func ServiceRouteBindingHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getServiceRouteBinding).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getServiceRouteBindings).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, createServiceRouteBinding).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateServiceRouteBinding).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteServiceRouteBinding).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/parameters", getServiceRouteBindingParameter).Methods("GET")
}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get a service route binding
// @Description This endpoint retrieves the service route binding by GUID.
// @Tags Service Route Binding
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceRouteBinding Guid"
// @Param include query []string false "Optionally include a list of related resources in the response; valid values are space.organization and service_offering" collectionFormat(multi)
// @Success 200 {object} ServiceRouteBinding
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_route_bindings/{guid} [GET]
func getServiceRouteBinding(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceRouteBinding
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary List service route bindings
// @Description This endpoint retrieves the service route bindings the user has access to.
// @Tags Service Route Binding
// @Produce  json
// @Security ApiKeyAuth
// @Param route_guids query []string false "Comma-delimited list of route guids to filter by" collectionFormat(csv)
// @Param service_instance_guids query []string false "Comma-delimited list of service instance guids to filter by" collectionFormat(csv)
// @Param service_instance_names query []string false "Comma-delimited list of service instance names to filter by" collectionFormat(csv)
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param guids query []string false "Comma-delimited list of service route binding guids to filter by" collectionFormat(csv)
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param include query []string false "Optionally include a list of unique related resources in the response; valid values are space and spaceorganization" collectionFormat(multi)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Success 200 {object} ServiceRouteBindingList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_route_bindings [GET]
func getServiceRouteBindings(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceRouteBindingList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, SpaceDeveloper'
// @Summary Create a service route binding
// @Description This endpoint creates a new route service binding. The service instance and the route must be in the same space.
// @Description To bind a route to a user-provided service instance, the service instance must have the route_service_url property set.
// @Description To bind a route to a managed service instance, the service offering must be bindable, and the service offering must have route_forwarding set in the requires property.
// @Tags Service Route Binding
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateServiceRouteBinding body CreateServiceRouteBinding true "Create ServiceRouteBinding"
// @Success 201 {object} ServiceRouteBinding
// @Success 202 {object} ServiceRouteBinding
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_route_bindings [POST]
func createServiceRouteBinding(w http.ResponseWriter, r *http.Request) {
	var pBody CreateServiceRouteBinding
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
		var final ServiceRouteBinding
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Update a service route binding
// @Description This endpoint updates a service route binding with labels and annotations.
// @Tags Service Route Binding
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceRouteBinding Guid"
// @Param UpdateServiceRouteBinding body UpdateServiceRouteBinding true "Update ServiceRouteBinding"
// @Success 200 {object} ServiceRouteBinding
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_route_bindings/{guid} [PATCH]
func updateServiceRouteBinding(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdateServiceRouteBinding
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
		var final ServiceRouteBinding
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Delete a service route binding
// @Description This endpoint deletes a service route binding. When deleting route bindings originating from user provided service instances, the delete operation does not require interactions with service brokers, therefore the API will respond synchronously to the delete request. Consequently, deleting route bindings from managed service instances responds with a job which can be used to track the progress of the delete operation.
// @Tags Service Route Binding
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceRouteBinding Guid"
// @Success 201 {object} string "ok"
// @Success 204 {object} string "ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_route_bindings/{guid} [DELETE]
func deleteServiceRouteBinding(w http.ResponseWriter, r *http.Request) {
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
// @Summary Get parameters for a route binding
// @Description Queries the Service Broker for the parameters associated with this service route binding. The broker catalog must have enabled the bindings_retrievable feature for the Service Offering. Check the Service Offering object for the value of this feature flag. This endpoint is not available for User-Provided Service Instances.
// @Tags Service Route Binding
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceRouteBinding Guid"
// @Success 200 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_route_bindings/{guid}/parameters [GET]
func getServiceRouteBindingParameter(w http.ResponseWriter, r *http.Request) {
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
