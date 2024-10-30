/*
 * Copyright © 2024 Ninohana.
 */

package lcu

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func prettyPrint(v any) {
	prettyJson, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(prettyJson))
}

func prettyPrintWithTag(tag string, v any) {
	fmt.Printf("=========%s=========\n", tag)
	prettyPrint(v)
	fmt.Printf("=========%s=========\n", strings.Repeat("==", len(tag)/3))
}

func TestNewLcuClient(t *testing.T) {
	summoner, err := lcu.GetSummonerByName("我玉玉了#55165") // 班德尔城，大佬带带我
	if err != nil {
		t.Error(err)
	}
	prettyPrint(summoner)

	_ = lcu.StartWebsocket(nil, nil)
	//_ = lcuClient.StartWebsocket(func(err error) {
	//	panic(err)
	//}, func(message string) bool {
	//	fmt.Println(message)
	//	return true
	//})
	err = lcu.Subscribe("OnJsonApiEvent", func(data interface{}) {
		fmt.Println(data)
	})
	if err != nil {
		t.Errorf("订阅失败")
	}

	err = lcu.Unsubscribe("OnJsonApiEvent")
	if err != nil {
		t.Errorf("取消订阅失败")
	}
}

func TestNewSgpClient(t *testing.T) {
	sgpToken, _ := lcu.GetSgpToken()
	sgpClient := NewSgpClient(sgpToken.AccessToken, CQ100)

	summoner, err := sgpClient.GetSummonerByName("我玉玉了")
	if err != nil {
		t.Error(err)
	} else {
		prettyPrint(summoner)

		gamingInfo, err := sgpClient.GetGamingInfoByPuuid(summoner.Puuid)
		if err != nil {
			t.Error(err)
		} else {
			prettyPrintWithTag("获取正在发生的对局信息", gamingInfo)
		}

		jwt, err := sgpClient.GetJwtByPuuid(summoner.Puuid)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println(jwt)
		}
	}

	isValid, err := sgpClient.CheckName("我玉玉了")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("名字不重复？", isValid)
	}
	//sgpClient.RefreshToken()
}

// 测试自定义请求接口
//
// 使用lcu客户端发送请求，此时如果已经初始化过，则不需要传入鉴权信息，lcu客户端内部会自动管理，在需要时添加
//
// 可以在https://www.mingweisamuel.com/lcu-schema/tool/#/查到所有接口信息
func TestLcuClient_CustomRequest(t *testing.T) {
	// 创建一个http请求
	req, _ := http.NewRequest("GET", "/entitlements/v1/token", nil)
	// 发送请求
	resp, err := lcu.Do(req)
	if resp == nil {
		panic(err)
	}
	// 释放资源
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	// 获取请求返回值
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errRes := &errorResponse{}
		_ = json.Unmarshal(body, errRes)
		t.Error(errRes)
		return
	}
	fmt.Println(string(body))
}
