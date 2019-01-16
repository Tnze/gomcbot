package gomcbot

//World record all of the things in the world where player at
type World struct {
	Entities map[int32]Entity
	chunks   map[ChunkLoc]*Chunk
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

//ChunkLoc record location of a Chunk
type ChunkLoc struct {
	X, Y int
}

//Entity 表示一个实体
type Entity interface {
	EntityID() int32
}
