/*
 * Copyright © 2024 Ninohana.
 */

package lol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const remain = "Go!!!!!!!!!"

func httpGet(client http.Client, url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if resp == nil || err != nil {
		fmt.Println("请求失败")
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	return body
}

func httpPost(client http.Client, url string, payload map[string]interface{}) []byte {
	payloadBytes, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	resp, _ := client.Do(req)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	return body
}
