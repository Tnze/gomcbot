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
			fmt.Println(int(g.player.X), int(g.player.Y), int(g.player.Z))
			// go func() {
			// 	g.Dig(-32, 66, -239)
			// }()
			go func() {
				for {
					time.Sleep(time.Second * 3)
					for i := -5; i < 5; i++ {
						for j := -5; j < 5; j++ {
							fmt.Printf("%3d", g.GetBlock(int(g.player.X)+i, int(g.player.Y)-1, int(g.player.Z)+j).id)
						}
						fmt.Println()
					}
					fmt.Printf("from %d to %d, %d to %d\n", int(g.player.X)-5, int(g.player.X)+4, int(g.player.Z)-5, int(g.player.Z)+4)
					fmt.Println()
				}
			}()

		case PlayerDeadEvent:
			fmt.Println("Player Dead")
		case InventoryChangeEvent:
			fmt.Println("Inventory Change")
		default:
			// fmt.Println(e)
		}
	}
}
