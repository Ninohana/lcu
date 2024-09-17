package lol

import (
	"testing"
)

var lcu = NewLcuClient("61563", BasicAuth{"riot", "MtTLGA_EQXPGKOHEu4c_tQ"})

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
