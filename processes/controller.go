package processes

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var uris = "processes"

func ProcessHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getProcess).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris, getProcess).Methods("GET")
}

// Permitted roles "Admin", "Admin" Read-Only, "Global Auditor", "Org Manager", "Space Auditor", "Space Developer", "Space Manager"
func getProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Process
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	}
}
