package util

import (
	"fmt"
	bot "github.com/Tnze/gomcbot"
	"testing"
	"time"
)

func TestTweenLook(t *testing.T) {
	p := bot.Auth{
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
		case bot.PlayerSpawnEvent:
			fmt.Println("Player Spawn")
			go func() {
				for {
					TweenLook(g, -60, 60, time.Second*2)

					TweenLook(g, 60, -60, time.Second*2)
				}
			}()

		default:
			// fmt.Println(e)
		}
	}
}

func TestLineMove(t *testing.T) {
	p := bot.Auth{
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
		case bot.PlayerSpawnEvent:
			fmt.Println("Player Spawn")
			go func() {
				for {
					err := TweenLineMove(g, 9, -130, time.Second*1)
					if err != nil {
						fmt.Println(err)
					}
					err = TweenLineMove(g, 9, -130-5, time.Second*1)
					if err != nil {
						fmt.Println(err)
					}
					err = TweenLineMove(g, 9+5, -130-5, time.Second*1)
					if err != nil {
						fmt.Println(err)
					}
					err = TweenLineMove(g, 9+5, -130, time.Second*1)
					if err != nil {
						fmt.Println(err)
					}
				}
			}()

		default:
			// fmt.Println(e)
		}
	}
}

func TestTweenJump(t *testing.T) {
	p := bot.Auth{
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
		case bot.PlayerSpawnEvent:
			fmt.Println("Player Spawn")
			go func() {
				for {
					TweenJump(g)
					time.Sleep(time.Second)
				}
			}()

		default:
			// fmt.Println(e)
		}
	}
}
