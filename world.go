package gomcbot

type World struct {
	Entities map[int32]Entity
	chunks   map[Location]*Chunk
}

type Chunk struct {
	sections [16]Section
}

type Section struct {
	blocks [16][16][16]Block
}

type Block struct {
	id uint
}

type Location struct {
	X, Y int
}

//Entity 表示一个实体
type Entity interface {
	EntityID() int32
}
