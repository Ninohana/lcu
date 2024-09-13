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

func prettyPrintJson(data map[string]interface{}) {
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(jsonData))
}

func prettyPrintJsonWithTag(tag string, data map[string]interface{}) {
	fmt.Printf("=========%s=========\n", tag)
	prettyPrintJson(data)
	fmt.Printf("=========%s=========\n", strings.Repeat("==", len(tag)/3))
}
func TestNewLcuClient(t *testing.T) {
	lcu := NewLcuClient("57851", BasicAuth{"riot", "R1aBr6N1bmBaTS_D0g5sjw"})

	prettyPrintJsonWithTag("LCU获取召唤师信息", lcu.GetSummonerByName("我玉玉了#55165")) // 班德尔城，大佬带带我

	lcu.StartWebsocket(nil, nil)
	err := lcu.Subscribe("OnJsonApiEvent", func(data interface{}) {
		fmt.Println(data)
	})
	if err != nil {
		t.Errorf("subscribe failed")
	}

	//err = lcu.Unsubscribe("OnJsonApiEvent")
	//if err != nil {
	//	panic(err)
	//}
}

func TestNewSgpClient(t *testing.T) {
	lcu := NewLcuClient("57851", BasicAuth{"riot", "R1aBr6N1bmBaTS_D0g5sjw"})

	sgpToken := lcu.GetSgpToken()

	sgp := NewSgpClient(sgpToken.AccessToken, Region{
		Code:     "cq100",
		Endpoint: "https://cq100-sgp.lol.qq.com:21019",
		Name:     "班德尔城",
	})

	prettyPrintJsonWithTag("SGP获取召唤师信息", sgp.GetSummonerByName("我玉玉了"))
	prettyPrintJsonWithTag("获取正在发生的对局信息", sgp.GetGamingInfoByPuuid("c9ea4cd2-fd41-5656-b615-49056d444271"))
	fmt.Println(sgp.GetJwtByPuuid("c9ea4cd2-fd41-5656-b615-49056d444271"))
	fmt.Println(sgp.CheckName("我玉玉了"))
	sgp.RefreshToken()
}
