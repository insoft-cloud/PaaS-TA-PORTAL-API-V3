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

//Permitted roles, Admin, Space Developer
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

//Permitted roles "Admin" Read-Only, "Admin", "Global Auditor", "Org Auditor", "Org Manager", "Space Auditor", "Space Developer", "Space Manager"
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

//Permitted roles "All Roles"
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

//Permitted roles "Admin", "Space Developer"
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

//Permitted roles "Admin", "Space Developer"
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

//Permitted roles
//Role	Notes
//Admin
//Admin Read-Only
//Global Auditor
//Org Auditor
//Org Billing Manager	Can only check if routes exist for a domain without an organization relationship
//Org Manager
//Space Auditor
//Space Developer
//Space Manager
func checkReservedRoutesForDomain(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/domains/"+guid+"/route_reservations"+"?"+query, nil, "GET", w, r)
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

//Permitted roles "Admin" Read-Only "Admin" "Global Auditor" "Org Auditor" "Org Manager" "Space Auditor" "Space Developer" "Space Manager"
func getDestinationsRoute(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/destinations"+"?"+query, nil, "GET", w, r)
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

func insertDestinationsForRoute(w http.ResponseWriter, r *http.Request) {
	var pBody insertDestinations
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

//Permitted roles
//Role
//Admin
//Space Developer
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

//Permitted roles "Admin", "Space Developer"
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

//Permitted roles "Admin", "Space Developer"
//Note: unmapped=true is a required query parameter; always include it.
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
