package resource_matches

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var uris = "resource_matches"

func ResourceMatchesHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createResourceMatch).Methods("POST")
}

//  @Description Permitted roles 'All Roles'
// 201 결과 안난오고 200에 결과 값
// @Summary Create a resource match
// @Description This endpoint returns a list of cached resources from the input list.
// @Tags Resource Matches
// @Produce json
// @Security ApiKeyAuth
// @Param resources body ResourceMatch true "List of resources to check for in the resource cache"
// @Success 200 {object} ResourceMatch
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /resource_matches [POST]
func createResourceMatch(w http.ResponseWriter, r *http.Request) {
	var pBody ResourceMatch
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
		var final ResourceMatch
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
