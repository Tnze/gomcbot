package gomcbot

import (
	pk "./packet"
	"fmt"
)

type World struct {
	Entities map[int32]Entity
}

type Chunk struct {
}

type Section struct {
}

type Block struct {
}

//Entity 表示一个实体
type Entity interface {
	EntityID() int32
}

func unpackChunkDataPacket(p *pk.Packet) {
	// index := 0
	// X := pk.UnpackInt32(p.Data)
	// Y := pk.UnpackInt32(p.Data[4:])
	// index += 8
	// fmt.Println("Chunk: (", X, ", ", Y, ")") //Debug: Show Chunk loc
	// FullChunk := p.Data[index] != 0x00
	// index++
	// PrimaryBitMask, len := pk.UnpackVarInt(p.Data[index:])
	// index += len
	// Size, len := pk.UnpackVarInt(p.Data[index:])
	// index += len
	// Data := p.Data[index : index+int(Size)]
	// index += int(Size)
	// /*NumberofBlockEntities*/ _, len = pk.UnpackVarInt(p.Data[index:])
	// index += len

	// readChunkColumn(FullChunk, PrimaryBitMask, Data)
	// fmt.Println(FullChunk, PrimaryBitMask, Data, NumberofBlockEntities)
}

func readChunkColumn(isFull bool, mask int32, data []byte) (c *Chunk) {
	index := 0
	for sectionY := 0; sectionY < 16; sectionY++ {
		if (mask & (1 << uint(sectionY))) != 0 { // Is the given bit set in the mask?
			BitsPerBlock := data[index]
			index++
			fmt.Println("If use section palettes:", sectionY, BitsPerBlock)

			if BitsPerBlock < 9 { //section palettes is used
				PaletteLength, len := pk.UnpackVarInt(data[index:])
				index += len

				Palette := make([]int32, PaletteLength) //Read Palette
				for i := 0; i < int(PaletteLength); i++ {
					Palette[i], len = pk.UnpackVarInt(data[index:])
					index += len
				}
			}

			DataArrayLength, len := pk.UnpackVarInt(data[index:])
			index += len
			fmt.Println("DataArrayLength", DataArrayLength)
			DataArray := make([]int64, DataArrayLength)
			for i := 0; i < int(DataArrayLength); i++ {
				DataArray[i] = pk.UnpackInt64(data[index:])
				index += 8
			}

			//Read BlockLight (ignore)
			index += int(DataArrayLength) / 2

			index += int(DataArrayLength) / 2
		}
	}
	if isFull { //need recive Biomes datas
		index += 256 * 4
	}
	return
}
