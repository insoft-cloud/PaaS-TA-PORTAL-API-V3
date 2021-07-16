package environment_variable_groups

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	_ "net/url"
)

var uris = "environment_variable_groups"

func EnvironmentVariableGroupsHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{name}", getEnvironmentVariableGroup).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{name}", updateEnvironmentVariableGroup).Methods("PATCH")

}

// @Description Permitted Roles 'All'
// @Summary Get an environment variable group
// @Description
// @Tags Environment Variable Groups
// @Produce  json
// @Security ApiKeyAuth
// @Param name path string true "Environment variable name"
// @Success 200 {object} EnvironmentVariableGroup
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /environment_variable_groups/{name} [GET]
func getEnvironmentVariableGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "staging" {
		return
	}
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+name, nil, "GET", w, r)
	if rBodyResult {
		var final EnvironmentVariableGroup
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin'
// @Summary Update environment variable group
// @Description Update the environment variable group. The variables given in the request will be merged with the existing environment variable group. Any requested variables with a value of null will be removed from the group. Environment variable names may not start with VCAP_. PORT is not a valid environment variable.
// @Tags Environment Variable Groups
// @Produce  json
// @Security ApiKeyAuth
// @Param name path string true "Environment variable name"
// @Param interface{} body interface{} true "Update environment variable group"
// @Success 200 {object} EnvironmentVariableGroup
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /environment_variable_groups/{name} [PATCH]
func updateEnvironmentVariableGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var pBody interface{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+name, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final EnvironmentVariableGroup
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
