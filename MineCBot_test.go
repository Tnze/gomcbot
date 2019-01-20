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

	for e := range events {
		switch e {
		case PlayerSpawnEvent:
			fmt.Println(g.player.X, g.player.Y, g.player.Z)
			go func() {
				for {
					time.Sleep(time.Millisecond * 500)
					player := g.GetPlayer()
					fmt.Println("HeldItem: ", player.HeldItem)
					fmt.Println("Inventory: ", player.Inventory)
				}
			}()
		case PlayerDeadEvent:
			fmt.Println("Player Dead")
		case InventoryChangeEvent:
			fmt.Println("Inventory Change")
		default:
			fmt.Println(e)
		}
	}
}
