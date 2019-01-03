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
	fmt.Println(resp)
}
```
