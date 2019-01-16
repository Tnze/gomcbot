package gomcbot

import (
	pk "./packet"
	"fmt"
)

// Player includes the player's status
type Player struct {
	entityID int32
	UUID     [2]int64 //128bit UUID

	X, Y, Z    float64
	Yaw, Pitch float32
	OnGround   bool

	HeldItem  int //拿着的物品栏位
	Inventory []Solt
}

//EntityID get player's entity ID
func (p *Player) EntityID() int32 {
	return p.entityID
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
	X, Y, Z int
}

//HandleGame recive server packet and response them correct
func (g *Game) HandleGame() error {
	g.sendChan = make(chan pk.Packet, 64)
	go func() {
		for {
			p := <-g.sendChan
			err := g.sendPacket(&p)
			if err != nil {
				fmt.Printf("send packet in game fail: %v\n", err)
			}
		}
	}()
	for {

		// fmt.Println("Reading Packet")
		//Recive Packet
		pack, err := g.recvPacket()
		if err != nil {
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
			g.sendChan <- *g.settings.pack()
		case 0x3D:
			handleHeldItemPacket(g, pack)
		case 0x22:
			handleChunkDataPacket(g, pack)
		case 0x32:
			handlePlayerPositionAndLookPacket(g, pack)
			fmt.Println("Pos: ", g.player.X, g.player.Y, g.player.Z)
			sendPlayerPositionAndLookPacket(g) // to confirm the spawn position
			fmt.Println("Send PositionAndLookPacket")
		case 0x54:
			handleDeclareRecipesPacket(g, pack)
		case 0x29:
			handleEntityLookAndRelativeMove(g, pack)
		case 0x39:
			handleEntityHeadLook(g, pack)
		case 0x28:
			handleEntityRelativeMovePacket(g, pack)
		case 0x21:
			handleKeepAlivePacket(g, pack)
		case 0x27:
			handleEntityPacket(g, pack)
		case 0x05:
			handleSpawnPlayer(g, pack)
		case 0x15:
			//handleWindowItems(g, pack)
		default:
			// fmt.Printf("ignore pack id %X\n", pack.ID)
		}
		// g.player.Yaw += 1
		// if g.player.Yaw > 360 {
		// 	g.player.Yaw -= 360
		// }
		sendPlayerLookPacket(g)
	}
}

func handleJoinGamePacket(g *Game, p *pk.Packet) {
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
	// fmt.Println("Plugin Packet: ", p)
}

func handleServerDifficultyPacket(g *Game, p *pk.Packet) {
	g.Info.Difficulty = int(p.Data[0])
}

func handleSpawnPositionPacket(g *Game, p *pk.Packet) {
	g.Info.SpawnPosition.X, g.Info.SpawnPosition.Y, g.Info.SpawnPosition.Z =
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
	c, x, y := unpackChunkDataPacket(p, g.Info.Dimension == 0)
	g.world.chunks[ChunkLoc{x, y}] = c
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
	ID, index := pk.UnpackVarInt(p.Data)
	E := g.world.Entities[ID]
	if E != nil {
		P, ok := E.(*Player)
		if !ok {
			return
		}
		DeltaX := pk.UnpackInt16(p.Data[index:])
		DeltaY := pk.UnpackInt16(p.Data[index+2:])
		DeltaZ := pk.UnpackInt16(p.Data[index+4:])
		index += 3 * 2
		P.Yaw += float32(p.Data[index]) * (1.0 / 256)
		P.Pitch += float32(p.Data[index+1]) * (1.0 / 256)
		index += 2
		P.OnGround = p.Data[index] != 0x00

		P.X += float64(DeltaX) / 128
		P.Y += float64(DeltaY) / 128
		P.Z += float64(DeltaZ) / 128
	}
}

func handleEntityHeadLook(g *Game, p *pk.Packet) {

}

func handleEntityRelativeMovePacket(g *Game, p *pk.Packet) {
	ID, index := pk.UnpackVarInt(p.Data)
	E := g.world.Entities[ID]
	if E != nil {
		P, ok := E.(*Player)
		if !ok {
			return
		}
		DeltaX := pk.UnpackInt16(p.Data[index:])
		DeltaY := pk.UnpackInt16(p.Data[index+2:])
		DeltaZ := pk.UnpackInt16(p.Data[index+4:])
		index += 3 * 2
		P.OnGround = p.Data[index] != 0x00

		P.X += float64(DeltaX) / 128
		P.Y += float64(DeltaY) / 128
		P.Z += float64(DeltaZ) / 128
	}
}

func handleKeepAlivePacket(g *Game, p *pk.Packet) {
	KeepAliveID := pk.UnpackInt64(p.Data)
	sendKeepAlivePacket(g, KeepAliveID)
}

func handleEntityPacket(g *Game, p *pk.Packet) {
	// initialize an entity.
}

func handleSpawnPlayer(g *Game, p *pk.Packet) {
	np := new(Player)
	var index int
	np.entityID, index = pk.UnpackVarInt(p.Data[index:])
	np.UUID[0] = pk.UnpackInt64(p.Data[index:])
	np.UUID[1] = pk.UnpackInt64(p.Data[index+8:])
	index += 16
	np.X = pk.UnpackDouble(p.Data[index:])
	np.Y = pk.UnpackDouble(p.Data[index+8:])
	np.Z = pk.UnpackDouble(p.Data[index+16:])
	index += 3 * 8
	np.Yaw = float32(p.Data[index]) * (1.0 / 256)
	np.Pitch = float32(p.Data[index+1]) * (1.0 / 256)
	index += 2

	g.world.Entities[np.entityID] = np //把该玩家添加到全局实体表里面
}

func handleWindowItems(g *Game, p *pk.Packet) {
	index := 0

	WindowID := p.Data[index]
	index++

	Count := pk.UnpackInt16(p.Data[index:])
	index += 2

	solts := make([]Solt, Count)
	var len int
	for i := int16(0); i < Count; i++ {
		solts[i], len = unpackSolt(p.Data[index:])
		index += len
	}

	switch WindowID {
	case 0: //is player inventory
		g.player.Inventory = solts
	}
}

func sendTeleportConfirmPacket(g *Game, TeleportID int32) {
	g.sendChan <- pk.Packet{
		ID:   0x00,
		Data: pk.PackVarInt(TeleportID),
	}
}

func sendPlayerPositionAndLookPacket(g *Game) {
	var data []byte
	data = append(data, pk.PackDouble(g.player.X)...)
	data = append(data, pk.PackDouble(g.player.Y)...)
	data = append(data, pk.PackDouble(g.player.Z)...)
	data = append(data, pk.PackFloat(g.player.Yaw)...)
	data = append(data, pk.PackFloat(g.player.Pitch)...)
	data = append(data, pk.PackBoolean(g.player.OnGround))

	g.sendChan <- pk.Packet{
		ID:   0x11,
		Data: data,
	}
}

func sendKeepAlivePacket(g *Game, KeepAliveID int64) {
	g.sendChan <- pk.Packet{
		ID:   0x0E,
		Data: pk.PackUint64(uint64(KeepAliveID)),
	}
	fmt.Println("Keep Alive")
}

func sendPlayerLookPacket(g *Game) {
	var data []byte
	data = append(data, pk.PackFloat(g.player.Yaw)...)
	data = append(data, pk.PackFloat(g.player.Pitch)...)
	data = append(data, pk.PackBoolean(g.player.OnGround))
	g.sendChan <- pk.Packet{
		ID:   0x12,
		Data: data,
	}
}

//GetWorld return the current World
func (g *Game) GetWorld() World {
	return g.world
}
