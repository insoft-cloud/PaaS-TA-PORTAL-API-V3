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

// @Description Permitted roles 'Admin, Space Developer'
// @Summary Create a package
// @Description
// @Tags Packages
// @Produce json
// @Security ApiKeyAuth
// @Param CreatePackages body CreatePackage true "Create Packages"
// @Success 200 {object} Package
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /packages [POST]
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

// @Description Permitted roles 'Admin, Admin Read-Only, Global Auditor, Org Manager, Space Auditor, Space Developer, Space Manager'
// @Summary Get a package
// @Description
// @Tags Packages
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "package guid"
// @Success 200 {object} Package
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /packages/{guid} [GET]
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

// @Description Permitted roles 'All Roles'
// @Summary List packages
// @Description Retrieve all packages the user has access to.
// @Tags Packages
// @Produce json
// @Security ApiKeyAuth
// @Param guids query []string false "Comma-delimited list of package guids to filter by" collectionFormat(csv)
// @Param states query []string false "Comma-delimited list of package states to filter by" collectionFormat(csv)
// @Param types query []string false "Comma-delimited list of package types to filter by" collectionFormat(csv)
// @Param app_guids query []string false "Comma-delimited list of app guids to filter by" collectionFormat(csv)
// @Param space_guids query []string false "Comma-delimited list of space guids to filter by" collectionFormat(csv)
// @Param organization_guids query []string false "Comma-delimited list of organization guids to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by; defaults to ascending. Prepend with - to sort descending. Valid values are created_at, updated_at"
// @Param label_selector query string false "A query string containing a list of label selector requirements"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} Package
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /packages [GET]
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

// @Description Permitted roles 'Admin, Admin Read-Only, Global Auditor, Org Manager, Space Auditor, Space Developer, Space Manager'
// @Summary List packages for an app
// @Description Retrieve packages for an app that the user has access to.
// @Tags Packages
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "app guid"
// @Param guids query []string false "Comma-delimited list of package guids to filter by" collectionFormat(csv)
// @Param states query []string false "Comma-delimited list of package states to filter by" collectionFormat(csv)
// @Param types query []string false "Comma-delimited list of package types to filter by" collectionFormat(csv)
// @Param page query integer false "Page to display; valid values are integers >= 1"
// @Param per_page query integer false "Number of results per page; valid values are 1 through 5000"
// @Param order_by query string false "Value to sort by; defaults to ascending. Prepend with - to sort descending. Valid values are created_at, updated_at"
// @Param created_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Param updated_ats query string false "Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators"
// @Success 200 {object} Package
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /apps/{guid}/packages [GET]
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

// @Description Permitted roles 'Admin, Space Developer'
// parameter를 못받아서 수정(update)이 안됨 metadata
// @Summary Update a package
// @Description
// @Tags Packages
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "package guid"
// @Param UpdatePackage body UpdatePackage false "Update Packages"
// @Success 200 {object} Package
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /packages/{guid} [PATCH]
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

// @Description Permitted Roles 'Admin, Space Developer'
// 결과 202 나와야하는데, 200 나옵니다. -> 삭제는 실행됨.
// @Summary Delete a package
// @Description
// @Tags Packages
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "package guid"
// @Success 200 {object} Package
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /packages/{guid} [DELETE]
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

// @Description Permitted roles 'Admin, Space Developer'
// postman으로 테스트 완료(복사 성공), swagger 분기처리 어떻게 해야될지 모르겠음.
// @Summary Copy a package
// @Description This endpoint copies the bits of a source package to a target package.
// @Tags Packages
// @Produce json
// @Security ApiKeyAuth
// @Param source_guid query []string true "GUID of the source package to copy from" collectionFormat(csv)
// @Param CopyPackages body CopyPackage true "Copy Packages"
// @Success 200 {object} Package
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /packages [POST]
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
// @Description Permitted roles 'Admin, Space Developer'
// @Summary Download package bits
// @Description This endpoint downloads the bits of an existing package.
// @Description When using a remote blobstore, such as AWS, the response is a redirect to the actual location of the bits.
// @Description If the client is automatically following redirects, then the OAuth token that was used to communicate with Cloud Controller will be replayed on the new redirect request.
// @Description Some blobstores may reject the request in that case. Clients may need to follow the redirect without including the OAuth token.
// @Tags Packages
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "package guid"
// @Success 200 {object} Package
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /packages/{guid}/download [GET]
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

// @Description Permitted roles 'Admin, Space Developer'
// @Summary Upload package bits
// @Description This upload endpoint takes a multi-part form requests for packages of type bits.
// @Description The request requires either a .zip file uploaded under the bits field or a list of resource match objects under the resources field. These field may be used together.
// @Description The resources field in the request accepts the v2 resources object format.
// @Tags Packages
// @Produce json
// @Security ApiKeyAuth
// @Param guid path string true "package guid"
// @Param bits formData file false "A binary zip file containing the package bits"
// @Param resources formData object false "Fingerprints of the application bits that have previously been pushed to Cloud Foundry, formatted as resource match objects"
// @Success 200 {object} Package
// @Failure 400,404 {object} config.Error
// @Failure 500 {object} config.Error
// @Failure default {object} config.Error
// @Router /packages/{guid}/upload [POST]
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
