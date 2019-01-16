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

## Join a server
```go
package main

import (
	"fmt"
	gmcb "github.com/Tnze/gomcbot"
)

func main() {
	p := Auth{
		Name: "Mi_Xi_Xi",
		UUID: "ff7a038f-265c-4d42-b0cf-04c575896469",
		AsTk: "",
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
