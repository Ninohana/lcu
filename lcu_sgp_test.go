/*
 * Copyright © 2024 Ninohana.
 */

package lol

import (
	"encoding/json"
	"fmt"
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
	lcu := NewLcuClient("55378", BasicAuth{"riot", "WleZotZkvvdIHcUfT9Pa4Q"})

	summoner, err := lcu.GetSummonerByName("我玉玉了#55165") // 班德尔城，大佬带带我
	if err != nil {
		t.Error(err)
	}
	prettyPrint(summoner)

	_ = lcu.StartWebsocket(nil, nil)
	//_ = lcu.StartWebsocket(func(err error) {
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
	lcu := NewLcuClient("55378", BasicAuth{"riot", "WleZotZkvvdIHcUfT9Pa4Q"})

	sgpToken, _ := lcu.GetSgpToken()
	sgp := NewSgpClient(sgpToken.AccessToken, HN10)

	summoner, err := sgp.GetSummonerByName("我玉玉了")
	if err != nil {
		t.Error(err)
	} else {
		prettyPrint(summoner)

		gamingInfo, err := sgp.GetGamingInfoByPuuid(summoner.Puuid)
		if err != nil {
			t.Error(err)
		} else {
			prettyPrintWithTag("获取正在发生的对局信息", gamingInfo)
		}

		jwt, err := sgp.GetJwtByPuuid(summoner.Puuid)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println(jwt)
		}
	}

	isValid, err := sgp.CheckName("我玉玉了")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("名字不重复？", isValid)
	}
	sgp.RefreshToken()
}
