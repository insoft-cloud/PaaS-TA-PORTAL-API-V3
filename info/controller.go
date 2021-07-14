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

// @Summary Get platform info
// @Description
// @Tags Info
// @Produce  json
// @Success 200 {object} PlatformInfo
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /info [GET]
func getPlatformInfo(w http.ResponseWriter, r *http.Request) {
	rBody, rBodyResult := config.Curl("/v3/"+uris, nil, "GET", w, r)
	if rBodyResult {
		var final PlatformInfo
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin Admin Read-Only Global Auditor
// @Summary Get platform usage summary
// @Description This endpoint retrieves a high-level summary of usage across the entire Cloud Foundry installation.
// @Tags Info
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} PlatformUsageSummary
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /info/usage_summary [GET]
func getPlatformUsageSummary(w http.ResponseWriter, r *http.Request) {
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/usage_summary", nil, "GET", w, r)
	if rBodyResult {
		var final PlatformUsageSummary
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
