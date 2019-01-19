package gomcbot

import (
	"math"
)

// SetPosition method move your character around.
// if delta is too big that the server will ignore.
func (g *Game) SetPosition(x, y, z float64, onGround bool) {
	g.motion <- func() {
		g.player.X, g.player.Y, g.player.Z = x, y, z
		g.player.OnGround = onGround
		sendPlayerPositionPacket(g) //向服务器更新位置
	}
}

// LookAt method turn player's hand and make it look at a point.
func (g *Game) LookAt(x, y, z float64) {
	x0, y0, z0 := g.player.X, g.player.Y, g.player.Z
	x, y, z = x-x0, y-y0, z-z0

	r := math.Sqrt(x*x + y*y + z*z)
	yaw := -math.Atan2(x, z) / math.Pi * 180
	for yaw < 0 {
		yaw = 360 + yaw
	}
	pitch := -math.Asin(y/r) / math.Pi * 180

	g.LookYawPitch(float32(yaw), float32(pitch))
}

// LookYawPitch set player's hand to the direct by yaw and pitch.
// yaw can be [0, 360) and pitch can be (-180, 180).
// if |pitch|>90 the player's hand will be very strange.
func (g *Game) LookYawPitch(yaw, pitch float32) {
	g.motion <- func() {
		g.player.Yaw, g.player.Pitch = yaw, pitch
		sendPlayerLookPacket(g) //向服务器更新朝向
	}
}
