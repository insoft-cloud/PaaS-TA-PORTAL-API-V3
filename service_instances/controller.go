package service_instances

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "service_instances"

func ServiceInstanceHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createServiceInstance).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/user-provide", createServiceInstanceUserProvide).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris, getServiceInstances).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getServiceInstance).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/credentials", getCredential).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/parameters", getParameter).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateServiceInstance).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteServiceInstance).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/shared_spaces", getShareSpacesRelationship).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/shared_spaces", shareServiceInstance).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/shared_spaces/{space_guid}", unShareServiceInstance).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/shared_spaces/usage_summary", getUsageSummary).Methods("GET")
}

// @Description Permitted Roles 'Admin, SpaceDeveloper'
// @Summary Create a service instance
// @Description This endpoint creates a new service instance. Service instances can be of type managed or user-provided, and the required parameters are different for each type. User provided service instances do not require interactions with service brokers.
// @Description If failures occur when creating managed service instances, the API might execute orphan mitigation steps accordingly to cases outlined in the OSBAPI specification
// @Tags Service Instances
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateServiceInstanceProvide body CreateServiceInstanceProvide true "Create ServiceInstanceProvide"
// @Success 202 {object} ServiceInstance
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_instances [POST]
func createServiceInstance(w http.ResponseWriter, r *http.Request) {
	var pBody CreateServiceInstance
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
		var final ServiceInstance
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, SpaceDeveloper'
// @Summary Create a service instance
// @Description This endpoint creates a new service instance. Service instances can be of type managed or user-provided, and the required parameters are different for each type. User provided service instances do not require interactions with service brokers.
// @Description If failures occur when creating managed service instances, the API might execute orphan mitigation steps accordingly to cases outlined in the OSBAPI specification
// @Tags Service Instances
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateServiceInstanceProvide body CreateServiceInstanceProvide true "Create ServiceInstanceProvide"
// @Success 202 {object} ServiceInstance
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_instances [POST]
func createServiceInstanceUserProvide(w http.ResponseWriter, r *http.Request) {
	var pBody CreateServiceInstanceProvide
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
		var final ServiceInstance
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary List service instances
// @Description This endpoint retrieves the service instances the user has access to. At the moment, this endpoint only returns managed service instances. This may change in the future.
// @Description This includes access granted by service instance sharing.
// @Tags Service Instances
// @Produce  json
// @Security ApiKeyAuth
// @Param names query []string false "Comma-delimited list of service instance names to filter by" collectionFormat(csv)
// @Param space_guids query []string false "Comma-delimited list of space guids to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param service_plan_guids query []string false "Comma-delimited list of service plan guids to filter by" collectionFormat(csv)
// @Param service_plan_names query []string false "Comma-delimited list of service plan names to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param fields[space] query string false "string enums" Enums(guid, name, relationships.organization)
// @Param fields[space.organization] query string false "string enums" Enums(guid, name)
// @Param fields[service_plan] query string false "string enums" Enums(guid, name, relationships.service_offering)
// @Param fields[service_plan.service_offering] query string false "string enums" Enums(guid, name, description, documentation_url, tags, relationships.service_broker)
// @Param fields[service_plan.service_offering.service_broker] query string false "string enums" Enums(guid, name)
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} ServiceInstanceList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_instances [GET]
func getServiceInstances(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceInstanceList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get a service instance
// @Description This endpoint retrieves the service instance by GUID.
// @Tags Service Instances
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceInstance Guid"
// @Param fields[space] query string false "string enums" Enums(guid, name)
// @Param fields[space.organization] query string false "string enums" Enums(guid, name)
// @Param fields[service_plan] query string false "string enums" Enums(guid, name)
// @Param fields[service_plan.service_offering] query string false "string enums" Enums(name, guid, description, documentation_url, tags)
// @Param fields[service_plan.service_offering.service_broker] query string false "string enums" Enums(guid, name)
// @Success 200 {object} ServiceInstance
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_instances/{guid} [GET]
func getServiceInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceInstance
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Admin Read-Only Space Developer Space Manager'
// @Summary Get credentials for a user-provided service instance
// @Description Retrieves the credentials for a user-provided service instance. This endpoint is not available for managed service instances.
// @Tags Service Instances
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceInstance Guid"
// @Success 200 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_instances/{guid}/credentials [GET]
func getCredential(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/credentials", nil, "GET", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Queries the Service Broker for the parameters associated with this service instance. The broker catalog must have enabled the instances_retrievable feature for the Service Offering. Check the Service Offering object for the value of this feature flag.
// @Description Retrieves the credentials for a user-provided service instance. This endpoint is not available for managed service instances.
// @Tags Service Instances
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceInstance Guid"
// @Success 200 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_instances/{guid}/parameters [GET]
func getParameter(w http.ResponseWriter, r *http.Request) {
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

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Update a service instance
// @Description
// @Tags Service Instances
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceCredentialBinding Guid"
// @Param UpdateServiceInstance body UpdateServiceInstance true "Update ServiceInstance"
// @Success 200 {object} ServiceInstance
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_instances/{guid} [PATCH]
func updateServiceInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdateServiceInstance
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
		var final ServiceInstance
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Delete a service instance
// @Description This endpoint deletes a service instance and any associated service credential bindings or service route bindings. The service instance is removed from all spaces where it is available.
// @Description User provided service instances do not require interactions with service brokers, therefore the API will respond synchronously to the delete request.
// @Description For managed service instances, the API will respond asynchronously. If a service credential binding or service route binding cannot be deleted synchronously, then the operation will fail, and the deletion of the binding will continue in the background. The operation can be retried until it is successful.
// @Tags Service Instances
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceCredentialBinding Guid"
// @Param purge query boolean false "If true, deletes the service instance and all associated resources without any interaction with the service broker."
// @Success 202 {object} ServiceInstance
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_instances/{guid} [DELETE]
func deleteServiceInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"?"+query, nil, "DELETE", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary List shared spaces relationship
// @Description This endpoint lists the spaces that the service instance has been shared to.
// @Tags Service Instances
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceInstance Guid"
// @Param fields[space] query string false "string enums" Enums(guid, name, relationships.organization)
// @Param fields[space.organization] query string false "string enums" Enums(guid, name)
// @Success 200 {object} SharedSpacesRelationship
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_instances/{guid}/relationships/shared_spaces [GET]
func getShareSpacesRelationship(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/shared_spaces?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final SharedSpacesRelationship
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Share a service instance to other spaces
// @Description This endpoint shares the service instance with the specified spaces. In order to share into a space the requesting user must be a space developer in the target space.
// @Tags Service Instances
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceInstance Guid"
// @Param SharedServiceInstanceToOtherSpaces body SharedServiceInstanceToOtherSpaces true "Shared ServiceInstanceToOtherSpaces"
// @Success 200 {object} SharedSpacesRelationship
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_instances/{guid}/relationships/shared_spaces [POST]
func shareServiceInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody SharedServiceInstanceToOtherSpaces
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/shared_spaces", reqBody, "POST", w, r)
	if rBodyResult {
		var final SharedSpacesRelationship
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Unshare a service instance from another space
// @Description This endpoint unshares the service instance from the specified space. This will automatically unbind any applications bound to this service instance in the specified space. Unsharing a service instance from a space will not delete any service keys.
// @Tags Service Instances
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceInstance Guid"
// @Param space_guid path string true "SharedSpaces Guid"
// @Success 200 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_instances/{guid}/relationships/shared_spaces/{space_guid} [DELETE]
func unShareServiceInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	spaceGuid := vars["space_guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/shared_spaces/"+spaceGuid, nil, "DELETE", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get usage summary in shared spaces
// @Description This endpoint returns the number of bound apps in spaces where the service instance has been shared to.
// @Tags Service Instances
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "ServiceInstance Guid"
// @Success 200 {object} GetUsageSummaryInSharedSpace
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /service_instances/{guid}/relationships/shared_spaces/usage_summary [DELETE]
func getUsageSummary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/shared_spaces/usage_summary", nil, "GET", w, r)
	if rBodyResult {
		var final GetUsageSummaryInSharedSpace
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
