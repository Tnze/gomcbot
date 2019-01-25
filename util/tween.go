package util

import (
	"fmt"
	bot "github.com/Tnze/gomcbot"
	"math"
	"time"
)

// TweenLook do tween animation at player's head.
func TweenLook(g *bot.Game, yaw, pitch float32, t time.Duration) {
	p := g.GetPlayer()
	start := time.Now()
	yaw0, pitch0 := p.Yaw, p.Pitch
	ofstY, ofstP := yaw-yaw0, pitch-pitch0
	var scale float32
	for scale < 1 {
		scale = float32(time.Since(start)) / float32(t)
		g.LookYawPitch(yaw0+ofstY*scale, pitch0+ofstP*scale)
		time.Sleep(time.Millisecond * 50)
	}
}

//TweenLineMove allows you smoothly move on plane. You can't move in Y axis
func TweenLineMove(g *bot.Game, x, z float64) error {
	p := g.GetPlayer()
	start := time.Now()
	x0, y0, z0 := p.GetPosition()

	if similar(x0, x) && similar(z0, z) {
		return nil
	}

	y0 = math.Floor(y0)
	ofstX, ofstZ := x-x0, z-z0
	t := time.Duration(float64(time.Second) * (math.Sqrt(ofstX*ofstX+ofstZ*ofstZ) / 4.2))
	var scale float64
	for scale < 1 {
		scale = float64(time.Since(start)) / float64(t)
		g.SetPosition(x0+ofstX*scale, y0, z0+ofstZ*scale, true)
		time.Sleep(time.Millisecond * 50)
	}

	p = g.GetPlayer()
	if !similar(p.X, x) || !similar(p.Z, z) {
		return fmt.Errorf("wrongly move")
	}
	return nil
}

func similar(a, b float64) bool {
	return a-b < 1 && b-a < 1
}

//TweenJump simulate player jump make no headway
func TweenJump(g *bot.Game) {
	p := g.GetPlayer()
	y := math.Floor(p.Y)
	for tick := 0; tick < 11; tick++ {
		h := -1.7251e-8 + 0.4591*float64(tick) - 0.0417*float64(tick)*float64(tick)

		g.SetPosition(p.X, y+h, p.Z, false)
		time.Sleep(time.Millisecond * 50)
	}
	g.SetPosition(p.X, p.Y, p.Z, true)
}

//TweenJumpTo simulate player jump up a block
func TweenJumpTo(g *bot.Game, x, z int) {
	p := g.GetPlayer()
	y := math.Floor(p.Y)
	for tick := 0; tick < 7; tick++ {
		h := -1.7251e-8 + 0.4591*float64(tick) - 0.0417*float64(tick)*float64(tick)

		g.SetPosition(p.X, y+h, p.Z, false)
		time.Sleep(time.Millisecond * 50)
	}
	TweenLineMove(g, float64(x)+0.5, float64(z)+0.5)
	CalibratePos(g)
}
