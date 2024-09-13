/*
 * Copyright © 2024 Ninohana.
 */

package lol

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Sgp 封装了SGP API。
type Sgp struct {
	Client *http.Client
	auth   OAuth
	region Region
}

// OAuth 鉴权信息
type OAuth struct {
	AccessToken string `json:"accessToken"`
}

func (oa *OAuth) toString() string {
	return "Bearer " + oa.AccessToken
}

type oAuthTransport struct {
	transport http.RoundTripper
	auth      OAuth
}

func (transport *oAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", transport.auth.toString())
	req.Header.Set("Content-Type", "application/json")
	return transport.transport.RoundTrip(req)
}

// Region 大区信息
type Region struct {
	Code     string `json:"code"`
	Endpoint string `json:"endpoint"`
	Name     string `json:"name"`
}

// NewSgpClient 创建一个Sgp客户端。
func NewSgpClient(accessToken string, region Region) *Sgp {
	sgp := new(Sgp)
	sgp.region = region
	sgp.auth = OAuth{
		AccessToken: accessToken,
	}
	sgp.Client = &http.Client{
		Transport: &oAuthTransport{
			auth: sgp.auth,
			transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // 跳过证书验证
				},
			},
		},
	}
	return sgp
}

// RefreshToken 刷新token。
//
// 疑似过时
func (sgp *Sgp) RefreshToken() {
	url := fmt.Sprintf("%s/session-external/v1/session/refresh", sgp.region.Endpoint)
	lst := map[string]interface{}{
		"lst": sgp.auth.AccessToken,
	}
	res := httpPost(*sgp.Client, url, lst)
	println(string(res))
}

// GetSummonerByName 通过召唤师名称获取召唤师信息。
//
// 更新后召唤师名称可能重复，如果想查询的召唤师名称有重复，该函数不可靠。
// 考虑更换使用lcu中的函数指定召唤师的tagline精确查询。
//
// name: 召唤师名称
//
// 该函数构造请求URL，使用HTTP请求获取召唤师信息，并返回响应的主体内容。
func (sgp *Sgp) GetSummonerByName(name string) (rJson map[string]interface{}) {
	//url := fmt.Sprintf("%s/summoner-ledge/v1/regions/%s/summoners/names", region.Endpoint, region.Code)
	url := fmt.Sprintf("%s/summoner-ledge/v1/regions/%s/summoners/name/%s", sgp.region.Endpoint, sgp.region.Code, name)
	res := httpGet(*sgp.Client, url)
	_ = json.Unmarshal(res, &rJson)
	return rJson
}

// CheckName 检查召唤师名称是否可用。
func (sgp *Sgp) CheckName(name string) bool {
	url := fmt.Sprintf("%s/summoner-ledge/v1/regions/%s/checkname?summonerName=%s", sgp.region.Endpoint, sgp.region.Code, name)
	res := httpGet(*sgp.Client, url)
	isValid, _ := strconv.ParseBool(string(res))
	return isValid
}

// GetJwtByPuuid 通过puuid获取jwt。
//
// 用于访问客户端聊天功能
func (sgp *Sgp) GetJwtByPuuid(puuid string) string {
	url := fmt.Sprintf("%s/summoner-ledge/v1/regions/%s/summoners/puuid/%s/jwt", sgp.region.Endpoint, sgp.region.Code, puuid)
	res := httpGet(*sgp.Client, url)
	return string(res)
}

// GetGamingInfoByPuuid 获取正在进行的对局信息。
func (sgp *Sgp) GetGamingInfoByPuuid(puuid string) (rJson map[string]interface{}) {
	url := fmt.Sprintf("%s/gsm/v1/ledge/spectator/region/%s/puuid/%s", sgp.region.Endpoint, sgp.region.Code, puuid)
	res := httpGet(*sgp.Client, url)
	_ = json.Unmarshal(res, &rJson)
	return rJson
}
