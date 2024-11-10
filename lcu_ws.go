package lcu

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type lcuWebsocket struct {
	conn       *websocket.Conn
	wg         *sync.WaitGroup
	onError    func(error)
	onPanic    func(any)
	dispatcher map[string]func(*lwsMessageContent)
}

type lwsMessageContent struct {
	Data      any    `json:"data"`
	EventType string `json:"eventType"`
	Uri       string `json:"uri"`
}

// Subscribe 订阅一个客户端事件。
//
// event: 事件
//
// handler: 回调函数
func (ws *lcuWebsocket) Subscribe(event string, handler func(*lwsMessageContent)) error {
	err := ws.conn.WriteJSON([]interface{}{5, event})
	if err != nil {
		return err
	}
	ws.dispatcher[event] = handler
	return nil
}

// Unsubscribe 取消订阅一个客户端事件。
func (ws *lcuWebsocket) Unsubscribe(event string) error {
	err := ws.conn.WriteJSON([]interface{}{6, event})
	if err != nil {
		return err
	}
	delete(ws.dispatcher, event)
	return nil
}

// StartWebsocket 启动一个websocket连接。
//
// 可用于获取客户端事件
//
// onError: 错误回调，可为nil
//
// onPanic: 异常回调，可为nil
func (lcu *lcuClient) StartWebsocket(onError func(error), onPanic func(any)) (*lcuWebsocket, error) {
	dialer := websocket.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 跳过证书验证
		},
	}
	conn, _, err := dialer.Dial(
		"wss://127.0.0.1:"+lcu.Port,
		http.Header{"Authorization": []string{lcu.Auth.toString()}},
	)
	if err != nil {
		fmt.Println("连接失败")
		return nil, err
	}

	if onError == nil {
		onError = func(err error) {
			fmt.Printf("Websocket错误: %v\n", err)
		}
	}
	if onPanic == nil {
		onPanic = func(recover any) {
			fmt.Printf("Websocket异常: %v\n", recover)
		}
	}
	ws := &lcuWebsocket{
		conn:       conn,
		wg:         new(sync.WaitGroup),
		onError:    onError,
		onPanic:    onPanic,
		dispatcher: make(map[string]func(*lwsMessageContent)),
	}

	go ws.listen()

	return ws, nil
}

func (ws *lcuWebsocket) listen() {
	ws.wg.Add(1)
	defer func() {
		ws.wg.Done()
		if r := recover(); r != nil {
			ws.onPanic(r)
		}
	}()

	for {
		_, message, err := ws.conn.ReadMessage()
		if err != nil {
			ws.onError(err)
			continue
		}
		if len(message) == 0 {
			continue
		}

		proto := new([]any)
		// The structure of proto message like: [code.(int), event.(string), data.(lwsMessageContent)]
		err = json.Unmarshal(message, &proto)
		if err != nil {
			ws.onError(err)
			continue
		}

		var msgContent lwsMessageContent
		t := (*proto)[2].(map[string]any)
		msgContent.Data = t["data"]
		msgContent.EventType = t["eventType"].(string)
		msgContent.Uri = t["uri"].(string)
		ws.dispatcher[(*proto)[1].(string)](&msgContent)
	}
}
