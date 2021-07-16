package spaces

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "spaces"

func ServiceRouteBindingHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createSpace).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getSpace).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getSpaces).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateSpace).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteSpace).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/isolation_segment", assignedIsolationSegment).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/isolation_segment", manageIsolationSegment).Methods("PATCH")
}

// @Description Permitted Roles 'Admin, Org Manager'
// @Summary Create a space
// @Description
// @Tags Spaces
// @Produce  json
// @Security ApiKeyAuth
// @Param Space body Space true "Create Space"
// @Success 201 {object} Space
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces [POST]
func createSpace(w http.ResponseWriter, r *http.Request) {
	var pBody CreateSpace
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
		var final Space
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get a space
// @Description This endpoint retrieves the specified space object.
// @Tags Spaces
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Space Guid"
// @Param include query []string false "Optionally include additional related resources in the response; valid value is organization" collectionFormat(multi)
// @Success 200 {object} Space
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces/{guid} [GET]
func getSpace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final Space
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted Roles 'All'
// @Summary List spaces
// @Description Retrieve all spaces the user has access to.
// @Tags Spaces
// @Produce  json
// @Security ApiKeyAuth
// @Param names query []string false "Comma-delimited list of space names to filter by" collectionFormat(csv)
// @Param guids query []string false "Comma-delimited list of space guids to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param include query []string false "Optionally include a list of unique related resources in the response; valid values are space and spaceorganization" collectionFormat(multi)
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} SpaceList
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces [GET]
func getSpaces(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final SpaceList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted Roles 'Admin Org Manager Space Manager'
// @Summary Update a space
// @Description
// @Tags Spaces
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Space Guid"
// @Param UpdateSpace body UpdateSpace true "Update Space"
// @Success 200 {object} Space
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces/{guid} [PATCH]
func updateSpace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdateSpace
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
		var final Space
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

// @Description Permitted Roles 'Admin Org Manager'
// @Summary Delete a space
// @Description When a space is deleted, the user roles associated with the space will be deleted.
// @Tags Spaces
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Space Guid"
// @Success 202 {object} string "ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces/{guid} [DELETE]
func deleteSpace(w http.ResponseWriter, r *http.Request) {
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

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get assigned isolation segment
// @Description
// @Tags Spaces
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Space Guid"
// @Success 200 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces/{guid}/relationships/isolation_segment [GET]
func assignedIsolationSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/isolation_segment", nil, "GET", w, r)
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

// @Description Permitted Roles 'Admin Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Manage isolation segment
// @Description This endpoint assigns an isolation segment to the space. The isolation segment must be entitled to the space’s parent organization.
// @Description Apps will not run in the newly assigned isolation segment until they are restarted.
// @Tags Spaces
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "Space Guid"
// @Param IsolationSegment body IsolationSegment true "IsolationSegment"
// @Success 200 {object} object
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces/{guid}/relationships/isolation_segment [PATCH]
func manageIsolationSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody IsolationSegment
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/isolation_segment", reqBody, "PATCH", w, r)
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
