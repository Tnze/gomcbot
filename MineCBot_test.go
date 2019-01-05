package gomcbot

import (
	// "./authenticate"
	"testing"
)

func TestPingAndList(t *testing.T) {
	resp, err := PingAndList("localhost", 25565)
	if err != nil {
		t.Errorf("ping and list server fail: %v", err)
	}
	t.Log("Status:" + resp)
}

func TestJoinServer(t *testing.T) {
	p := Player{
		Name: "",
		UUID: "",
		AsTk: "",
	}
	_, err := p.JoinServer("localhost", 25565)
	if err != nil {
		t.Error(err)
	}
	for {
	}
}
