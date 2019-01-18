package gomcbot

import (
	"github.com/Tnze/gomcbot/authenticate"
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
	resp, err := authenticate.Authenticate("", "")
	if err != nil {
		t.Fatal(err)
	}
	p := Auth{
		Name: resp.SelectedProfile.Name,
		UUID: resp.SelectedProfile.ID,
		AsTk: resp.AccessToken,
	}
	g, err := p.JoinServer("localhost", 25565)
	if err != nil {
		t.Fatal(err)
	}
	err = g.HandleGame()
	if err != nil {
		t.Fatal(err)
	}
}

func TestJoinServerOffline(t *testing.T) {

	p := Auth{
		Name: "",
		UUID: "",
		AsTk: "",
	}
	g, err := p.JoinServer("localhost", 25565)
	if err != nil {
		t.Fatal(err)
	}
	err = g.HandleGame()
	if err != nil {
		t.Fatal(err)
	}
}
