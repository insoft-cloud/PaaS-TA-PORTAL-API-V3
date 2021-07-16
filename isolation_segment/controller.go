package isolation_segments

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "isolation_segments"

func IsolationSegmentsHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createIsolationSegment).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getIsolationSegment).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getIsolationSegments).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/organizations", getOrganizationsRelationship).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/spaces", spacesRelationship).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updateIsolationSegment).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deleteIsolationSegment).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/organizations", entitleOrganizationsIsolationSegment).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/relationships/organizations/{org_guid}", revokeOrganizationsIsolationSegment).Methods("DELETE")

}

// @Description Permitted Roles 'Admin'
// @Summary Create an isolation segment
// @Description
// @Tags IsolationSegment
// @Produce  json
// @Security ApiKeyAuth
// @Param CreateIsolationSegment body CreateIsolationSegment true "Create Isolation Segment"
// @Success 201 {object} IsolationSegment
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /isolation_segments [POST]
func createIsolationSegment(w http.ResponseWriter, r *http.Request) {
	var pBody CreateIsolationSegment
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
		var final IsolationSegment
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary Get an isolation segment
// @Description
// @Tags IsolationSegment
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "IsolationSegment Guid"
// @Success 200 {object} IsolationSegment
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /isolation_segments/{guid} [GET]
func getIsolationSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final IsolationSegment
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'All'
// @Summary List isolation segments
// @Description Retrieves all isolation segments to which the user has access. For admin, this is all the isolation segments in the system. For an org manager, this is the isolation segments in the allowed list for any organization to which the user belongs. For any other user, this is the isolation segments assigned to any spaces to which the user has access.
// @Tags IsolationSegment
// @Produce  json
// @Security ApiKeyAuth
// @Param guids query []string false "Comma-delimited list of isolation segment guids to filter by" collectionFormat(csv)
// @Param names query []string false "Comma-delimited list of isolation segment names to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name, state"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param lifecycle_type query string false "Lifecycle type to filter by; valid values are buildpack, docker"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} IsolationSegments
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /isolation_segments [GET]
func getIsolationSegments(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final IsolationSegments
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles All Roles
// @Summary List organizations relationship
// @Description This endpoint lists the organizations entitled for the isolation segment. For an Admin, this will list all entitled organizations in the system. For any other user, this will list only the entitled organizations to which the user belongs.
// @Tags IsolationSegment
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "IsolationSegment Guid"
// @Success 200 {object} OrganizationsRelationship
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /isolation_segments/{guid}/relationships/organizations [GET]
func getOrganizationsRelationship(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/organizations", nil, "GET", w, r)
	if rBodyResult {
		var final OrganizationsRelationship
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin, Admin Read-Only Global Auditor Org Manager Space Auditor Space Developer Space Manager'
// @Summary List spaces relationship
// @Description This endpoint lists the spaces to which the isolation segment is assigned. For an Admin, this will list all associated spaces in the system. For an org manager, this will list only those associated spaces belonging to orgs for which the user is a manager. For any other user, this will list only those associated spaces to which the user has access.
// @Tags IsolationSegment
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "IsolationSegment Guid"
// @Success 200 {object} SpacesRelationship
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /isolation_segments/{guid}/relationships/spaces [GET]
func spacesRelationship(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/spaces", nil, "GET", w, r)
	if rBodyResult {
		var final SpacesRelationship
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin
// @Summary Update an isolation segment
// @Description
// @Tags IsolationSegment
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "IsolationSegment Guid"
// @Param UpdateIsolationSegment body UpdateIsolationSegment true "Update IsolationSegment"
// @Success 200 {object} IsolationSegment
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /isolation_segments/{guid} [PATCH]
func updateIsolationSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody UpdateIsolationSegment
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
		var final IsolationSegment
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin
// @Summary Delete an isolation segment
// @Description An isolation segment cannot be deleted if it is entitled to any organization.
// @Tags IsolationSegment
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "IsolationSegment Guid"
// @Success 204  {object} string "No Content"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /isolation_segments/{guid} [DELETE]
func deleteIsolationSegment(w http.ResponseWriter, r *http.Request) {
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

// @Description Permitted Roles 'Admin'
// @Summary Entitle organizations for an isolation segment
// @Description This endpoint entitles the specified organizations for the isolation segment. In the case where the specified isolation segment is the system-wide shared segment, and if an organization is not already entitled for any other isolation segment, then the shared isolation segment automatically gets assigned as the default for that organization.
// @Tags IsolationSegment
// @Produce  json
// @Security ApiKeyAuth
// @Param EntitleOrganizationsIsolationSegment body EntitleOrganizationsIsolationSegment true "Entitle Organizations IsolationSegment"
// @Param guid path string true "IsolationSegment Guid"
// @Success 200 {object} OrganizationsRelationship
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /isolation_segments/{guid}/relationships/organizations [POST]
func entitleOrganizationsIsolationSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]

	var pBody EntitleOrganizationsIsolationSegment
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}

	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/organizations", reqBody, "POST", w, r)
	if rBodyResult {
		var final OrganizationsRelationship
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles Admin
// @Summary Revoke entitlement to isolation segment for an organization
// @Description This endpoint revokes the entitlement for the specified organization to the isolation segment. If the isolation segment is assigned to a space within an organization, the entitlement cannot be revoked. If the isolation segment is the organization’s default, the entitlement cannot be revoked.
// @Tags IsolationSegment
// @Produce  json
// @Security ApiKeyAuth
// @Param org_guid path string true "Organization Guid"
// @Param guid path string true "IsolationSegment Guid"
// @Success 204  {object} string "No Content"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /isolation_segments/{guid}/relationships/organizations/{org_guid} [DELETE]
func revokeOrganizationsIsolationSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	org_guid := vars["org_guid"]

	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/relationships/organizations/"+org_guid, nil, "DELETE", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
