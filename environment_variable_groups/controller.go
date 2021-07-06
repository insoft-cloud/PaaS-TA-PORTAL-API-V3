package environment_variable_groups

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"fmt"
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

//Permitted All Roles
func getEnvironmentVariableGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Println("getEnvironmentVariableGroup?" + name)
	if name == "staging" {
		return
	}
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+name, nil, "GET", w, r)
	if rBodyResult {
		var final EnvironmentVariableGroup
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//Permitted roles 'Admin'
func updateEnvironmentVariableGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Println("updateEnvironmentVariableGroup?" + name)
	if name == "staging" {
		return
	}
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
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
