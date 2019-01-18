package gomcbot

import (
	"bytes"
	"fmt"
	pk "github.com/Tnze/gomcbot/packet"
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

		reader := bytes.NewReader(pack.Data)

		switch pack.ID {
		case 0x25:
			handleJoinGamePacket(g, reader)
		case 0x19:
			handlePluginPacket(g, reader)
		case 0x0D:
			handleServerDifficultyPacket(g, reader)
		case 0x49:
			handleSpawnPositionPacket(g, reader)
		case 0x2E:
			handlePlayerAbilitiesPacket(g, reader)
			g.sendChan <- *g.settings.pack()
		case 0x3D:
			handleHeldItemPacket(g, reader)
		case 0x22:
			handleChunkDataPacket(g, pack)
		case 0x32:
			handlePlayerPositionAndLookPacket(g, reader)
			fmt.Println("Pos: ", g.player.X, g.player.Y, g.player.Z)
			sendPlayerPositionAndLookPacket(g) // to confirm the spawn position
			fmt.Println("Send PositionAndLookPacket")
		case 0x54:
			handleDeclareRecipesPacket(g, reader)
		case 0x29:
			handleEntityLookAndRelativeMove(g, reader)
		case 0x39:
			handleEntityHeadLook(g, reader)
		case 0x28:
			handleEntityRelativeMovePacket(g, reader)
		case 0x21:
			handleKeepAlivePacket(g, reader)
		case 0x27:
			handleEntityPacket(g, reader)
		case 0x05:
			handleSpawnPlayer(g, reader)
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

func handleJoinGamePacket(g *Game, r *bytes.Reader) error {
	eid, err := pk.UnpackInt32(r)
	if err != nil {
		return err
	}
	g.Info.EntityID = int(eid)
	gamemode, err := r.ReadByte()
	if err != nil {
		return err
	}
	g.Info.Gamemode = int(gamemode & 0x7)
	g.Info.Hardcore = gamemode&0x8 != 0
	dimension, err := pk.UnpackInt32(r)
	if err != nil {
		return err
	}
	g.Info.Dimension = int(dimension)
	difficulty, err := r.ReadByte()
	if err != nil {
		return err
	}
	g.Info.Difficulty = int(difficulty)
	// ignore Max Players
	g.Info.LevelType, err = pk.UnpackString(r)
	if err != nil {
		return err
	}
	rdi, err := r.ReadByte()
	if err != nil {
		return err
	}
	g.Info.ReducedDebugInfo = rdi != 0x00
	return nil
}

func handlePluginPacket(g *Game, r *bytes.Reader) {
	// fmt.Println("Plugin Packet: ", p)
}

func handleServerDifficultyPacket(g *Game, r *bytes.Reader) error {
	diff, err := r.ReadByte()
	if err != nil {
		return err
	}
	g.Info.Difficulty = int(diff)
	return nil
}

func handleSpawnPositionPacket(g *Game, r *bytes.Reader) (err error) {
	g.Info.SpawnPosition.X, g.Info.SpawnPosition.Y, g.Info.SpawnPosition.Z, err =
		pk.UnpackPosition(r)
	return
}

func handlePlayerAbilitiesPacket(g *Game, r *bytes.Reader) error {
	f, err := r.ReadByte()
	if err != nil {
		return err
	}
	g.abilities.Flags = int8(f)
	g.abilities.FlyingSpeed, err = pk.UnpackFloat(r)
	if err != nil {
		return err
	}
	g.abilities.FieldofViewModifier, err = pk.UnpackFloat(r)
	return err
}

func handleHeldItemPacket(g *Game, r *bytes.Reader) error {
	hi, err := r.ReadByte()
	if err != nil {
		return err
	}
	g.player.HeldItem = int(hi)
	return nil
}

func handleChunkDataPacket(g *Game, p *pk.Packet) error {
	c, x, y, err := unpackChunkDataPacket(p, g.Info.Dimension == 0)
	g.world.chunks[ChunkLoc{x, y}] = c
	return err
}

func handlePlayerPositionAndLookPacket(g *Game, r *bytes.Reader) error {
	x, err := pk.UnpackDouble(r)
	if err != nil {
		return err
	}
	y, err := pk.UnpackDouble(r)
	if err != nil {
		return err
	}
	z, err := pk.UnpackDouble(r)
	if err != nil {
		return err
	}
	yaw, err := pk.UnpackFloat(r)
	if err != nil {
		return err
	}
	pitch, err := pk.UnpackFloat(r)
	if err != nil {
		return err
	}

	flags, err := r.ReadByte()
	if err != nil {
		return err
	}

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
	TeleportID, _ := pk.UnpackVarInt(r)
	sendTeleportConfirmPacket(g, TeleportID)
	return nil
}

func handleDeclareRecipesPacket(g *Game, r *bytes.Reader) {
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

func handleEntityLookAndRelativeMove(g *Game, r *bytes.Reader) error {
	ID, err := pk.UnpackVarInt(r)
	if err != nil {
		return err
	}
	E := g.world.Entities[ID]
	if E != nil {
		P, ok := E.(*Player)
		if !ok {
			return nil
		}
		DeltaX, err := pk.UnpackInt16(r)
		if err != nil {
			return err
		}
		DeltaY, err := pk.UnpackInt16(r)
		if err != nil {
			return err
		}
		DeltaZ, err := pk.UnpackInt16(r)
		if err != nil {
			return err
		}

		yaw, err := r.ReadByte()
		if err != nil {
			return err
		}

		pitch, err := r.ReadByte()
		if err != nil {
			return err
		}
		P.Yaw += float32(yaw) * (1.0 / 256)
		P.Pitch += float32(pitch) * (1.0 / 256)

		og, err := r.ReadByte()
		if err != nil {
			return err
		}
		P.OnGround = og != 0x00

		P.X += float64(DeltaX) / 128
		P.Y += float64(DeltaY) / 128
		P.Z += float64(DeltaZ) / 128
	}
	return nil
}

func handleEntityHeadLook(g *Game, r *bytes.Reader) {

}

func handleEntityRelativeMovePacket(g *Game, r *bytes.Reader) error {
	ID, err := pk.UnpackVarInt(r)
	if err != nil {
		return err
	}
	E := g.world.Entities[ID]
	if E != nil {
		P, ok := E.(*Player)
		if !ok {
			return nil
		}
		DeltaX, err := pk.UnpackInt16(r)
		if err != nil {
			return err
		}
		DeltaY, err := pk.UnpackInt16(r)
		if err != nil {
			return err
		}
		DeltaZ, err := pk.UnpackInt16(r)
		if err != nil {
			return err
		}

		og, err := r.ReadByte()
		if err != nil {
			return err
		}
		P.OnGround = og != 0x00

		P.X += float64(DeltaX) / 128
		P.Y += float64(DeltaY) / 128
		P.Z += float64(DeltaZ) / 128
	}
	return nil
}

func handleKeepAlivePacket(g *Game, r *bytes.Reader) (err error) {
	KeepAliveID, err := pk.UnpackInt64(r)
	sendKeepAlivePacket(g, KeepAliveID)
	return
}

func handleEntityPacket(g *Game, r *bytes.Reader) {
	// initialize an entity.
}

func handleSpawnPlayer(g *Game, r *bytes.Reader) (err error) {
	np := new(Player)
	np.entityID, err = pk.UnpackVarInt(r)
	if err != nil {
		return
	}
	np.UUID[0], err = pk.UnpackInt64(r)
	if err != nil {
		return
	}
	np.UUID[1], err = pk.UnpackInt64(r)
	if err != nil {
		return
	}
	np.X, err = pk.UnpackDouble(r)
	if err != nil {
		return
	}
	np.Y, err = pk.UnpackDouble(r)
	if err != nil {
		return
	}
	np.Z, err = pk.UnpackDouble(r)
	if err != nil {
		return
	}

	yaw, err := r.ReadByte()
	if err != nil {
		return err
	}

	pitch, err := r.ReadByte()
	if err != nil {
		return err
	}

	np.Yaw = float32(yaw) * (1.0 / 256)
	np.Pitch = float32(pitch) * (1.0 / 256)

	g.world.Entities[np.entityID] = np //把该玩家添加到全局实体表里面
	return nil
}

func handleWindowItems(g *Game, r *bytes.Reader) (err error) {
	WindowID, err := r.ReadByte()
	if err != nil {
		return
	}

	Count, err := pk.UnpackInt16(r)
	if err != nil {
		return
	}

	solts := make([]Solt, Count)
	for i := int16(0); i < Count; i++ {
		solts[i], err = unpackSolt(r)
		if err != nil {
			return
		}
	}

	switch WindowID {
	case 0: //is player inventory
		g.player.Inventory = solts
	}
	return nil
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
