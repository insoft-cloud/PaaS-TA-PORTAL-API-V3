package config

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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

func FileCurl3(key string, url string, method string, w http.ResponseWriter, r *http.Request) (interface{}, bool) {
	uploaded, handler, err := r.FormFile(key)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
		final := ErrorMessage("BuildPack Upload Error :: "+err.Error(), 500, w)
		return final, false
	}
	fmt.Println(123213)
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	fw, err := writer.CreateFormFile(key, handler.Filename)
	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
		final := ErrorMessage("BuildPack Upload Error :: "+err.Error(), 500, w)
		return final, false
	}
	fmt.Println(456456)
	_, err = io.Copy(fw, uploaded)
	if err != nil {
		fmt.Println(3)
		fmt.Println(err)
		final := ErrorMessage("BuildPack Upload Error :: "+err.Error(), 500, w)
		return final, false
	}
	writer.Close()
	req, err := http.NewRequest(method, GetDomainConfig()+url, bytes.NewReader(buffer.Bytes()))
	if err != nil {
		fmt.Println(66666)
		fmt.Println(err)
		return err, false
	}
	req.Header.Set("Authorization", r.Header.Get("cf-Authorization"))
	req.Header.Add("Content-type", "multipart/form-data")
	fmt.Println(req)
	res, err := Client.Do(req)
	if err != nil {
		fmt.Println(4)
		fmt.Println(err)
		return err, false
	}
	fmt.Println(res)
	fmt.Println(789789)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err, false
	} else if res.StatusCode > 400 {
		fmt.Println(err)
		w.WriteHeader(res.StatusCode)
		return body, false
	}
	fmt.Println(5)
	return body, true
}

func FileCurl2(key string, url string, method string, w http.ResponseWriter, r *http.Request) (interface{}, bool) {
	reader, err := r.MultipartReader()
	var filename string
	var filer *os.File
	//에러가 발생하면 던진다
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, false
	}

	//for로 복수 파일이 있는 경우에 모든 파일이 끝날때까지 읽는다
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		//파일 명이 없는 경우는 skip
		if part.FileName() == "" {
			continue
		}

		//uploadedfile 디렉토리에 받았던 파일 명으로 파일을 만든다
		filer, err = os.Create(part.FileName())
		filename = part.FileName()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			filer.Close()
			return nil, false
		}

		//만든 파일에 읽었던 파일 내용을 모두 복사
		_, err = io.Copy(filer, part)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			filer.Close()
			return nil, false
		}
	}
	filer.Close()
	tt, _ := os.Open(filename)
	defer tt.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile(key, filepath.Base(filename))
	io.Copy(part, tt)
	writer.Close()
	req, err := http.NewRequest(method, GetDomainConfig()+url, body)

	if err != nil {
		fmt.Println(66666)
		fmt.Println(err)
		return err, false
	}
	req.Header.Set("Authorization", r.Header.Get("cf-Authorization"))
	req.Header.Add("Content-type", writer.FormDataContentType())

	if err != nil {
		fmt.Println(err)
		return err, false
	}
	res, err := Client.Do(req)
	if err != nil {
		os.Remove(filename)
		return err, false
	}
	bodyr, err := ioutil.ReadAll(res.Body)
	if err != nil {
		os.Remove(filename)
		return err, false
	} else if res.StatusCode > 400 {
		os.Remove(filename)
		w.WriteHeader(res.StatusCode)
		return bodyr, false
	}
	fmt.Println("파이널")
	os.Remove("test.zip")
	return nil, false
}

func FileCurl(key string, url string, method string, w http.ResponseWriter, r *http.Request) (interface{}, bool) {
	uploaded, handler, err := r.FormFile(key)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
		return final, false
	}
	defer uploaded.Close()
	file, _ := os.Create(handler.Filename)
	_, err = io.Copy(file, uploaded)
	if err != nil {
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
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
		return final, false
	}
	Refile.Close()
	writer.Close()
	req, err := http.NewRequest(method, GetDomainConfig()+url, body)
	if err != nil {
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
		os.Remove(handler.Filename)
		return final, false
	}
	req.Header.Set("Authorization", r.Header.Get("cf-Authorization"))
	req.Header.Add("Content-type", writer.FormDataContentType())
	res, err := Client.Do(req)
	if err != nil {
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
		os.Remove(handler.Filename)
		return final, false
	}
	tbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
		os.Remove(handler.Filename)
		return final, false
	} else if res.StatusCode > 400 {
		w.WriteHeader(res.StatusCode)
		final := ErrorMessage("File Upload Error :: "+err.Error(), 500, w)
		os.Remove(handler.Filename)
		return final, false
	}
	os.Remove(handler.Filename)
	return tbody, true
}
