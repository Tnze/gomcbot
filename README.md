# GoMCbot
A golang minecraft bot package

# Usage

## PingList a minecraft server

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
	fmt.Println(resp)// for json format of resp see https://wiki.vg/Server_List_Ping#Response
}
```

## Get a mojang AccessToken (Login)
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

## Join a server (online)
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
	g, err := Auth.JoinServer("localhost", 25565)
	if err != nil {
		panic(err)
	}

	//Handle game
	err = g.HandleGame()
	if err != nil {
		panic(err)
	}
}

```

## Join a server (offline)
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
	
	err = game.HandleGame()
	if err != nil {
		panic(err)
	}
}
```
