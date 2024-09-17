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
	sgpToken, _ := lcu.GetSgpToken()
	sgp := NewSgpClient(sgpToken.AccessToken, CQ100)

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
	//sgp.RefreshToken()
}

func TestLcu_GetSummonerByPuuid(t *testing.T) {
	type args struct {
		puuid string
	}
	tests := []struct {
		name         string
		args         args
		wantSummoner bool
		wantErr      bool
	}{
		{"base", args{"c9ea4cd2-fd41-5656-b615-49056d444271"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSummoner, err := lcu.GetSummonerByPuuid(tt.args.puuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSummonerByPuuid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSummoner != nil && !tt.wantSummoner {
				t.Errorf("GetSummonerByPuuid() gotSummoner = %v, wantSummoner %v", gotSummoner, tt.wantSummoner)
			}
		})
	}
}
