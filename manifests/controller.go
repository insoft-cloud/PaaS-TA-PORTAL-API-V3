package manifests

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var uris = "spaces"

func ManifestsHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/actions/apply_manifest", applyManifest).Methods("POST")
	myRouter.HandleFunc("/v3/apps/{guid}/manifest", generateManifest).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/manifest_diff", createManifestDiff).Methods("POST")

}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Apply a manifest to a space
// @Description Apply changes specified in a manifest to the named apps and their underlying processes. The apps must reside in the space. These changes are additive and will not modify any unspecified properties or remove any existing environment variables, routes, or services.
// @Description Apply manifest will only trigger an immediate update for the “disk_quota”, “instances”, and “memory” properties. All other properties require an app restart to take effect.
// @Tags Manifests
// @Produce  json
// @Security ApiKeyAuth
// @Param binary body string true "manifest.yml" format(binary)
// @Param guid path string true "Space Guid"
// @Success 202 {object} string "ok"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces/{guid}/actions/apply_manifest [POST]
func applyManifest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.ManifestCurl("/v3/"+uris+"/"+guid+"/actions/apply_manifest", reqBody, "POST", w, r)
	if rBodyResult {
		json.NewEncoder(w).Encode("OK")
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Admin Read-Only Space Developer'
// @Summary Generate a manifest for an app
// @Description Generate a manifest for an app and its underlying processes.
// @Tags Manifests
// @Produce  json
// @Security ApiKeyAuth
// @Param guid path string true "App Guid"
// @Success 200 {object} string "---"
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces/{guid}/manifest [GET]
func generateManifest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.ManifestCurl("/v3/apps/"+guid+"/manifest", nil, "GET", w, r)
	if rBodyResult {
		fmt.Println(rBody)
		json.NewEncoder(w).Encode(rBody)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// @Description Permitted Roles 'Admin Space Developer'
// @Summary Generate a manifest for an app
// @Description This endpoint returns a JSON representation of the difference between the provided manifest and the current state of a space.
// @Description Currently, this endpoint can only diff version 1 manifests.
// @Tags Manifests
// @Produce  json
// @Security ApiKeyAuth
// @Param binary body string true "manifest.yml" format(binary)
// @Param guid path string true "Space Guid"
// @Success 202 {object} object Diff
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /spaces/{guid}/manifest [GET]
func createManifestDiff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.ManifestCurl("/v3/spaces/"+guid+"/manifest_diff", reqBody, "POST", w, r)
	if rBodyResult {
		var final Diff
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
