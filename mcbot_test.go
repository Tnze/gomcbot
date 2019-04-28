package gomcbot_test

import (
	"fmt"
	"testing"
	// "time"

	bot "github.com/Tnze/gomcbot"
	auth "github.com/Tnze/gomcbot/util/authenticate"
)

func TestPingAndList(t *testing.T) {
	resp, err := bot.PingAndList("jdao.online", 25566)
	if err != nil {
		t.Errorf("ping and list server fail: %v", err)
	}
	t.Log("Status:" + resp)
}

func TestJoinServerOffline(t *testing.T) {
	c := bot.NewClient()

	err := c.JoinServer("jdao.online", 25566)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Login success")
}

func TestJoinOnlineServer(t *testing.T) {
	c := bot.NewClient()
	//Login

	// This is the basic authenticate function.
	// Maybe you could get more control of login process by using
	// https://github.com/JoshuaDoes/go-yggdrasil.
	resp, err := auth.Authenticate("email", "password")
	if err != nil {
		panic(err)
	}
	c.Auth = resp.ToAuth()

	//Join server
	err = c.JoinServer("localhost", 25565)
	if err != nil {
		panic(err)
	}

	//Handle game
	// events := game.GetEvents()
	// go game.HandleGame()

	// for e := range events { //Reciving events
	// 	switch e.(type) {
	// 	case bot.PlayerSpawnEvent:
	// 		fmt.Println("Player is spawned!")
	// 	}
	// }
}
