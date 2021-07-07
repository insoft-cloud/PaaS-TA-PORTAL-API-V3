package info

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var uris = "info"

func InforHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, getPlatformInfo).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/usage_summary", getPlatformUsageSummary).Methods("GET")
}

func getPlatformInfo(w http.ResponseWriter, r *http.Request) {
	rBody, rBodyResult := config.Curl("/v3/"+uris, nil, "GET", w, r)
	if rBodyResult {
		var final PlatformInfo
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted Roles Admin Admin Read-Only Global Auditor
func getPlatformUsageSummary(w http.ResponseWriter, r *http.Request) {
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/usage_summary", nil, "GET", w, r)
	if rBodyResult {
		var final PlatformUsageSummary
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
