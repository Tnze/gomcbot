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

## Join a server
> Note: gomcbot can't join a online-mode server currently.
```go
p := Player{
	Name: "Tnze",
	UUID: "58f6356e-b30c-4811-8bfc-d72a9ee99e73",
	AsTk: "",
}
game, err = p.JoinServer("localhost", 25565)
if err != nil {
	panic(err)
}
...
```
