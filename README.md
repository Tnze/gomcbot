# GoMCbot
A golang minecraft robot package

- Current Version: `1.13.2`  
- Protocol Version: `404`

## What can it do
GoMCbot play a role in a Minecraft Client which has no graphy but you can control it by programing.
These're what gomcbot can do and what can't do yet
- [x] Join online server
- [x] Join offline server
- [x] Parse chunk data
- [x] View inventory
- [x] Use item
- [x] Listening sound
- [x] Send chat message
- [x] Recive and **parse** chat message
- [x] Move around and look around
- [x] Swing hands
- [x] Mine blocks
- [ ] Place blocks
- [ ] Steer vehicle
- [ ] Steer boat
- [ ] Craft
- [ ] Pick item
- [ ] Use entity
- [ ] Edit book
- [ ] Enchant item
- [ ] Record entity
- [ ] Unload chunks

> I'm writing a mine robot so this list will be complete later. Create an issue if you want one of them.
# Usage
There are some basic usage example. For more information, see [Wiki](https://github.com/Tnze/gomcbot/wiki).

To use gomcbot, simply run `go get -u github.com/Tnze/gomcbot`
## Ping and List a Minecraft Server

```go
package main

import (
	"fmt"
	"github.com/Tnze/gomcbot"
)

func main() {
	resp, err := gomcbot.PingAndList("localhost", 25565)
	if err != nil {
		panic(err)
	}
	// see format of resp at https://wiki.vg/Server_List_Ping#Response
	fmt.Println(resp)
}
```

## Get a Mojang AccessToken (Login)
```go
package main

import (
	"fmt"
	auth "github.com/Tnze/gomcbot/authenticate"
)

func main() {
	resp, err := auth.Authenticate("Your Email", "Your Password")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
```

## Join a Server (online)
```go
package main

import (
	// mb "github.com/Tnze/gomcbot"
	auth "github.com/Tnze/gomcbot/authenticate"
)

func main() {
	//Login
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

	for e := range events {//Reciving events
		switch e {
		case mb.PlayerSpawnEvent:
			fmt.Println("Player is spawn!")
		}
	}
}

```

## Join a Server (offline)
```go
package main

import (
	mb "github.com/Tnze/gomcbot"
)

func main() {
	p := mb.Auth{
		Name: "YourGameName",
		UUID: "YourUUID",
	}

	game, err := p.JoinServer("localhost", 25565)
	if err != nil {
		panic(err)
	}
	events := game.GetEvents()

	go game.HandleGame()
	for e := range events {
		switch e {
		case mb.PlayerSpawnEvent:
			fmt.Println("Player is spawn!")
		}
	}
}
```
