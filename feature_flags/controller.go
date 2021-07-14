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

// @Description Permitted Roles All Roles
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

// @Description Permitted Roles All Roles
// @Summary List feature flags
// @Description Retrieve all feature_flags.
// @Tags Feature Flags
// @Produce  json
// @Security ApiKeyAuth
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
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

// @Description Permitted Roles Admin
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
