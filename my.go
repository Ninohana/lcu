/*
 * Copyright © 2024 Ninohana.
 */

package lol

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const remain = "Go!!!!!!!!!"

type ErrorResponse struct {
	ErrorCode             string      `json:"errorCode"`
	HttpStatus            int         `json:"httpStatus"`
	ImplementationDetails interface{} `json:"implementationDetails"`
	Message               string      `json:"message"`
}
type authSetter interface {
	setAuth(*http.Request)
}

type AuthTransport struct {
	http.RoundTripper
	authSetter
}

func (at AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	at.setAuth(req)
	return at.RoundTripper.RoundTrip(req)
}

func httpGet(client *http.Client, url string) ([]byte, *ErrorResponse) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if resp == nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errRes := &ErrorResponse{}
		_ = json.Unmarshal(body, errRes)
		return nil, errRes
	}
	return body, nil
}

func httpPost(client *http.Client, url string, payload map[string]interface{}) ([]byte, *ErrorResponse) {
	payloadBytes, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	resp, err := client.Do(req)
	if resp == nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errRes := &ErrorResponse{}
		_ = json.Unmarshal(body, errRes)
		return nil, errRes
	}
	return body, nil
}
