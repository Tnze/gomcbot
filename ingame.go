package gomcbot

import (
	pk "./packet"
	"fmt"
)

// Player includes the player's status
type Player struct {
	HeldItem   int     //拿着的物品栏位
	X, Y, Z    float64 //Player Position
	Yaw, Pitch float32
}

//PlayerInfo content player info in server
type PlayerInfo struct {
	EntityID         int      //实体ID
	Gamemode         int      //游戏模式
	Hardcore         bool     //是否是极限模式
	Dimension        int      //维度
	Difficulty       int      //难度
	LevelType        string   //地图类型
	ReducedDebugInfo bool     //减少调试信息
	SpawnPosition    Position //主世界出生点
}

// PlayerAbilities defines what player can do
type PlayerAbilities struct {
	Flags               int8
	FlyingSpeed         float32
	FieldofViewModifier float32
}

//Position is a 3D vector
type Position struct {
	x, y, z int
}

//HandleGame recive server packet and response them correct
func (g *Game) HandleGame() error {
	g.sendChan = make(chan pk.Packet, 64)
	for {
		if len(g.sendChan) > 0 { //如果有包要发送则发送包
			p := <-g.sendChan
			err := g.sendPacket(&p)
			if err != nil {
				return fmt.Errorf("send packet in game fail: %v", err)
			}
		}

		//Recive Packet
		pack, err := g.recvPacket()
		if err != nil {
			// fmt.Printf("recv packet in game fail: %v\n", err)
			// continue
			return fmt.Errorf("recv packet in game fail: %v", err)
		}

		switch pack.ID {
		case 0x25:
			handleJoinGamePacket(g, pack)
		case 0x19:
			handlePluginPacket(g, pack)
		case 0x0D:
			handleServerDifficultyPacket(g, pack)
		case 0x49:
			handleSpawnPositionPacket(g, pack)
		case 0x2E:
			handlePlayerAbilitiesPacket(g, pack)
		case 0x3D:
			handleHeldItemPacket(g, pack)
		case 0x22:
			handleChunkDataPacket(g, pack)
		case 0x32:
			handlePlayerPositionAndLookPacket(g, pack)
		case 0x54:
			handleDeclareRecipesPacket(g, pack)
		case 0x29:
			handleEntityLookAndRelativeMove(g, pack)
		case 0x39:
			handleEntityHeadLook(g, pack)
		case 0x28:
			handleEntityRelativeMove(g, pack)
		default:
			fmt.Printf("ignore pack id %X\n", pack.ID)
		}
	}
}

func handleJoinGamePacket(g *Game, p *pk.Packet) {
	// https://wiki.vg/Protocol#Join_Game
	g.Info.EntityID = int(pk.UnpackInt32(p.Data))
	gamemode := p.Data[1]
	g.Info.Gamemode = int(gamemode & 0x7)
	g.Info.Hardcore = gamemode&0x8 != 0
	g.Info.Dimension = int(pk.UnpackInt32(p.Data[5:]))
	g.Info.Difficulty = int(p.Data[9])
	// ignore Max Players
	var ltlen int
	g.Info.LevelType, ltlen = pk.UnpackString(p.Data[11:])
	g.Info.ReducedDebugInfo = p.Data[11+ltlen] != 0x00
}

func handlePluginPacket(g *Game, p *pk.Packet) {
	fmt.Println("ignore Plugin Packet")
	// fmt.Println("Plugin Packet: ", p)
}

func handleServerDifficultyPacket(g *Game, p *pk.Packet) {
	g.Info.Difficulty = int(p.Data[0])
}

func handleSpawnPositionPacket(g *Game, p *pk.Packet) {
	g.Info.SpawnPosition.x, g.Info.SpawnPosition.y, g.Info.SpawnPosition.z =
		pk.UnpackPosition(p.Data)
}

func handlePlayerAbilitiesPacket(g *Game, p *pk.Packet) {
	g.abilities.Flags = int8(p.Data[0])
	g.abilities.FlyingSpeed = pk.UnpackFloat(p.Data[1:])
	g.abilities.FieldofViewModifier = pk.UnpackFloat(p.Data[5:])
}

func handleHeldItemPacket(g *Game, p *pk.Packet) {
	g.player.HeldItem = int(p.Data[0])
}

func handleChunkDataPacket(g *Game, p *pk.Packet) {

}

func handlePlayerPositionAndLookPacket(g *Game, p *pk.Packet) {
	x := pk.UnpackDouble(p.Data)
	y := pk.UnpackDouble(p.Data[8:])
	z := pk.UnpackDouble(p.Data[16:])
	yaw := pk.UnpackFloat(p.Data[24:])
	pitch := pk.UnpackFloat(p.Data[28:])

	flags := p.Data[32]

	if flags&0x01 == 0 {
		g.player.X = x
	} else {
		g.player.X += x
	}
	if flags&0x02 == 0 {
		g.player.Y = y
	} else {
		g.player.Y += y
	}
	if flags&0x04 == 0 {
		g.player.Z = z
	} else {
		g.player.Z += z
	}
	if flags&0x08 == 0 {
		g.player.Yaw = yaw
	} else {
		g.player.Yaw += yaw
	}
	if flags&0x10 == 0 {
		g.player.Pitch = pitch
	} else {
		g.player.Pitch += pitch
	}
	//confirm this packet with Teleport Confirm
	TeleportID, _ := pk.UnpackVarInt(p.Data[33:])
	sendTeleportConfirmPacket(g, TeleportID)
}

func handleDeclareRecipesPacket(g *Game, p *pk.Packet) {
	//Ignore Declare Recipes Packet

	// NumRecipes, index := pk.UnpackVarInt(p.Data)
	// for i := 0; i < int(NumRecipes); i++ {
	// 	RecipeID, len := pk.UnpackString(p.Data[index:])
	// 	index += len
	// 	Type, len := pk.UnpackString(p.Data[index:])
	// 	index += len
	// 	switch Type {
	// 	case "crafting_shapeless":
	// 	}
	// }
}

func handleEntityLookAndRelativeMove(g *Game, p *pk.Packet) {

}

func handleEntityHeadLook(g *Game, p *pk.Packet) {

}

func handleEntityRelativeMove(g *Game, p *pk.Packet) {

}

func sendTeleportConfirmPacket(g *Game, TeleportID int32) {
	g.sendChan <- pk.Packet{
		ID:   0x00,
		Data: pk.PackVarInt(TeleportID),
	}
}
