package minecbot

import (
	"testing"
)

func TestConnection(t *testing.T) {
	_, err := NewConnection("localhost", 25565)
	if err != nil {
		t.Errorf("connect server fail: %v", err)
	}
}

func TestPingAndList(t *testing.T) {
	resp, err := PingAndList("localhost", 25565)
	if err != nil {
		t.Errorf("ping and list server fail: %v", err)
	}
	t.Log("List:" + resp)
}
