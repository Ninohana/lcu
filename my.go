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

type errorResponse struct {
	ErrorCode             string      `json:"errorCode"`
	HttpStatus            int         `json:"httpStatus"`
	ImplementationDetails interface{} `json:"implementationDetails"`
	Message               string      `json:"message"`
}

// responseError 接口返回的错误信息。
type responseError struct {
	Message string
}

func (error responseError) Error() string {
	return error.Message
}

type authSetter interface {
	setAuth(*http.Request)
}

type authTransport struct {
	http.RoundTripper
	authSetter
}

func (at authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	at.setAuth(req)
	return at.RoundTripper.RoundTrip(req)
}

func httpGet(client *http.Client, url string) ([]byte, *errorResponse) {
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
		errRes := &errorResponse{}
		_ = json.Unmarshal(body, errRes)
		return nil, errRes
	}
	return body, nil
}

func httpPost(client *http.Client, url string, payload map[string]interface{}) ([]byte, *errorResponse) {
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
		errRes := &errorResponse{}
		_ = json.Unmarshal(body, errRes)
		return nil, errRes
	}
	return body, nil
}
