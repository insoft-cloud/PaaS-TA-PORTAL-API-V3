package routes

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "routes"

func RouteHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createRoute).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getRoute).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getRoutes).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris, getAppRoutes).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateRoute).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteRoute).Methods("DELETE")
	myRouter.HandleFunc("/v3/domains/{guid}/route_reservations", checkReservedRoutesForDomain).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/destinations", getDestinationsRoute).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/destinations", insertDestinationsForRoute).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/destinations", replaceAllDestinationsForRoute).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/destinations/{destination_guid}", removeDestinationForRoute).Methods("DELETE")
	myRouter.HandleFunc("/v3/space/{guid}/routes?unmapped=true", deleteUnmappedRoutesForSpace).Methods("DELETE")

}

// @Description Permitted roles 'Admin, Space Developer'
// @Summary Create a route
// @Description
// @Tags Routes
// @Produce json
// @Security ApiKeyAuth
// @Param CreateRoutes body CreateRoute true "Create Routes"
// @Success 202 {object} Route
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /routes [POST]
func createRoute(w http.ResponseWriter, r *http.Request) {
	var pBody CreateRoute
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
		var final Route
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles 'Admin Read-Only, Admin, Global Auditor, Org Auditor, Org Manager, Space Auditor, Space Developer, Space Manager'
// @Summary Get a route
// @Description
// @Tags Routes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "app guid"
// @Param include query string false "Optionally include additional related resources in the response Valid values are domain, space.organization, space"
// @Success 200 {object} Route
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /routes/{guid} [GET]
func getRoute(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final Route
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles 'All Roles'
// @Summary List routes
// @Description Retrieve all routes the user has access to.
// @Tags Routes
// @Produce json
// @Security ApiKeyAuth
// @Param app_guids query []string false "Comma-delimited list of app guids to filter by" collectionFormat(csv)
// @Param domain_guids query []string false "Comma-delimited list of domain guids to filter by" collectionFormat(csv)
// @Param hosts query []string false "Comma-delimited list of hostnames to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param paths query []string false "Comma-delimited list of paths to filter by (e.g. /path1,/path2)" collectionFormat(csv)
// @Param ports query []integer false "Comma-delimited list of ports to filter by (e.g. 3306,5432)" collectionFormat(csv)
// @Param space_guids query []string false "Comma-delimited list of space guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by; defaults to ascending. Prepend with - to sort descending. Valid values are created_at, updated_at"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param include query string false "Optionally include a list of unique related resources in the response Valid values are domain, space.organization, space"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} Route
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /routes [GET]
func getRoutes(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final RouteList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles 'Admin, Admin Read-Only, Global Auditor, Org Manager, Space Auditor, Space Developer, Space Manager'
// @Summary List routes for an app
// @Description Retrieve all routes that have destinations that point to the given app.
// @Tags Routes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "apps guid"
// @Param domain_guids query []string false "Comma-delimited list of domain guids to filter by" collectionFormat(csv)
// @Param hosts query []string false "Comma-delimited list of hostnames to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param paths query []string false "Comma-delimited list of paths to filter by (e.g. /path1,/path2)" collectionFormat(csv)
// @Param ports query []integer false "Comma-delimited list of ports to filter by (e.g. 3306,5432)" collectionFormat(csv)
// @Param space_guids query []string false "Comma-delimited list of space guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by; defaults to ascending. Prepend with - to sort descending. Valid values are created_at, updated_at"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Success 200 {object} Route
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/routes [GET]
func getAppRoutes(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final RouteList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles 'Admin, Space Developer'
// @Summary Update a route
// @Description
// @Tags Routes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "route guid"
// @Param UpdateRoute body UpdateRoute false "Update Routes"
// @Success 200 {object} Route
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /routes/{guid} [PATCH]
func updateRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdateRoute
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final Route
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles 'Admin, Space Developer'
// @Summary Delete a route
// @Description
// @Tags Routes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "route guid"
// @Success 202 {object} string "Accepted"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /routes/{guid} [DELETE]
func deleteRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "DELETE", w, r)
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

// @Description Permitted roles 'Admin, Admin Read-Only, Global Auditor, Org Auditor, Org Billing Manager Can only check if routes exist for a domain without an organization relationship,
// @Description Org Manager, Space Auditor, Space Developer, Space Manager'
// @Summary Check reserved routes for a domain
// @Description Check if a specific route for a domain exists, regardless of the user’s visibility for the route in case the route belongs to a space the user does not belong to.
// @Tags Routes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "domain guid"
// @Param host query string false "Hostname to filter by; defaults to empty string if not provided and only applicable to http routes"
// @Param path query string false "Path to filter by; defaults to empty string if not provided and only applicable to http routes"
// @Param port query integer false "Port to filter by; only applicable to tcp routes and required for tcp routes"
// @Success 200 {object} Route
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /domains/{guid}/route_reservations [GET]
func checkReservedRoutesForDomain(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/domains/"+guid+"/route_reservations"+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final CheckReservedRoutesForDomain
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles 'Admin, Admin Read-Only, Global Auditor, Org Auditor, Org Manager, Space Auditor, Space Developer, Space Manager'
// @Summary List destinations for a route
// @Description Retrieve all destinations associated with a route.
// @Tags Routes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "route guid"
// @Param guids path string false "Comma-delimited list of destination guids to filter by"
// @Param app_guids path string false "Comma-delimited list of app guids to filter by"
// @Success 200 {object} Route
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /routes/{guid}/destinations [GET]
func getDestinationsRoute(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/destinations"+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final DestinationsRouteList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles 'Admin, Space Developer'
// @Summary Insert destinations for a route
// @Description Add one or more destinations to a route, preserving any existing destinations.
// @Description Note that weighted destinations cannot be added with this endpoint. To add weighted destinations, replace all destinations for a route at once using the replace destinations endpoint.
// @Tags Routes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "route guid"
// @Param destinations body string true "List of destinations to add to route; destinations without process.type specified will get process type "web" by default"
// @Success 200 {object} Route
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /routes/{guid}/destinations [POST]
func insertDestinationsForRoute(w http.ResponseWriter, r *http.Request) {
	var pBody InsertDestinations
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

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/destinations", reqBody, "POST", w, r)
	if rBodyResult {
		var final Route
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles 'Admin, Space Developer'
// @Summary Replace all destinations for a route
// @Description Replaces all destinations for a route, removing any destinations not included in the provided list.
// @Description If using weighted destinations, all destinations provided here must have a weight specified,
// @Description and all weights for this route must sum to 100. If not, all provided destinations must not have a weight.
// @Description Mixing weighted and unweighted destinations for a route is not allowed.
// @Tags Routes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "route guid"
// @Param replaceAllDestinationsForRoute body ReplaceAllDestinationRoute true "List of destinations use for route. Destinations without process.type specified will get process type "web" by default"
// @Success 200 {object} Route
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /routes/{guid}/destinations [PATCH]
func replaceAllDestinationsForRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody ReplaceAllDestinationRoute
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/destinations", reqBody, "PATCH", w, r)
	if rBodyResult {
		var final Route
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted roles 'Admin, Space Developer'
// @Summary Remove destination for a route
// @Description Remove a destination from a route.
// @Tags Routes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "route guid"
// @Param destination_guid path string true "destination guid"
// @Success 204 {object} string "No Content"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /routes/{guid}/destinations/{destination_guid} [DELETE]
func removeDestinationForRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	destinationGuid := vars["destination_guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/destinations/"+destinationGuid, nil, "DELETE", w, r)
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

// @Description Permitted roles "Admin", "Space Developer"
//Note: unmapped=true is a required query parameter; always include it.
// @Summary Delete unmapped routes for a space
// @Description Deletes all routes in a space that are not mapped to any applications and not bound to any service instances.
// @Description Note: unmapped=true is a required query parameter; always include it.
// @Tags Routes
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "route guid"
// @Param destination_guid path string true "destination guid"
// @Param query path string true "unmapped=true"
// @Success 202 {object} string "Accepted"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /space/{guid}/routes [DELETE]
func deleteUnmappedRoutesForSpace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/space/"+guid+"/"+uris+"?unmapped=true", nil, "DELETE", w, r)
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
