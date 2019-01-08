package gomcbot

import (
	pk "./packet"
)

type World struct {
}

type Chunk struct {
}

type Section struct {
}

type Block struct {
}

func unpackChunkDataPacket(p *pk.Packet) {
	// X, index := pk.UnpackVarInt(p.Data)
	// Y, len := pk.UnpackVarInt(p.Data[index:])
	// index += len
	// FullChunk := p.Data[index] != 0x00
	// index++
	// PrimaryBitMask, len := pk.UnpackVarInt(p.Data[index:])
	// index += len
	// Size, len := pk.UnpackVarInt(p.Data[index:])
	// index += len
	// Data := p.Data[index : index+int(Size)]
	// index += int(Size)
	// NumberofBlockEntities, len := pk.UnpackVarInt(p.Data[index:])
	// index += len
}
