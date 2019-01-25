package gomcbot

import (
	// "./util"
	"fmt"
	"testing"
	// "time"
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
		Name: "Name",
		UUID: "UUID",
		AsTk: "",
	}
	g, err := p.JoinServer("localhost", 25565)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Login success")
	events := g.GetEvents()
	go g.HandleGame()

	for e := range events {
		switch e {
		case PlayerSpawnEvent:
			fmt.Println("Player Spawn")
		case PlayerDeadEvent:
			fmt.Println("Player Dead")
		case InventoryChangeEvent:
			fmt.Println("Inventory Change")
		case BlockChangeEvent:

		default:
			// fmt.Println(e)
		}
	}
}
