/*
 * Copyright © 2024 Ninohana.
 */

package lol

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Lcu 封装了 League Client API
type Lcu struct {
	Client        *http.Client
	Port          string
	Auth          BasicAuth
	authTransport authTransport
	websocket     lcuWebsocket
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
	lcu *Lcu
}

func (l *localTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "https"
	req.URL.Host = "127.0.0.1:" + l.lcu.Port
	return l.RoundTripper.RoundTrip(req)
}

// NewLcuClient 创建一个Lcu客户端。
func NewLcuClient(port string, auth BasicAuth) *Lcu {
	lcu := new(Lcu)
	lcu.Port = port
	lcu.Auth = auth
	lcu.authTransport = authTransport{
		&localTransport{
			&http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // 跳过证书验证
				},
			},
			lcu,
		},
		lcu.Auth,
	}
	lcu.Client = &http.Client{Transport: lcu.authTransport}
	return lcu
}

// ResponseError 接口返回的错误信息。
type ResponseError struct {
	Message string
}

func (error ResponseError) Error() string {
	return error.Message
}

// GetSgpToken 获取SGP Token。
func (lcu *Lcu) GetSgpToken() (token *SgpToken, err error) {
	res, errRes := httpGet(lcu.Client, "/entitlements/v1/token")
	if errRes != nil {
		return nil, &ResponseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &token)
	return token, nil
}

// GetSummonerByName 通过召唤师名称获取召唤师信息。
func (lcu *Lcu) GetSummonerByName(name string) (summoner *Summoner, err error) {
	path := fmt.Sprintf("/lol-summoner/v1/summoners?name=%s", url.QueryEscape(name))
	res, errRes := httpGet(lcu.Client, path)
	if errRes != nil {
		return nil, &ResponseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &summoner)
	return summoner, nil
}

// GetSummonerByPuuid 通过召唤师puuid获取召唤师信息。
func (lcu *Lcu) GetSummonerByPuuid(puuid string) (summoner *Summoner, err error) {
	path := fmt.Sprintf("/lol-summoner/v2/summoners/puuid/%s", puuid)
	res, errRes := httpGet(lcu.Client, path)
	if errRes != nil {
		return nil, &ResponseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &summoner)
	return summoner, nil
}

// GetSummonerGamesByPuuid 通过puuid获取召唤师对局信息。
//
// begin: 从第多少条开始
//
// end: 到第多少条
func (lcu *Lcu) GetSummonerGamesByPuuid(puuid string, begin int, end int) (games *GamesInfo, err error) {
	path := fmt.Sprintf(
		"/lol-match-history/v1/products/lol/%s/matches?begIndex=%d&endIndex=%d",
		puuid, begin, end)
	res, errRes := httpGet(lcu.Client, path)
	if errRes != nil {
		return nil, &ResponseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &games)
	return games, nil
}

func (lcu *Lcu) GetGameInfoByGameId(gameId int64) (game *GameInfo, err error) {
	path := fmt.Sprintf("/lol-match-history/v1/games/%d", gameId)
	res, errRes := httpGet(lcu.Client, path)
	if errRes != nil {
		return nil, &ResponseError{Message: errRes.Message}
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
func (lcu *Lcu) Spectate(name string, tagline string, puuid string) (isSuccess bool, err error) {
	url := "/lol-spectator/v1/spectate/launch"
	payload := map[string]interface{}{
		"allowObserveMode":     "ALL",
		"dropInSpectateGameId": fmt.Sprintf("%s#%s", name, tagline),
		"gameQueueType":        "",
		"puuid":                puuid,
	}
	res, errRes := httpPost(lcu.Client, url, payload)
	if errRes != nil {
		return false, &ResponseError{Message: errRes.Message}
	}
	return len(res) == 0, nil
}

// GetServiceEndpoint 获取SGP服务地址。
// 形如 "https://cq100-sgp.lol.qq.com:21019"
//
// panic: 获取失败
func (lcu *Lcu) GetServiceEndpoint() string {
	url := "/lol-platform-config/v1/namespaces/PlayerPreferences/ServiceEndpoint"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		panic(&ResponseError{Message: errRes.Message})
	}
	return strings.ReplaceAll(string(res), `"`, "")
}

// GetPlatformId 获取平台ID。
// 形如 "HN10" "CQ100"
//
// panic: 获取失败
func (lcu *Lcu) GetPlatformId() string {
	url := "/lol-platform-config/v1/namespaces/LoginDataPacket/platformId"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		panic(&ResponseError{Message: errRes.Message})
	}
	return strings.ReplaceAll(string(res), `"`, "")
}

func (lcu *Lcu) GetReplaysConfiguration() (configuration *ReplaysConfigurationV1, err error) {
	url := "/lol-replays/v1/configuration"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		return nil, &ResponseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &configuration)
	return configuration, nil
}

// GetRoflsPath 获取回放文件保存路径。
func (lcu *Lcu) GetRoflsPath() string {
	url := "/lol-replays/v1/rofls/path"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		panic(&ResponseError{Message: errRes.Message})
	}
	return strings.ReplaceAll(string(res), `"`, "")
}

// GetRoflsDefaultPath 获取回放文件保存路径。
func (lcu *Lcu) GetRoflsDefaultPath() string {
	url := "/lol-replays/v1/rofls/path/default"
	res, errRes := httpGet(lcu.Client, url)
	if errRes != nil {
		panic(&ResponseError{Message: errRes.Message})
	}
	return strings.ReplaceAll(string(res), `"`, "")
}
