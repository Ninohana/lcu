github.com/Ninohana/lol -  封装完好的LCU API库，及SGP支持

# 功能

 - 订阅客户端事件
 - 获取正在进行的对局信息
 - 通过召唤师名称获取召唤师信息（LCU）
 - 通过召唤师名称获取召唤师信息（SGP）
 - 检查召唤师名称是否可用（有重复）
 - 通过puuid获取jwt

# 快速开始

直接在Go项目中import即可
```go
import "github.com/Ninohana/lol"
```

使用示例，更详细的使用方法可以查看根目录下的测试类`lcu_sgp_test.go`
```go
// 创建LCU客户端
lcu := NewLcuClient("62529", BasicAuth{"riot", "JDJE18RKuT3fldK5yc2xuA"})

// 获取召唤师信息
summoner, _ := lcu.GetSummonerByName("我玉玉了#55165")
fmt.Println(summoner)

// 开启长连接
lcu.StartWebsocket(nil, nil)
// 监听事件
lcu.Subscribe("OnJsonApiEvent", func(data interface{}) {
		fmt.Println(data) // 直接输出
})

// 创建SGP客户端
sgpToken, _ := lcu.GetSgpToken() // 获取token
sgp := NewSgpClient(sgpToken.AccessToken, Region{
		Code:     "cq100",
		Endpoint: "https://cq100-sgp.lol.qq.com:21019",
		Name:     "班德尔城",
})

// 获取正在发生的对局信息
gamingInfo, _ := sgp.GetGamingInfoByPuuid(summoner.Puuid)
fmt.Println(gamingInfo)
```
