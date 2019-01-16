package gomcbot

import (
	pk "./packet"
	// "fmt"
)

func unpackChunkDataPacket(p *pk.Packet, hasSkyLight bool) (c *Chunk, x, y int) {
	index := 0

	//区块坐标
	X := pk.UnpackInt32(p.Data)
	Y := pk.UnpackInt32(p.Data[4:])
	index += 8
	// fmt.Println("Chunk: (", X, ", ", Y, ")") //Debug: Show Chunk loc

	FullChunk := p.Data[index] != 0x00
	index++

	//主掩码
	PrimaryBitMask, len := pk.UnpackVarInt(p.Data[index:])
	index += len

	//区块数据
	Size, len := pk.UnpackVarInt(p.Data[index:])
	index += len
	Data := p.Data[index : index+int(Size)]
	index += int(Size)

	//实体信息
	// NumberofBlockEntities, len := pk.UnpackVarInt(p.Data[index:])
	// index += len

	//解析区块数据
	return readChunkColumn(FullChunk, PrimaryBitMask, Data, hasSkyLight),
		int(X), int(Y)
}

func readChunkColumn(isFull bool, mask int32, data []byte, hasSkyLight bool) *Chunk {
	index := 0
	var c Chunk
	for sectionY := 0; sectionY < 16; sectionY++ {
		if (mask & (1 << uint(sectionY))) != 0 { // Is the given bit set in the mask?
			BitsPerBlock := data[index]
			index++

			//读调色板
			palette := make(map[uint]uint)
			if BitsPerBlock < 9 {
				length, len := pk.UnpackVarInt(data[index:])
				index += len
				for id := uint(0); id < uint(length); id++ {
					stateID, len := pk.UnpackVarInt(data[index:])
					index += len

					palette[id] = uint(stateID)
				}
			}

			//Section数据
			DataArrayLength, len := pk.UnpackVarInt(data[index:])
			index += len
			DataArray := make([]int64, DataArrayLength)
			for i := 0; i < int(DataArrayLength); i++ {
				DataArray[i] = pk.UnpackInt64(data[index:])
				index += 8
			}

			//用数据填充区块
			fillSection(&c.sections[sectionY], perBits(BitsPerBlock), DataArray, palette)

			//throw BlockLight data
			index += 2048
			if hasSkyLight {
				//throw SkyLight data
				index += 2048
			}
		}
	}
	if isFull { //need recive Biomes datas
		index += 256 * 4
	}

	// fmt.Println(c)
	return &c
}

const defaultBitsPerBlock = 14

func perBits(BitsPerBlock byte) uint {
	switch {
	case BitsPerBlock <= 4:
		return 4
	case BitsPerBlock < 9:
		return uint(BitsPerBlock)
	default:
		return defaultBitsPerBlock
	}
}

func fillSection(s *Section, bpb uint, DataArray []int64, palette map[uint]uint) {
	perOffset := 64 - bpb

	for n := 0; n < 16*16*16; n++ {
		offset := uint(n * int(bpb))
		data := uint(DataArray[offset/64])
		data <<= offset % 64
		data >>= perOffset
		if offset%64 > 64-bpb {
			data &= uint(DataArray[offset/64+1]) >>
				(bpb + offset%64 - 64)
		}
		if bpb < 9 {
			s.blocks[n%16][n%(16*16)/16][n/(16*16)].id = palette[data]
		} else {
			s.blocks[n%16][n%(16*16)/16][n/(16*16)].id = data
		}
	}
}
