package apps

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "apps"

func AppHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createApp).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris, getApps).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getApp).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateApp).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteApp).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/droplets/current", getAppDroplet).Methods("get")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/current_droplet", getAppDropletAssociation).Methods("get")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/env", getAppEnv).Methods("get")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/environment_variables", getAppEnvVariables).Methods("get")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/permissions", getAppPermissions).Methods("get")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/current_droplet", setAppDroplet).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/ssh_enabled", getAppSSH).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/actions/start", startApp).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/actions/stop", stopApp).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/actions/restart", restartApp).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/environment_variables", setAppEnv).Methods("PATCH")
}

// @Description Permitted Roles 'Admin, SpaceDeveloper'
// @Summary Create an app
// @Description
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param App body CreateApp true "Create App"
// @Success 200 {object} App
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps [POST]
func createApp(w http.ResponseWriter, r *http.Request) {
	var pBody CreateApp
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
		var final App
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get an app
// @Description
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Param include query []string false "Optionally include additional related resources in the response; valid values are space and space.organization" collectionFormat(multi)
// @Success 200 {object} App
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid} [GET]
func getApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final App
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary List apps
// @Description Retrieve all apps the user has access to.
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guids query []string false "Comma-delimited list of app guids to filter by" collectionFormat(csv)
// @Param names query []string false "Comma-delimited list of app names to filter by" collectionFormat(csv)
// @Param space_guids query []string false "Comma-delimited list of space guids to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param stacks query []string false "Comma-delimited list of stack names to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param lifecycle_type query string false "Lifecycle type to filter by; valid values are buildpack, docker"
// @Param include query []string false "Optionally include a list of unique related resources in the response; valid values are space and spaceorganization" collectionFormat(multi)
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} AppList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps [GET]
func getApps(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final AppList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Update an app
// @Description
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Param App body UpdateApp true "Update App"
// @Success 200 {object} App
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid} [PATCH]
func updateApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateApp
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
		var final App
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Delete an app
// @Description
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid} [DELETE]
func deleteApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "DELETE", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get current droplet
// @Description
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} AppDroplet
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/droplets/current [GET]
func getAppDroplet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/droplets/current", nil, "GET", w, r)
	if rBodyResult {
		var final AppDroplet
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get current droplet association for an app
// @Description This endpoint retrieves the current droplet relationship for an app.
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} AppDropletAssociation
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/relationships/current_droplet [GET]
func getAppDropletAssociation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/current_droplet", nil, "GET", w, r)
	if rBodyResult {
		var final AppDropletAssociation
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Space Developer'
// @Summary Get environment for an app
// @Description Retrieve the environment variables that will be provided to an app at runtime. It will include environment variables for Environment Variable Groups and Service Bindings.
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} AppEnv
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/env [GET]
func getAppEnv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/env", nil, "GET", w, r)
	if rBodyResult {
		var final AppEnv
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Space Developer'
// @Summary Get environment variables for an app
// @Description Retrieve the environment variables that are associated with the given app. For the entire list of environment variables that will be available to the app at runtime, see the env endpoint.
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} AppEnvVariable
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/environment_variables [GET]
func getAppEnvVariables(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/environment_variables", nil, "GET", w, r)
	if rBodyResult {
		var final AppEnvVariable
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get permissions
// @Description Get the current user’s permissions for the given app. If a user can see an app, then they can see its basic data. Only admin, read-only admins, and space developers can read sensitive data.
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} AppPermission
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/permissions [GET]
func getAppPermissions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/permissions", nil, "GET", w, r)
	if rBodyResult {
		var final AppPermission
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin Space Developer
// @Summary Set current droplet
// @Description Set the current droplet for an app. The current droplet is the droplet that the app will use when running.
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Param AppSetDroplet body AppSetDroplet true "App Set Droplet"
// @Success 200 {object} AppDropletAssociation
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/relationships/current_droplet [PATCH]
func setAppDroplet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody AppSetDroplet
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/current_droplet", reqBody, "PATCH", w, r)
	if rBodyResult {
		var final AppDropletAssociation
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get SSH enabled for an app
// @Description Returns if an application’s runtime environment will accept ssh connections. If ssh is disabled, the reason field will describe whether it is disabled globally, at the space level, or at the app level.
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} AppSSH
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/ssh_enabled [GET]
func getAppSSH(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/ssh_enabled", nil, "GET", w, r)
	if rBodyResult {
		var final AppSSH
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Start an app
// @Description
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} App
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/actions/start [POST]
func startApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/actions/start", nil, "POST", w, r)
	if rBodyResult {
		var final App
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin Space Developer
// @Summary Stop an app
// @Description
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} App
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/actions/stop [POST]
func stopApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/actions/stop", nil, "POST", w, r)
	if rBodyResult {
		var final App
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin Space Developer
// @Summary Stop an app
// @Description This endpoint will synchronously stop and start an application. Unlike the start and stop actions, this endpoint will error if the app is not successfully stopped in the runtime. For restarting applications without downtime, see the deployments resource.
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} App
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/actions/restart [POST]
func restartApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/actions/restart", nil, "POST", w, r)
	if rBodyResult {
		var final App
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Stop an app
// @Description Update the environment variables associated with the given app. The variables given in the request will be merged with the existing app environment variables. Any requested variables with a value of null will be removed from the app. Environment variable names may not start with VCAP_. PORT is not a valid environment variable.
// @Tags Apps
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Param App Env body AppEnvVar true "App Env"
// @Success 200 {object} AppEnvVariable
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/environment_variables [PATCH]
func setAppEnv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody AppEnvVar
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/environment_variables", reqBody, "PATCH", w, r)
	if rBodyResult {
		var final AppEnvVariable
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
