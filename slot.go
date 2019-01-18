package gomcbot

import (
	"bytes"
	pk "github.com/Tnze/gomcbot/packet"
	// "github.com/Nightgunner5/go.nbt"
)

//Solt 表示物品栏的一格
type Solt struct {
	ID    int
	Count byte
}

type soltNBT struct {
}

func unpackSolt(b *bytes.Reader) (Solt, error) {
	index := 0
	p, err := b.ReadByte()
	if err != nil {
		return Solt{}, err
	}
	Present := p != 0x00
	index++
	if Present {
		itemID, err := pk.UnpackVarInt(b)
		if err != nil {
			return Solt{}, err
		}
		count, err := b.ReadByte()
		if err != nil {
			return Solt{}, err
		}
		index++

		//nbt.Unmarshal(nbt.Uncompressed)

		return Solt{
			ID:    int(itemID),
			Count: count,
		}, nil
	}
	return Solt{}, nil
}
