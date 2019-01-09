package gomcbot

import (
	pk "./packet"
	"testing"
)

func TestUnpackChunk(t *testing.T) {
	p := pk.Packet{
		ID:   34,
		Data: []byte{},
	}
	unpackChunkDataPacket(&p)
}
