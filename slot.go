package gomcbot

import (
	pk "./packet"
	// "bytes"
	// "github.com/Nightgunner5/go.nbt"
)

//Solt 物品栏的一格
type Solt struct {
	ID    int
	Count byte
}

type soltNBT struct {
}

func unpackSolt(b []byte) (Solt, int) {
	index := 0
	Present := b[index] != 0x00
	index++
	if Present {
		itemID, len := pk.UnpackVarInt(b[index:])
		index += len
		count := b[index]
		index++

		//nbt.Unmarshal(nbt.Uncompressed)

		return Solt{
			ID:    int(itemID),
			Count: count,
		}, index
	}
	return Solt{}, index
}
