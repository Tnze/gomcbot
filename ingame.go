package gomcbot

import (
	"bytes"
	"fmt"
	pk "github.com/Tnze/gomcbot/packet"
	"math"
	"time"
)

// Player includes the player's status.
type Player struct {
	entityID int32
	UUID     [2]int64 //128bit UUID

	X, Y, Z    float64
	Yaw, Pitch float32
	OnGround   bool

	HeldItem  int //拿着的物品栏位
	Inventory []Solt

	Health         float32 //血量
	Food           int32   //饱食度
	FoodSaturation float32 //食物饱和度
}

//EntityID get player's entity ID.
func (p *Player) EntityID() int32 {
	return p.entityID
}

//GetPosition return the player's position
func (p *Player) GetPosition() (x, y, z float64) {
	return p.X, p.Y, p.Z
}

//GetBlockPos return the position of the Block at player's feet
func (p *Player) GetBlockPos() (x, y, z int) {
	return int(math.Floor(p.X)), int(math.Floor(p.Y)), int(math.Floor(p.Z))
}

//PlayerInfo content player info in server.
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

// PlayerAbilities defines what player can do.
type PlayerAbilities struct {
	Flags               int8
	FlyingSpeed         float32
	FieldofViewModifier float32
}

//Position is a 3D vector.
type Position struct {
	X, Y, Z int
}

// HandleGame recive server packet and response them correctly.
// Note that HandleGame will block if you don't recive from Events.
func (g *Game) HandleGame() error {
	defer func() {
		close(g.events)
	}()

	errChan := make(chan error)

	//send
	g.sendChan = make(chan pk.Packet, 64)
	go func() {
		for p := range g.sendChan {
			err := g.sendPacket(&p)
			if err != nil {
				errChan <- fmt.Errorf("send packet in game fail: %v", err)
				return
			}
		}
	}()

	//recv
	g.recvChan = make(chan *pk.Packet, 64)
	go func() {
		for {
			pack, err := g.recvPacket()
			if err != nil {
				close(g.recvChan)
				errChan <- fmt.Errorf("recv packet in game fail: %v", err)
				return
			}

			g.recvChan <- pack
		}
	}()

	for {
		select {
		case err := <-errChan:
			close(g.sendChan)
			return err
		case pack, ok := <-g.recvChan:
			if !ok {
				break
			}
			err := handlePack(g, pack)
			if err != nil {
				// fmt.Println(fmt.Errorf("handle packet 0x%X error: %v", pack.ID, err))
				return fmt.Errorf("handle packet 0x%X error: %v", pack.ID, err)
			}
		case motion := <-g.motion:
			motion()
		}
	}
}

func handlePack(g *Game, p *pk.Packet) (err error) {
	reader := bytes.NewReader(p.Data)

	switch p.ID {
	case 0x25:
		err = handleJoinGamePacket(g, reader)
	case 0x19:
		handlePluginPacket(g, reader)
	case 0x0D:
		err = handleServerDifficultyPacket(g, reader)
	case 0x49:
		err = handleSpawnPositionPacket(g, reader)
	case 0x2E:
		err = handlePlayerAbilitiesPacket(g, reader)
		g.sendChan <- *g.settings.pack()
	case 0x3D:
		err = handleHeldItemPacket(g, reader)
	case 0x22:
		err = handleChunkDataPacket(g, p)
		g.events <- BlockChangeEvent{}
	case 0x32:
		err = handlePlayerPositionAndLookPacket(g, reader)
		sendPlayerPositionAndLookPacket(g) // to confirm the spawn position
	case 0x54:
		// handleDeclareRecipesPacket(g, reader)
	case 0x29:
		// err = handleEntityLookAndRelativeMove(g, reader)
	case 0x39:
		// handleEntityHeadLookPacket(g, reader)
	case 0x28:
		// err = handleEntityRelativeMovePacket(g, reader)
	case 0x21:
		err = handleKeepAlivePacket(g, reader)
	case 0x27:
		//handleEntityPacket(g, reader)
	case 0x05:
		// err = handleSpawnPlayerPacket(g, reader)
	case 0x15:
		err = handleWindowItemsPacket(g, reader)
	case 0x44:
		err = handleUpdateHealthPacket(g, reader)
	case 0x0E:
		err = handleChatMessagePacket(g, reader)
	case 0x0B:
		err = handleBlockChangePacket(g, reader)
		g.events <- BlockChangeEvent{}
	case 0x0F:
		err = handleMultiBlockChangePacket(g, reader)
		g.events <- BlockChangeEvent{}
	case 0x1B:
		// should assumes that the server has already closed the connection by the time the packet arrives.
		g.events <- DisconnectEvent{Text: "disconnect"}
		err = fmt.Errorf("disconnect")
	case 0x17:
	// 	err = handleSetSlotPacket(g, reader)
	case 0x4D:
		err = handleSoundEffect(g, reader)
	default:
		// fmt.Printf("ignore pack id %X\n", p.ID)
	}
	return
}

