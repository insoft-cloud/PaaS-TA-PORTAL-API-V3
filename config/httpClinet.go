package config

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var Client http.Client

func ClientSetting() {
	// 인증서 추가 고민....
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	Client = http.Client{Transport: tr}
}

func Curl(url string, tbody []byte, method string, w http.ResponseWriter, r *http.Request) (interface{}, bool) {
	req, _ := http.NewRequest(method, GetDomainConfig()+url, bytes.NewBuffer(tbody))
	req.Header.Set("Authorization", r.Header.Get("cf-Authorization"))
	req.Header.Set("Content-type", "application/json")
	w.Header().Set("content-type", "application/json")
	res, err := Client.Do(req)
	w.WriteHeader(res.StatusCode)
	//token := r.Header.Values("cf-Authorization") //cf token
	if err != nil {
		rErrs := &Errors{Code: 500, Detail: err.Error(), Title: "Portal API Error"}
		Errorlog.Error(method, "',", rErrs, ",", url)
		return rErrs, false
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		rErrs := &Errors{Code: 500, Detail: err.Error(), Title: "Portal API Error"}
		Errorlog.Error(method, "',", rErrs, ",", url)
		return rErrs, false
	} else if res.StatusCode > 400 {
		var final Error
		json.Unmarshal(body, &final)
		Errorlog.Error(method, "',", final, ",", url)
		return final, false
	}
	now := time.Now()
	jsonString := string(body)
	replace := strings.ReplaceAll(jsonString, "\n", "")
	fmt.Println(now, strings.ReplaceAll(replace, "    ", ""))
	// replaceBody := strings.ReplaceAll(replace, "\"", "")
	Infolog.Info(r.Proto, "", "[", r.URL, "]")
	return body, true
}

func ManifestCurl(url string, tbody []byte, method string, w http.ResponseWriter, r *http.Request) (interface{}, bool) {
	req, _ := http.NewRequest(method, GetDomainConfig()+url, bytes.NewBuffer(tbody))
	req.Header.Set("Authorization", r.Header.Get("cf-Authorization"))
	req.Header.Set("Content-type", "application/x-yaml")
	w.Header().Set("content-type", "application/json")
	res, err := Client.Do(req)
	w.WriteHeader(res.StatusCode)
	if err != nil {
		rErrs := &Errors{Code: 500, Detail: err.Error(), Title: "Portal API Error"}
		Errorlog.Println(method, "',", rErrs, ",", url)
		return rErrs, false
	}
	body, err := ioutil.ReadAll(res.Body)
	Infolog.Info(r.Proto, "", "[", r.URL, "]")

	if err != nil {
		rErrs := &Errors{Code: 500, Detail: err.Error(), Title: "Portal API Error"}
		Errorlog.Println(method, "',", rErrs, ",", url)
		return rErrs, false
	} else if res.StatusCode > 400 {
		var final Error
		json.Unmarshal(body, &final)
		Errorlog.Println(method, "',", final, ",", url)
		return final, false
	}
	if r.Method == http.MethodGet {
		return string(body), true
	} else {
		return body, true
	}
}

func FileCurl(key string, url string, method string, w http.ResponseWriter, r *http.Request) (interface{}, bool) {
	uploaded, handler, err := r.FormFile(key)
	if err != nil {
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
		Errorlog.Error(method, "',", final, ",", url)
		return final, false
	}
	defer uploaded.Close()
	file, _ := os.Create(handler.Filename)
	_, err = io.Copy(file, uploaded)
	if err != nil {
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
		Errorlog.Error(method, "',", final, ",", url)
		file.Close()
		os.Remove(handler.Filename)
		return final, false
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile(key, filepath.Base(handler.Filename))
	Refile, _ := os.Open(handler.Filename)
	_, err = io.Copy(part, Refile)
	if err != nil {
		Refile.Close()
		writer.Close()
		os.Remove(handler.Filename)
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
		Errorlog.Error(method, "',", final, ",", url)
		return final, false
	}
	Refile.Close()
	writer.Close()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
		Errorlog.Error(method, "',", final, ",", url)
		os.Remove(handler.Filename)
		return final, false
	}
	req.Header.Set("Authorization", r.Header.Get("cf-Authorization"))
	req.Header.Add("Content-type", writer.FormDataContentType())
	res, err := Client.Do(req)
	if err != nil {
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
		os.Remove(handler.Filename)
		Errorlog.Error(method, "',", final, ",", url)
		return final, false
	}
	tbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
		os.Remove(handler.Filename)
		Errorlog.Error(method, "',", final, ",", url)
		return final, false
	} else if res.StatusCode > 400 {
		w.WriteHeader(res.StatusCode)
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
		os.Remove(handler.Filename)
		Errorlog.Error(method, "',", final, ",", url)
		return final, false
	}
	os.Remove(handler.Filename)
	Infolog.Info(r.Proto, "", "[", r.URL, "]")
	return tbody, true
}

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler

	return handleCORS(handler)
}
