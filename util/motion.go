package util

import (
	bot "github.com/Tnze/gomcbot"
)

// CalibratePos moves player to the centre of block and fall player on the ground
func CalibratePos(g *bot.Game) {
	p := g.GetPlayer()
	x, y, z := p.GetBlockPos()
	for g.GetBlock(x, y-1, z).String() == "minecraft:air" {
		y--
	}
	g.SetPosition(float64(x)+0.5, float64(y), float64(z)+0.5, true)
}
