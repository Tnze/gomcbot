package gomcbot

//World record all of the things in the world where player at
type World struct {
	Entities map[int32]Entity
	chunks   map[chunkLoc]*Chunk
}

//Chunk store a 256*16*16 clolumn blocks
type Chunk struct {
	sections [16]Section
}

//Section store a 16*16*16 cube blocks
type Section struct {
	blocks [16][16][16]Block
}

//Block is the base of world
type Block struct {
	id uint
}

type chunkLoc struct {
	X, Y int
}

//Entity 表示一个实体
type Entity interface {
	EntityID() int32
}

// GetBlock return the block in the position (x, y, z)
func (w *World) GetBlock(x, y, z int) Block {
	c := w.chunks[chunkLoc{x >> 4, z >> 4}]
	if c != nil {
		cx, cy, cz := x&(16-1), y&(16-1), z&(16-1)
		/*
			n = n&(16-1)

			is equal to

			n %= 16
			if n < 0 { n += 16 }
		*/

		return c.sections[y/16].blocks[cx][cy][cz]
	}

	return Block{id: 0}
}
