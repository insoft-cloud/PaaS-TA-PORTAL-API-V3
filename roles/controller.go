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

// @Description Permitted roles 'Admin,
// @Description Org Manager	Can create roles in managed organizations and spaces within those organizations;
// @Description	can also create roles for users outside of managed organizations
// @Description when set_roles_by_username feature_flag is enabled; this requires identifying users by username and origin,
// @Description Space Manager Can create roles in managed spaces for users in their org'
// @Summary Create a role
// @Description This endpoint creates a new role for a user in an organization or space.
// @Description To create an organization role you must be an admin or organization manager in the organization associated with the role.
// @Description To create a space role you must be an admin, an organization manager in the parent organization of the space associated with the role, or a space manager in the space associated with the role.
// @Description For a user to be assigned a space role, the user must already have an organization role in the parent organization.
// @Description If the associated user is valid but does not exist in Cloud Controller’s database, a user resource will be created automatically.
// @Tags Roles
// @Produce json
// @Security ApiKeyAuth
// @Param CreateRoles body CreateRole true "Create Roles"
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

// @Description Permitted roles 'Admin, Admin Read-Only, Global Auditor, Org Manager	Can see roles in managed organizations or spaces in those organizations,
// @Description Org Auditor	Can only see organization roles in audited organizations, Org Billing Manager Can only see organization roles in billing-managed organizations,
// @Description Space Auditor Can see roles in audited spaces or parent organizations, Space Application Supporter (under development) Can see roles in supported spaces or parent organizations
// @Description Space Developer	Can see roles in developed spaces or parent organizations, Space Manager	Can see roles in managed spaces or parent organizations'
// @Summary Get a role
// @Description This endpoint gets an individual role resource.
// @Tags Roles
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "role guid"
// @Param include query []string false "Optionally include a list of unique related resources in the response; valid values are user, space, and organization"
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

// @Description Permitted roles 'All Roles'
// @Summary List roles
// @Description This endpoint lists roles that the user has access to.
// @Tags Roles
// @Produce json
// @Security ApiKeyAuth
// @Param guids query []string false "Comma-delimited list of role guids to filter by" collectionFormat(csv)
// @Param types query []string false "Comma-delimited list of role types to filter by" collectionFormat(csv)
// @Param space_guids query []string false "Comma-delimited list of space guids to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by; defaults to ascending. Prepend with - to sort descending. Valid values are created_at, updated_at"
// @Param include query []string false "Optionally include a list of unique related resources in the response; valid values are user, space, and organization"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
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

// @Description Permitted roles 'Admin, Org Manager Can delete roles in managed organizations or spaces in those organizations,
// @Description Space Manager Can delete roles in managed spaces'
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
