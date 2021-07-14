package roles

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
	Valid role types
	organization_user
	organization_auditor
	organization_manager
	organization_billing_manager
	space_auditor
	space_developer
	space_manager
	space_application_supporter - This role is under active development and is not supported
*/

var uris = "roles"

func RoleHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createRole).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getRole).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getRoles).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, deleteRole).Methods("DELETE")
}

//  @Description Permitted roles
// Role	Notes
// Admin
// Org Manager	Can create roles in managed organizations and spaces within those organizations;
//				can also create roles for users outside of managed organizations
//				when set_roles_by_username feature_flag is enabled; this requires identifying users by username and origin
// Space Manager	Can create roles in managed spaces for users in their org
// @Summary Create a role
// @Description This endpoint creates a new role for a user in an organization or space.
// @Description To create an organization role you must be an admin or organization manager in the organization associated with the role.
// @Description To create a space role you must be an admin, an organization manager in the parent organization of the space associated with the role, or a space manager in the space associated with the role.
// @Description For a user to be assigned a space role, the user must already have an organization role in the parent organization.
// @Description If the associated user is valid but does not exist in Cloud Controller’s database, a user resource will be created automatically.
// @Tags Roles
// @Produce json
// @Security ApiKeyAuth
// @Param type body string true "Role to create; see valid role types"
// @Param relationships.user body string true "A relationship to a user; the user can be defined by either a guid or, if the set_roles_by_username feature_flag is enabled, a username (with the option of including an origin to disambiguate it)"
// @Param relationships.organization body string true "A relationship to an organization; required only when creating an organization role"
// @Param relationships.space body string true "A relationship to a space; required only when creating a space role"
// @Success 200 {object} Role
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /roles [POST]
func createRole(w http.ResponseWriter, r *http.Request) {
	var pBody CreateRole
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
		var final Role
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//  @Description Permitted roles
// Role	Notes
// Admin
// Admin Read-Only
// Global Auditor
// Org Manager	Can see roles in managed organizations or spaces in those organizations
// Org Auditor	Can only see organization roles in audited organizations
// Org Billing Manager	Can only see organization roles in billing-managed organizations
// Space Auditor	Can see roles in audited spaces or parent organizations
// Space Application Supporter (under development)	Can see roles in supported spaces or parent organizations
// Space Developer	Can see roles in developed spaces or parent organizations
// Space Manager	Can see roles in managed spaces or parent organizations
// @Summary Get a role
// @Description This endpoint gets an individual role resource.
// @Tags Roles
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "role guid"
// @Param include path string false "Optionally include a list of unique related resources in the response; valid values are user, space, and organization"
// @Success 200 {object} Role
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /roles/{guid} [GET]
func getRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Role
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//  @Description Permitted roles "All Roles"
// @Summary List roles
// @Description This endpoint lists roles that the user has access to.
// @Tags Roles
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} Role
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /roles [GET]
func getRoles(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final RoleList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}

//  @Description Permitted roles
// Role	Notes
// Admin
// Org Manager	Can delete roles in managed organizations or spaces in those organizations
// Space Manager	Can delete roles in managed spaces
// @Summary List roles
// @Description This endpoint lists roles that the user has access to.
// @Tags Roles
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "role guid"
// @Success 200 {object} Role
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /roles/{guid} [DELETE]
func deleteRole(w http.ResponseWriter, r *http.Request) {
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
