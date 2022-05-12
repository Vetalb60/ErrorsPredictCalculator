// Package requests
//
//	________upload_file_req.go________
//
//	Upload file to database request.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	_FORM_TYPE = "file"
)

func (req *Requests) UploadFile(file_name string) (string, error) {
	str, err := SendPostRequest("http://"+req.GetCustomized().Host+"/upload", file_name, _FORM_TYPE)
	return str, err
}

func SendPostRequest(url string, filename string, filetype string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(filetype, filepath.Base(file.Name()))
	if err != nil {
		return "", err
	}

	io.Copy(part, file)
	writer.Close()

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	request.Close = true

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var dat map[string]interface{}

	err = json.Unmarshal(content, &dat)
	if err != nil {
		return "", err
	}

	return dat["message"].(string), nil
}
