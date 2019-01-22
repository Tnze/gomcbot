package gomcbot

import (
	"fmt"
	"testing"
)

func TestBlockString(t *testing.T) {
	for i := uint(0); i < 8598+1; i++ {
		fmt.Println(Block{id: i})
	}
}
