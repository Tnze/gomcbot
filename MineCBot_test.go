package gomcbot

import (
	"fmt"
	"testing"
	"time"
)

func TestPingAndList(t *testing.T) {
	resp, err := PingAndList("localhost", 25565)
	if err != nil {
		t.Errorf("ping and list server fail: %v", err)
	}
	t.Log("Status:" + resp)
}

func TestJoinServerOffline(t *testing.T) {
	p := Auth{
		Name: "Mi_Xi_Xi",
		UUID: "ff7a038f265c4d42b0cf04c575896469",
		AsTk: "",
	}
	g, err := p.JoinServer("localhost", 25565)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Login success")
	events := g.GetEvents()
	go g.HandleGame()
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 50)
			g.LookYawPitch(g.player.Yaw+10, g.player.Pitch)
			g.SetPosition(g.player.X+0.01, g.player.Y, g.player.Z)
		}
	}()
	for e := range events {
		switch e {
		case PlayerSpawnEvent:
			fmt.Println(g.player.X, g.player.Y, g.player.Z)
		case PlayerDeadEvent:
			fmt.Println("Player Dead")
		default:
			fmt.Println(e)
		}
	}
}
