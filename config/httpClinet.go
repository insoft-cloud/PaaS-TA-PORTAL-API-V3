package config

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
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
	res, err := Client.Do(req)
	if err != nil {
		return err, false
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err, false
	} else if res.StatusCode > 400 {
		w.WriteHeader(res.StatusCode)
		return body, false
	}
	return body, true

}
