package gomcbot

import (
	"fmt"
	"testing"
)

func TestItemsString(t *testing.T) {
	for i := 0; i < 789+1; i++ {
		fmt.Println(Solt{ID: i, Count: byte(i%64 + 1)})
	}
}
