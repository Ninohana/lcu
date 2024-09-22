/*
 * Copyright © 2024 Ninohana.
 */

package lcu

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// sgpClient 封装了SGP API。
type sgpClient struct {
	*http.Client
	Auth   OAuth
	Region Region
}

// OAuth 鉴权信息
type OAuth struct {
	AccessToken string `json:"accessToken"`
}

func (oa OAuth) toString() string {
	return "Bearer " + oa.AccessToken
}

func (oa OAuth) setAuth(req *http.Request) {
	req.Header.Set("Authorization", oa.toString())
	req.Header.Set("Content-Type", "application/json")
}

// Region 大区信息
type Region struct {
	Code     string `json:"code"`           // 大区代码
	Endpoint string `json:"endpoint"`       // 接口地址
	Name     string `json:"name,omitempty"` // 大区名
}

var HN10 = Region{
	Code:     "HN10",
	Endpoint: "https://hn10-k8s-sgp.lol.qq.com:21019",
	Name:     "黑色玫瑰",
}

var CQ100 = Region{
	Code:     "CQ100",
	Endpoint: "https://cq100-sgp.lol.qq.com:21019",
	Name:     "班德尔城",
}

//...

// NewSgpClient 创建一个Sgp客户端。
func NewSgpClient(accessToken string, region Region) *sgpClient {
	sgp := new(sgpClient)
	sgp.Region = region
	sgp.Auth = OAuth{accessToken}
	sgp.Client = &http.Client{
		Transport: authTransport{
			&http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // 跳过证书验证
				},
			},
			sgp.Auth,
		},
	}
	return sgp
}

// RefreshToken 刷新token。
//
// 疑似过时
func (sgp *sgpClient) RefreshToken() {
	url := fmt.Sprintf("%s/session-external/v1/session/refresh", sgp.Region.Endpoint)
	lst := map[string]interface{}{
		"lst": sgp.Auth.AccessToken,
	}
	res, errRes := httpPost(sgp.Client, url, lst)
	if errRes != nil {
		panic(errRes)
	}
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
func (sgp *sgpClient) GetSummonerByName(name string) (summoner *SummonerViaSgp, err error) {
	//url := fmt.Sprintf("%s/summoner-ledge/v1/regions/%s/summoners/names", region.Endpoint, region.Code)
	url := fmt.Sprintf("%s/summoner-ledge/v1/regions/%s/summoners/name/%s", sgp.Region.Endpoint, sgp.Region.Code, name)
	res, errRes := httpGet(sgp.Client, url)
	if errRes != nil {
		return nil, &responseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &summoner)
	return summoner, nil
}

// CheckName 检查召唤师名称是否可用。
func (sgp *sgpClient) CheckName(name string) (bool, error) {
	url := fmt.Sprintf("%s/summoner-ledge/v1/regions/%s/checkname?summonerName=%s", sgp.Region.Endpoint, sgp.Region.Code, name)
	res, errRes := httpGet(sgp.Client, url)
	if errRes != nil {
		return false, &responseError{Message: errRes.Message}
	}
	isValid, _ := strconv.ParseBool(string(res))
	return isValid, nil
}

// GetJwtByPuuid 通过puuid获取jwt。
//
// 用于访问客户端聊天功能
func (sgp *sgpClient) GetJwtByPuuid(puuid string) (string, error) {
	url := fmt.Sprintf("%s/summoner-ledge/v1/regions/%s/summoners/puuid/%s/jwt", sgp.Region.Endpoint, sgp.Region.Code, puuid)
	res, errRes := httpGet(sgp.Client, url)
	if errRes != nil {
		return "", &responseError{Message: errRes.Message}
	}
	return string(res), nil
}

// GetGamingInfoByPuuid 获取正在进行的对局信息。
func (sgp *sgpClient) GetGamingInfoByPuuid(puuid string) (gamingInfo *GamingInfo, err error) {
	url := fmt.Sprintf("%s/gsm/v1/ledge/spectator/region/%s/puuid/%s", sgp.Region.Endpoint, sgp.Region.Code, puuid)
	res, errRes := httpGet(sgp.Client, url)
	if errRes != nil {
		return nil, &responseError{Message: errRes.Message}
	}
	_ = json.Unmarshal(res, &gamingInfo)
	return gamingInfo, nil
}