func handleSoundEffect(g *Game, r *bytes.Reader) error {
	SoundID, err := pk.UnpackVarInt(r)
	if err != nil {
		return err
	}
	SoundCategory, err := pk.UnpackVarInt(r)
	if err != nil {
		return err
	}

	x, err := pk.UnpackInt32(r)
	if err != nil {
		return err
	}
	y, err := pk.UnpackInt32(r)
	if err != nil {
		return err
	}
	z, err := pk.UnpackInt32(r)
	if err != nil {
		return err
	}
	Volume, err := pk.UnpackFloat(r)
	if err != nil {
		return err
	}
	Pitch, err := pk.UnpackFloat(r)
	if err != nil {
		return err
	}
	g.events <- SoundEffectEvent{SoundID, SoundCategory, float64(x) / 8, float64(y) / 8, float64(z) / 8, Volume, Pitch}

	return nil
}

func handleSetSlotPacket(g *Game, r *bytes.Reader) error {
	windowID, err := r.ReadByte()
	if err != nil {
		return err
	}
	slot, err := pk.UnpackInt16(r)
	if err != nil {
		return err
	}
	slotData, err := unpackSolt(r)
	if err != nil {
		return err
	}

	switch int8(windowID) {
	case 0:
		if slot < 32 || slot > 44 {
			// return fmt.Errorf("slot out of range")
			break
		}
		fallthrough
	case -2:
		g.player.Inventory[slot] = slotData
		g.events <- InventoryChangeEvent(slot)
	}
	return nil
}

func handleMultiBlockChangePacket(g *Game, r *bytes.Reader) error {
	if !g.settings.ReciveMap {
		return nil
	}

	cX, err := pk.UnpackInt32(r)
	if err != nil {
		return err
	}
	cY, err := pk.UnpackInt32(r)
	if err != nil {
		return err
	}

	c := g.wd.chunks[chunkLoc{int(cX), int(cY)}]
	if c != nil {
		RecordCount, err := pk.UnpackVarInt(r)
		if err != nil {
			return err
		}

		for i := int32(0); i < RecordCount; i++ {
			xz, err := r.ReadByte()
			if err != nil {
				return err
			}
			y, err := r.ReadByte()
			if err != nil {
				return err
			}
			BlockID, err := pk.UnpackVarInt(r)
			if err != nil {
				return err
			}
			x, z := xz>>4, xz&0x0F

			c.sections[y/16].blocks[x][y%16][z] = Block{id: uint(BlockID)}
		}
	}

	return nil
}

func handleBlockChangePacket(g *Game, r *bytes.Reader) error {
	if !g.settings.ReciveMap {
		return nil
	}
	x, y, z, err := pk.UnpackPosition(r)
	if err != nil {
		return err
	}
	c := g.wd.chunks[chunkLoc{x >> 4, z >> 4}]
	if c != nil {
		id, err := pk.UnpackVarInt(r)
		if err != nil {
			return err
		}
		c.sections[y/16].blocks[x&15][y&15][z&15] = Block{id: uint(id)}
	}

	return nil
}

func handleChatMessagePacket(g *Game, r *bytes.Reader) error {

	s, err := pk.UnpackString(r)
	if err != nil {
		return err
	}
	pos, err := r.ReadByte()
	if err != nil {
		return err
	}
	cm, err := newChatMsg([]byte(s))
	if err != nil {
		return err
	}
	g.events <- ChatMessageEvent{cm, pos}

	return nil
}

func handleUpdateHealthPacket(g *Game, r *bytes.Reader) (err error) {
	g.player.Health, err = pk.UnpackFloat(r)
	if err != nil {
		return
	}
	g.player.Food, err = pk.UnpackVarInt(r)
	if err != nil {
		return
	}
	g.player.FoodSaturation, err = pk.UnpackFloat(r)
	if err != nil {
		return
	}

	if g.player.Health < 1 { //player is dead
		g.events <- PlayerDeadEvent{} //Dead event
		sendPlayerPositionAndLookPacket(g)
		time.Sleep(time.Second * 2)  //wait for 2 sec make it more like a human
		sendClientStatusPacket(g, 0) //status 0 means perform respawn
	}
	return
}

