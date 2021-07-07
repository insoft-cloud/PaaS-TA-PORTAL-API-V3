package packages

import (
	"PAAS-TA-PORTAL-V3/config"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

var uris = "packages"

func PackagesHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/v3/"+uris, createPackage).Methods("POST")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", getPackage).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris, getPackages).Methods("GET")
	myRouter.HandleFunc("/v3/apps/{guid}/"+uris, getPackagesForAnApp).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", updatePackage).Methods("PATCH")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}", deletePackage).Methods("DELETE")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/download", downloadPackage).Methods("GET")
	myRouter.HandleFunc("/v3/"+uris+"/{guid}/upload", uploadPackage).Methods("POST")
}

// Permitted roles "Admin", "Space Developer"
func createPackage(w http.ResponseWriter, r *http.Request) {
	var pBody CreatePackage
	if r.URL.Query().Get("source_guid") != "" {
		copyPackage(w, r)
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
		var final Package
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// Permitted roles
//"Admin", "Admin" Read-Only, "Global Auditor", "Org Manager", "Space Auditor", "Space Developer", "Space Manager"
func getPackage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, nil, "GET", w, r)
	if rBodyResult {
		var final Package
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

func getPackages(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final PackageList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

//Permitted roles
//"Admin", "Admin" Read-Only, "Global Auditor", "Org Manager", "Space Auditor", "Space Developer", "Space Manager"
func getPackagesForAnApp(w http.ResponseWriter, r *http.Request) {
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/apps/"+guid+"/"+uris+"?"+query, nil, "GET", w, r)
	if rBodyResult {
		var final PackageList
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// Permitted roles "Admin", "Space Developer"
// parameter를 못받아서 수정(update)이 안됨 metadata
func updatePackage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	var pBody UpdatePackage
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid, reqBody, "PATCH", w, r)
	if rBodyResult {
		var final Package
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// Permitted Roles "Admin", "Space Developer"
func deletePackage(w http.ResponseWriter, r *http.Request) {
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

// Permitted roles "Admin" "Space Developer"
func copyPackage(w http.ResponseWriter, r *http.Request) {
	var pBody CopyPackage
	query, _ := url.QueryUnescape(r.URL.Query().Encode())
	vResultI, vResultB := config.Validation(r, &pBody)
	if !vResultB {
		json.NewEncoder(w).Encode(vResultI)
		return
	}
	//호출
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, pBody)
	reqBody, _ = json.Marshal(pBody)
	rBody, rBodyResult := config.Curl("/v3/"+uris+"?"+query, reqBody, "POST", w, r)
	if rBodyResult {
		var final Package
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}

// PackageDownload pass status 302 안나오고 200으로 나옴
// Permitted roles "Admin", "Space Developer"
func downloadPackage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.Curl("/v3/"+uris+"/"+guid+"/download", nil, "GET", w, r)
	if rBodyResult {
		var final interface{}
		json.Unmarshal(rBody.([]byte), &final)
		fmt.Print(final)

		json.NewEncoder(w).Encode(final)
	} else {
		fmt.Print(rBody)
		//	var final interface{}
		//		json.Unmarshal(rBody.([]byte), &final)
		//		json.NewEncoder(w).Encode(final)
	}
}

// Permitted roles "Admin", "Space Developer"
func uploadPackage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	rBody, rBodyResult := config.FileCurl("bits", "/v3/"+uris+"/"+guid+"/upload", "POST", w, r)
	if rBodyResult {
		var final Package
		json.Unmarshal(rBody.([]byte), &final)
		json.NewEncoder(w).Encode(final)
	} else {
		json.NewEncoder(w).Encode(rBody)
	}
}
