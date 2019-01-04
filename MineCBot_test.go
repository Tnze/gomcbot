package gomcbot

import (
	"./authenticate"
	"testing"
)

func TestConnection(t *testing.T) {
	_, err := NewConnection("localhost", 25565)
	if err != nil {
		t.Errorf("connect server fail: %v", err)
	}
}

func TestPingAndList(t *testing.T) {
	resp, err := PingAndList("main.miaoscraft.ml", 25565)
	if err != nil {
		t.Errorf("ping and list server fail: %v", err)
	}
	t.Log("List:" + resp)
}

func TestLogin(t *testing.T) {
	conn, err := NewConnection("localhost", 25565)
	if err != nil {
		t.Errorf("connect server fail: %v", err)
	}
	auth := authenticate.Authenticate("Tnze", "password")
	err = conn.Login(auth)
	if err != nil {
		t.Errorf("login fail: %v", err)
	}
}