func handleJoinGamePacket(g *Game, r *bytes.Reader) error {
	eid, err := pk.UnpackInt32(r)
	if err != nil {
		return fmt.Errorf("read EntityID fail: %v", err)
	}
	g.Info.EntityID = int(eid)
	gamemode, err := r.ReadByte()
	if err != nil {
		return fmt.Errorf("read gamemode fail: %v", err)
	}
	g.Info.Gamemode = int(gamemode & 0x7)
	g.Info.Hardcore = gamemode&0x8 != 0
	dimension, err := pk.UnpackInt32(r)
	if err != nil {
		return fmt.Errorf("read dimension fail: %v", err)
	}
	g.Info.Dimension = int(dimension)
	difficulty, err := r.ReadByte()
	if err != nil {
		return fmt.Errorf("read difficulty fail: %v", err)
	}
	g.Info.Difficulty = int(difficulty)
	// ignore Max Players
	_, err = r.ReadByte()
	if err != nil {
		return fmt.Errorf("read MaxPlayers fail: %v", err)
	}

	g.Info.LevelType, err = pk.UnpackString(r)
	if err != nil {
		return fmt.Errorf("read LevelType fail: %v", err)
	}
	rdi, err := r.ReadByte()
	if err != nil {
		return fmt.Errorf("read ReducedDebugInfo fail: %v", err)
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
	if !g.settings.ReciveMap {
		return nil
	}

	c, x, y, err := unpackChunkDataPacket(p, g.Info.Dimension == 0)
	g.wd.chunks[chunkLoc{x, y}] = c
	return err
}

var isSpawn bool

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

	//handle PlayerSpawnEvent
	if !isSpawn {
		g.events <- PlayerSpawnEvent{}
		isSpawn = true
	}
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
	E := g.wd.Entities[ID]
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

func handleEntityHeadLookPacket(g *Game, r *bytes.Reader) {

}

func handleEntityRelativeMovePacket(g *Game, r *bytes.Reader) error {
	ID, err := pk.UnpackVarInt(r)
	if err != nil {
		return err
	}
	E := g.wd.Entities[ID]
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

func handleSpawnPlayerPacket(g *Game, r *bytes.Reader) (err error) {
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

	g.wd.Entities[np.entityID] = np //把该玩家添加到全局实体表里面
	return nil
}

func handleWindowItemsPacket(g *Game, r *bytes.Reader) (err error) {
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
		g.events <- InventoryChangeEvent(-2)
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

func sendPlayerPositionPacket(g *Game) {
	var data []byte
	data = append(data, pk.PackDouble(g.player.X)...)
	data = append(data, pk.PackDouble(g.player.Y)...)
	data = append(data, pk.PackDouble(g.player.Z)...)
	data = append(data, pk.PackBoolean(g.player.OnGround))

	g.sendChan <- pk.Packet{
		ID:   0x10,
		Data: data,
	}
}

func sendKeepAlivePacket(g *Game, KeepAliveID int64) {
	g.sendChan <- pk.Packet{
		ID:   0x0E,
		Data: pk.PackUint64(uint64(KeepAliveID)),
	}
}

func sendClientStatusPacket(g *Game, status int32) {
	data := pk.PackVarInt(status)
	g.sendChan <- pk.Packet{
		ID:   0x03,
		Data: data,
	}
}

//hand could be 0: main hand, 1: off hand
func sendAnimationPacket(g *Game, hand int32) {
	data := pk.PackVarInt(hand)
	g.sendChan <- pk.Packet{
		ID:   0x27,
		Data: data,
	}
}

func sendPlayerDiggingPacket(g *Game, status int32, x, y, z int, face Face) {
	data := pk.PackVarInt(status)
	data = append(data, pk.PackPosition(x, y, z)...)
	data = append(data, byte(face))

	g.sendChan <- pk.Packet{
		ID:   0x18,
		Data: data,
	}
}

func sendUseItemPacket(g *Game, hand int32) {
	data := pk.PackVarInt(hand)
	g.sendChan <- pk.Packet{
		ID:   0x2A,
		Data: data,
	}
}
