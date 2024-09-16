package lol

import (
	"testing"
)

var lcu = NewLcuClient("63539", BasicAuth{"riot", "upIxPFmtVSfQWr7i4NeN2g"})

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
