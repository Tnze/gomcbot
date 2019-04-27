package gomcbot_test

import (
	"fmt"
	bot "github.com/Tnze/gomcbot"
	auth "github.com/Tnze/gomcbot/authenticate"
)

func ExamplePingAndList() {
	resp, err := bot.PingAndList("localhost", 25565)
	if err != nil {
		panic(err)
	}
	// see format of resp at https://wiki.vg/Server_List_Ping#Response
	fmt.Println(resp)
}

func Example_joinOfflineServer() {
	Auth := bot.Auth{
		Name: "Steve",
	}
	//Join server
	game, err := Auth.JoinServer("localhost", 25565)
	if err != nil {
		panic(err)
	}

	//Handle game
	events := game.GetEvents()
	go game.HandleGame()

	for e := range events { //Reciving events
		switch e.(type) {
		case bot.PlayerSpawnEvent:
			fmt.Println("Player is spawned!")
		}
	}
}

func Example_joinOnlineServer() {
	//Login

	// This is the basic authenticate function.
	// Maybe you could get more control of login process by using
	// https://github.com/JoshuaDoes/go-yggdrasil.
	resp, err := auth.Authenticate("email", "password")
	if err != nil {
		panic(err)
	}
	Auth := resp.ToAuth()

	//Join server
	game, err := Auth.JoinServer("localhost", 25565)
	if err != nil {
		panic(err)
	}

	//Handle game
	events := game.GetEvents()
	go game.HandleGame()

	for e := range events { //Reciving events
		switch e.(type) {
		case bot.PlayerSpawnEvent:
			fmt.Println("Player is spawned!")
		}
	}
}
