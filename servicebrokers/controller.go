package servicebrokers

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

func ServiceBrokerHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/service_brokers", createServiceBroker).Methods("POST")
	myRouter.HandleFunc("/v3/service_brokers/{guid}", getServiceBroker).Methods("GET")
	myRouter.HandleFunc("/v3/service_brokers/", getServiceBrokers).Methods("GET")
	myRouter.HandleFunc("/v3/service_brokers/{guid}", updateServiceBrokers).Methods("PATCH")
	myRouter.HandleFunc("/v3/service_brokers/{guid}", deleteServiceBrokers).Methods("DELETE")
}

//Permitted roles 'Admin Space Developer'
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

	rBody, rBodyResult := config.Curl("/v3/service_brokers", reqBody, "POST", w, r)
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

//Permitted roles 'Admin Admin Read-Only Global Auditor Space Developer (only space-scoped brokers)'
func getServiceBroker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/service_brokers/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceBroker
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin Admin Read-Only Global Auditor Space Developer (only space-scoped brokers)'
func getServiceBrokers(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/service_brokers?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final ServiceBrokerList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin Space Developer (only space-scoped brokers)'
func updateServiceBrokers(w http.ResponseWriter, r *http.Request) {
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

	rBody, rBodyResult := config.Curl("/v3/service_brokers/"+guid, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final ServiceBroker
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin Space Developer (only space-scoped brokers)'
func deleteServiceBrokers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/service_brokers/"+guid, nil, "DELETE", w, r)
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
