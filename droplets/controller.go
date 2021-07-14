package droplets

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "droplets"

func DropletHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createDroplets).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getDroplet).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getDroplets).Methods("GET") //List droplet
	myRouter.HandleFunc("/v3/packages/{guid}/"+uris, getDropletsPackages).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris, getDropletsApp).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateDroplet).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteDroplet).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/download", downloadDroplet).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/upload", uploadDroplet).Methods("POST")

}

// @Description Permitted Roles 'Admin, SpaceDeveloper'
// @Summary Create a droplet
// @Description This endpoint is only for creating a droplet without a package. To create a droplet based on a package, see Create a build.
// @Tags Droplets
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateDroplet body CreateDroplet true "Create Droplet"
// @Success 200 {object} Droplet
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /droplets [POST]
func createDroplets(w http.ResponseWriter, r *http.Request) {
	var pBody CreateDroplet
	if r.URL.Query().Get("source_guid") != "" {
		copyDroplet(w, r)
		return
	}
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
		var final Droplet
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get a droplet
// @Description
// @Tags Droplets
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Droplet Guid"
// @Success 200 {object} Droplet
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /droplets/{guid} [GET]
func getDroplet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Droplet
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All Roles'
// @Summary List droplets
// @Description
// @Tags Droplets
// @Produce  json
// @Security ApiKeyAuth
// @Param guids query []string false "Comma-delimited list of droplet to filter by" collectionFormat(csv)
// @Param states query []string false "Comma-delimited list of droplet states to filter by" collectionFormat(csv)
// @Param app_guids query []string false "Comma-delimited list of app guids to filter by" collectionFormat(csv)
// @Param space_guids query []string false "Comma-delimited list of space guids to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param guid path string true "Organization Guid"
// @Success 200 {object} DropletList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /droplets [GET]
func getDroplets(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final DropletList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}

}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary List droplets for a package
// @Description Retrieve a list of droplets belonging to a package.
// @Tags Droplets
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Packages Guid"
// @Param guids query []string false "Comma-delimited list of droplet to filter by" collectionFormat(csv)
// @Param states query []string false "Comma-delimited list of droplet states to filter by" collectionFormat(csv)
// @Param app_guids query []string false "Comma-delimited list of app guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Success 200 {object} DropletList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /packages/{guid}/droplets [GET]
func getDropletsPackages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/packages/"+guid+"/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final DropletList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary List droplets for an app
// @Description Retrieve a list of droplets belonging to an app.
// @Tags Droplets
// @Produce  json
// @Security ApiKeyAuth
// @Param guids query []string false "Comma-delimited list of droplet to filter by" collectionFormat(csv)
// @Param states query []string false "Comma-delimited list of droplet states to filter by" collectionFormat(csv)
// @Param current query []string false "Comma-delimited list of app guids to filter by" collectionFormat(csv)
// @Param space_guids query boolean false "If true, only include the droplet currently assigned to the app"
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param guid path string true "Droplet Guid"
// @Success 200 {object} DropletList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/droplets [GET]
func getDropletsApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final DropletList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Org Manager Space Developer Space Manager'
// @Summary Update a droplet
// @Description
// @Tags Droplets
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Droplet Guid"
// @Param UpdateDroplet body UpdateDroplet true "Update Droplet"
// @Success 200 {object} Droplet
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /droplets/{guid} [PATCH]
func updateDroplet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateDroplet
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
		var final Droplet
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin  Space Developer
// @Summary Delete a droplet
// @Description
// @Tags Droplets
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Droplet Guid"
// @Success 202 {object} string	"ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /droplets/{guid} [DELETE]
func deleteDroplet(w http.ResponseWriter, r *http.Request) {
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

// @Description Permitted Roles Admin  Space Developer
// @Summary Copy a droplet
// @Description Copy a droplet to a different app. The copied droplet excludes the environment variables listed on the source droplet.
// @Tags Droplets
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Droplet Guid"
// @Param CopyDroplet body CopyDroplet true "CopyDroplet"
// @Param source_guid query string true "Source guid of the droplet to be copied"
// @Success 200 {object} Droplet
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /droplets/{guid} [POST]
func copyDroplet(w http.ResponseWriter, r *http.Request) {
	var pBody CopyDroplet
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?source_guid", reqBody, "POST", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager
// @Summary Download droplet bits
// @Description Download a gzip compressed tarball file containing a Cloud Foundry compatible droplet. When using a remote blobstore, such as AWS, the response is a redirect to the actual location of the bits. If the client is automatically following redirects, then the OAuth token that was used to communicate with Cloud Controller will be relayed on the new redirect request. Some blobstores may reject the request in that case. Clients may need to follow the redirect without including the OAuth token.
// @Tags Droplets
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Droplet Guid"
// @Success 302 {object} string
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /droplets/{guid}/download [GET]
func downloadDroplet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/download", nil, "GET", w, r)
	if rBodyResult {
		var final Droplet
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager
// @Summary Upload droplet bits
// @Description Upload a gzip compressed tarball file containing a Cloud Foundry compatible droplet. The file must be sent as part of a multi-part form.
// @Tags Droplets
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Droplets Guid"
// @Param bits formData file true "A gzip compressed tarball file with .tgz extension containing the droplet bits"
// @Success 200 {object} Droplet
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /droplets/{guid}/upload [POST]
func uploadDroplet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.FileCurl("bits", "/v3/"+uris+"/"+guid+"/upload", "POST", w, r)
	if rBodyResult {
		var final Droplet
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
