package feature_flags

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "feature_flags"

func FeatureFlagHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{name}", getFeatureFlag).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getFeatureFlags).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{name}", updateFeatureFlags).Methods("PATCH")
}

//Permitted Roles All Roles
// @Summary Get a feature flag
// @Description
// @Tags Feature Flags
// @Produce  json
// @Security ApiKeyAuth
// @Param name path string true "FeatureFlag name"
// @Success 200 {object} FeatureFlags
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /feature_flags/{name} [GET]
func getFeatureFlag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+name, nil, "GET", w, r)
	if rBodyResult {
		var final FeatureFlags
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//Permitted Roles All Roles
// @Summary List feature flags
// @Description Retrieve all feature_flags.
// @Tags Feature Flags
// @Produce  json
// @Security ApiKeyAuth
// @Param query query string false "query"
// @Success 200 {object} GetFeatureFlags
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /feature_flags [GET]
func getFeatureFlags(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final GetFeatureFlags
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//Permitted Roles Admin
// @Summary Update a feature flag
// @Description
// @Tags Feature Flags
// @Produce  json
// @Security ApiKeyAuth
// @Param name path string true "FeatureFlag Name"
// @Param UpdateFeatureFlags body UpdateFeatureFlags true "Update Feature Flags"
// @Success 200 {object} FeatureFlags
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /feature_flags/{name} [PATCH]
func updateFeatureFlags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var pBody UpdateFeatureFlags
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+name, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final FeatureFlags
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
