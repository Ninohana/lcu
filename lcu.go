/*
 * Copyright © 2024 Ninohana.
 */

package lcu

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// lcuClient 封装了 League client API
type lcuClient struct {
	*http.Client
	Port      string
	Auth      BasicAuth
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

func (ba BasicAuth) setAuth(req *http.Request) {
	req.SetBasicAuth(ba.UserName, ba.Password)
}

type localTransport struct {
	http.RoundTripper
	lcu *lcuClient
}

func (l *localTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "https"
	req.URL.Host = "127.0.0.1:" + l.lcu.Port
	return l.RoundTripper.RoundTrip(req)
}

// NewLcuClient 创建一个Lcu客户端。
func NewLcuClient(port string, auth BasicAuth) *lcuClient {
	lcu := new(lcuClient)
	lcu.Port = port
	lcu.Auth = auth
	lcu.Client = &http.Client{
		Transport: authTransport{
			&localTransport{
				&http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true, // 跳过证书验证
					},
				},
				lcu,
			},
			lcu.Auth,
		},
	}
	return lcu
}

// GetSgpToken 获取SGP Token。
func (lcu *lcuClient) GetSgpToken() (token *SgpToken, err error) {
	res, errRes := httpGet(lcu.Client, "/entitlements/v1/token")
	if errRes != nil {
		return nil, &responseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &token)
	return token, nil
}

// GetSummonerByName 通过召唤师名称获取召唤师信息。
func (lcu *lcuClient) GetSummonerByName(name string) (summoner *Summoner, err error) {
	path := fmt.Sprintf("/lol-summoner/v1/summoners?name=%s", url.QueryEscape(name))
	res, errRes := httpGet(lcu.Client, path)
	if errRes != nil {
		return nil, &responseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &summoner)
	return summoner, nil
}

// GetSummonerByPuuid 通过召唤师puuid获取召唤师信息。
func (lcu *lcuClient) GetSummonerByPuuid(puuid string) (summoner *Summoner, err error) {
	path := fmt.Sprintf("/lol-summoner/v2/summoners/puuid/%s", puuid)
	res, errRes := httpGet(lcu.Client, path)
	if errRes != nil {
		return nil, &responseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &summoner)
	return summoner, nil
}

// GetSummonerGamesByPuuid 通过puuid获取召唤师对局信息。
//
// begin: 从第多少条开始
//
// end: 到第多少条
func (lcu *lcuClient) GetSummonerGamesByPuuid(puuid string, begin int, end int) (games *GamesInfo, err error) {
	path := fmt.Sprintf(
		"/lol-match-history/v1/products/lol/%s/matches?begIndex=%d&endIndex=%d",
		puuid, begin, end)
	res, errRes := httpGet(lcu.Client, path)
	if errRes != nil {
		return nil, &responseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &games)
	return games, nil
}

func (lcu *lcuClient) GetGameInfoByGameId(gameId int64) (game *GameInfo, err error) {
	path := fmt.Sprintf("/lol-match-history/v1/games/%d", gameId)
	res, errRes := httpGet(lcu.Client, path)
	if errRes != nil {
		return nil, &responseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &game)
	return game, nil
}

// Spectate 观战。
//
// summonerName: 召唤师名称
//
// puuid: puuid
//
// 返回接口返回
func (lcu *lcuClient) Spectate(name string, tagline string, puuid string) (isSuccess bool, err error) {
	url := "/lol-spectator/v1/spectate/launch"
	payload := map[string]interface{}{
		"allowObserveMode":     "ALL",
		"dropInSpectateGameId": fmt.Sprintf("%s#%s", name, tagline),
		"gameQueueType":        "",
		"puuid":                puuid,
	}
	res, errRes := httpPost(lcu.Client, url, payload)
	if errRes != nil {
		return false, &responseError{Message: errRes.Message}
	}
	return len(res) == 0, nil
}

// GetServiceEndpoint 获取SGP服务地址。
// 形如 "https://cq100-sgp.lol.qq.com:21019"
//
// panic: 获取失败
func (lcu *lcuClient) GetServiceEndpoint() string {
	url := "/lol-platform-config/v1/namespaces/PlayerPreferences/ServiceEndpoint"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		panic(&responseError{Message: errRes.Message})
	}
	return strings.ReplaceAll(string(res), `"`, "")
}

// GetPlatformId 获取平台ID。
// 形如 "HN10" "CQ100"
//
// panic: 获取失败
func (lcu *lcuClient) GetPlatformId() string {
	url := "/lol-platform-config/v1/namespaces/LoginDataPacket/platformId"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		panic(&responseError{Message: errRes.Message})
	}
	return strings.ReplaceAll(string(res), `"`, "")
}

func (lcu *lcuClient) GetReplaysConfiguration() (configuration *ReplaysConfigurationV1, err error) {
	url := "/lol-replays/v1/configuration"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		return nil, &responseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &configuration)
	return configuration, nil
}

// GetRoflsPath 获取回放文件保存路径。
func (lcu *lcuClient) GetRoflsPath() string {
	url := "/lol-replays/v1/rofls/path"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		panic(&responseError{Message: errRes.Message})
	}
	return strings.ReplaceAll(string(res), `"`, "")
}

// GetRoflsDefaultPath 获取回放文件保存路径。
func (lcu *lcuClient) GetRoflsDefaultPath() string {
	url := "/lol-replays/v1/rofls/path/default"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		panic(&responseError{Message: errRes.Message})
	}
	return strings.ReplaceAll(string(res), `"`, "")
}

func (lcu *lcuClient) GetCurrentSummonerProfile() (summonerProfile *SummonerProfile, err error) {
	url := "/lol-summoner/v1/current-summoner/summoner-profile"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		return nil, &responseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &summonerProfile)
	return summonerProfile, nil
}

func (lcu *lcuClient) GetGameFlowPhase() string {
	url := "/lol-gameflow/v1/gameflow-phase"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		fmt.Println(errRes.Message)
		return ""
	}
	return strings.ReplaceAll(string(res), `"`, "")
}

func (lcu *lcuClient) GetGameflowSession() (gameflowInfo *GameflowInfo, err error) {
	url := "/lol-gameflow/v1/session"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		return nil, &responseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &gameflowInfo)
	return gameflowInfo, nil
}

// AcceptTrade 接受英雄交换请求。
//
// actionId: 交换请求的ActionId
func (lcu *lcuClient) AcceptTrade(actionId int) error {
	url := fmt.Sprintf("/lol-champ-select/v1/session/trades/%d/accept", actionId)
	_, errRes := httpPost(lcu.Client, url, nil)
	if errRes != nil {
		return &responseError{Message: errRes.Message}
	}
	return nil
}

func (lcu *lcuClient) GetSelectSession() (selectSession *SelectSession, err error) {
	url := "/lol-champ-select/v1/session"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		return nil, &responseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &selectSession)
	return selectSession, nil
}
