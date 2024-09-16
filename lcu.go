/*
 * Copyright © 2024 Ninohana.
 */

package lol

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"strings"
)

// Lcu 封装了 League Client API
type Lcu struct {
	Client    *http.Client
	port      string
	auth      BasicAuth
	websocket lcuWebsocket
}

// BasicAuth 鉴权信息
type BasicAuth struct {
	UserName string
	Password string
}

func (ba BasicAuth) toString() string {
	b := fmt.Sprintf("%s:%s", ba.UserName, ba.Password)
	b = base64.StdEncoding.EncodeToString([]byte(b))
	return "Basic " + b
}

type basicAuthTransport struct {
	transport http.RoundTripper
	port      string
	basicAuth BasicAuth
}

func (transport *basicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(transport.basicAuth.UserName, transport.basicAuth.Password)
	req.URL.Scheme = "https"
	req.URL.Host = "127.0.0.1:" + transport.port
	return transport.transport.RoundTrip(req)
}

// NewLcuClient 创建一个Lcu客户端。
func NewLcuClient(port string, auth BasicAuth) *Lcu {
	lcu := new(Lcu)
	lcu.port = port
	lcu.auth = auth
	lcu.Client = &http.Client{
		Transport: &basicAuthTransport{
			transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // 跳过证书验证
				},
			},
			basicAuth: auth,
			port:      port,
		},
	}
	return lcu
}

// SgpToken 鉴权信息
type SgpToken struct {
	AccessToken string `json:"accessToken"`
	Issuer      string `json:"issuer"`
	Subject     string `json:"subject"`
	Token       string `json:"token"`
}

func (lcu *Lcu) GetSgpToken() (token *SgpToken, err error) {
	res, errRes := httpGet(*lcu.Client, "/entitlements/v1/token")
	if errRes != nil {
		return nil, &ResponseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &token)
	return token, nil
}

// GetSummonerByName 通过召唤师名称获取召唤师信息。
func (lcu *Lcu) GetSummonerByName(name string) (summoner *Summoner, err error) {
	path := fmt.Sprintf("/lol-summoner/v1/summoners?name=%s", url.QueryEscape(name))
	res, errRes := httpGet(*lcu.Client, path)
	if errRes != nil {
		return nil, &ResponseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &summoner)
	return summoner, nil
}

type lcuWebsocket struct {
	conn             *websocket.Conn
	onError          func(error)
	onUnmarshalError func(message string) bool
	dispatcher       map[string]func(interface{})
}

func (ws lcuWebsocket) subscribe(event string, handler func(interface{})) error {
	err := ws.conn.WriteJSON([]interface{}{5, event})
	if err != nil {
		return err
	}
	ws.dispatcher[event] = handler
	return nil
}

func (ws lcuWebsocket) unsubscribe(event string) error {
	err := ws.conn.WriteJSON([]interface{}{6, event})
	if err != nil {
		return err
	}
	delete(ws.dispatcher, event)
	return nil
}

// Subscribe 订阅一个客户端事件。
//
// event: 事件
//
// handler: 回调函数
func (lcu *Lcu) Subscribe(event string, handler func(interface{})) error {
	return lcu.websocket.subscribe(event, handler)
}

// Unsubscribe 取消订阅一个客户端事件。
func (lcu *Lcu) Unsubscribe(event string) error {
	return lcu.websocket.unsubscribe(event)
}

// StartWebsocket 启动一个websocket连接。
//
// 可用于获取客户端事件
//
// onError: 错误回调，可为nil
//
// onUnmarshalError: 解析错误回调，返回true继续解析，false结束解析，为nil则默认为true
//
// 返回错误
func (lcu *Lcu) StartWebsocket(onError func(error), onUnmarshalError func(message string) bool) error {
	dialer := websocket.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 跳过证书验证
		},
	}
	conn, _, err := dialer.Dial(
		"wss://127.0.0.1:"+lcu.port,
		http.Header{"Authorization": []string{lcu.auth.toString()}},
	)
	if err != nil {
		fmt.Println("连接失败")
		return err
	}

	if onError == nil {
		onError = func(err error) {
			fmt.Printf("Websocket错误: %v\n", err)
		}
	}
	if onUnmarshalError == nil {
		onUnmarshalError = func(message string) bool {
			return true
		}
	}
	lcu.websocket = lcuWebsocket{
		conn:             conn,
		onError:          onError,
		onUnmarshalError: onUnmarshalError,
		dispatcher:       make(map[string]func(interface{})),
	}

	go lcu.websocket.listen()

	return nil
}

func (ws lcuWebsocket) listen() {
	for {
		_, message, err := ws.conn.ReadMessage()
		if err != nil {
			ws.onError(err)
		}

		proto := new([]interface{})
		err = json.Unmarshal(message, &proto)
		if err != nil {
			if ws.onUnmarshalError(string(message)) {
				continue
			} else {
				break
			}
		}
		ws.dispatcher[(*proto)[1].(string)]((*proto)[2])
	}
	ws.conn.Close()
}

// ResponseError 接口返回的错误信息。
type ResponseError struct {
	Message string
}

func (error ResponseError) Error() string {
	return error.Message
}

// Spectate 观战。
//
// summonerName: 召唤师名称
// puuid: puuid
//
// 返回接口返回
func (lcu *Lcu) Spectate(name string, tagline string, puuid string) (isSuccess bool, err error) {
	url := "/lol-spectator/v1/spectate/launch"
	payload := map[string]interface{}{
		"allowObserveMode":     "ALL",
		"dropInSpectateGameId": fmt.Sprintf("%s#%s", name, tagline),
		"gameQueueType":        "",
		"puuid":                puuid,
	}
	res, errRes := httpPost(*lcu.Client, url, payload)
	if errRes != nil {
		return false, &ResponseError{Message: errRes.Message}
	}
	return len(res) == 0, nil
}

func (lcu *Lcu) GetServiceEndpoint() (string, error) {
	url := "/lol-platform-config/v1/namespaces/PlayerPreferences/ServiceEndpoint"
	res, errRes := httpGet(*lcu.Client, url)
	if errRes != nil {
		return "", &ResponseError{Message: errRes.Message}
	}
	return strings.ReplaceAll(string(res), `"`, ""), nil
}

func (lcu *Lcu) GetPlatformId() (string, error) {
	url := "/lol-platform-config/v1/namespaces/LoginDataPacket/platformId"
	res, errRes := httpGet(*lcu.Client, url)
	if errRes != nil {
		return "", &ResponseError{Message: errRes.Message}
	}
	return strings.ReplaceAll(string(res), `"`, ""), nil
}
