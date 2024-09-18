package lol

import (
	"testing"
)

var lcu = NewLcuClient("53559", BasicAuth{"riot", "npvaqXxkZnx0KMJu8jPSaA"})

func TestLcu_getServiceEndpoint(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"base", "https://cq100-sgp.lol.qq.com:21019", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lcu.GetServiceEndpoint()
			if got != tt.want {
				t.Errorf("getServiceEndpoint() got = %v, want %v", got, tt.want)
			}
		})
	}
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

func TestLcu_getPlatformId(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"base", "CQ100", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lcu.GetPlatformId()
			if got != tt.want {
				t.Errorf("getPlatformId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLcu_GetReplaysConfiguration(t *testing.T) {
	tests := []struct {
		name              string
		wantConfiguration bool
		wantErr           bool
	}{
		{"base", true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConfiguration, err := lcu.GetReplaysConfiguration()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReplaysConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (gotConfiguration == nil) == tt.wantConfiguration {
				t.Errorf("GetReplaysConfiguration() gotConfiguration = %v, want %v", gotConfiguration, tt.wantConfiguration)
				return
			}
		})
	}
}

func TestLcu_GetRoflsPath(t *testing.T) {
	tests := []struct {
		name       string
		wantNotNil bool
	}{
		{"base", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcu.GetRoflsPath(); (len(got) == 0) == tt.wantNotNil {
				t.Errorf("GetRoflsPath() = %v, want %v", got, tt.wantNotNil)
			}
		})
	}
}

func TestLcu_GetRoflsDefaultPath(t *testing.T) {
	tests := []struct {
		name       string
		wantNotNil bool
	}{
		{"base", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcu.GetRoflsDefaultPath(); (len(got) == 0) == tt.wantNotNil {
				t.Errorf("GetRoflsPath() = %v, want %v", got, tt.wantNotNil)
			}
		})
	}
}

func TestLcu_GetSummonerGamesByPuuid(t *testing.T) {
	type args struct {
		puuid string
		begin int
		end   int
	}
	tests := []struct {
		name      string
		args      args
		wantGames bool
		wantErr   bool
	}{
		{"base", args{"c9ea4cd2-fd41-5656-b615-49056d444271", 0, 4}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGames, err := lcu.GetSummonerGamesByPuuid(tt.args.puuid, tt.args.begin, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSummonerGamesByPuuid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//prettyPrint(gotGames)
			if gotGames == nil == tt.wantGames {
				t.Errorf("GetSummonerGamesByPuuid() gotGames = %v, want %v", gotGames, tt.wantGames)
			}
		})
	}
}

func TestLcu_GetGameInfoByGameId(t *testing.T) {
	type args struct {
		gameId int64
	}
	tests := []struct {
		name      string
		args      args
		wantGames bool
		wantErr   bool
	}{
		{"base", args{500203450300}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGames, err := lcu.GetGameInfoByGameId(tt.args.gameId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSummonerGamesByPuuid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//prettyPrint(gotGames)
			if gotGames == nil == tt.wantGames {
				t.Errorf("GetSummonerGamesByPuuid() gotGames = %v, want %v", gotGames, tt.wantGames)
			}
		})
	}
}
