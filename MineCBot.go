package gomcbot

import (
	"bufio"
	"bytes"
	"crypto/cipher"
	"fmt"
	pk "github.com/Tnze/gomcbot/packet"
	"io"
	"net"
)

// Game is the Object used to access Minecraft server
type Game struct {
	addr string
	port int
	conn net.Conn

	reciver *bufio.Reader
	sender  io.Writer

	threshold int
	Info      PlayerInfo
	abilities PlayerAbilities
	settings  Settings
	player    Player
	wd        world //the map data

	sendChan chan pk.Packet  //be used when HandleGame
	recvChan chan *pk.Packet //be used when HandleGame
	events   chan Event
	motion   chan func() //used to submit a function and HandleGame do

	soundCallBack func(sound int32, category int32, x, y, z float64, volume, pitch float32)
}

// PingAndList chack server status and list online player
func PingAndList(addr string, port int) (string, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", addr, port)) //连接
	if err != nil {
		return "", fmt.Errorf("cannot connect the server %q: %v", addr, err)
	}

	//握手
	hsPacket := newHandshakePacket(404, addr, port, 1)
	_, err = conn.Write(hsPacket.Pack(-1))
	if err != nil {
		return "", fmt.Errorf("send handshake packect fail: %v", err)
	}

	//请求服务器状态
	reqPacket := pk.Packet{
		ID:   0,
		Data: []byte{},
	}
	_, err = conn.Write(reqPacket.Pack(-1))
	if err != nil {
		return "", fmt.Errorf("send list packect fail: %v", err)
	}

	//服务器返回状态
	recv, err := pk.RecvPacket(bufio.NewReader(conn), false)
	if err != nil {
		return "", fmt.Errorf("recv list packect fail: %v", err)
	}
	s, _ := pk.UnpackString(bytes.NewReader(recv.Data))
	return string(s), nil
}

// JoinServer connect a Minecraft server.
// Return a JSON string about server status.
// see JSON format at https://wiki.vg/Server_List_Ping#Response
func (p *Auth) JoinServer(addr string, port int) (g *Game, err error) {
	//连接
	g = new(Game)
	g.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		err = fmt.Errorf("cannot connect the server %q: %v", addr, err)
		return
	}

	//init Game
	g.settings = DefaultSettings //默认设置
	g.reciver = bufio.NewReader(g.conn)
	g.sender = g.conn
	g.wd.Entities = make(map[int32]Entity)
	g.wd.chunks = make(map[chunkLoc]*Chunk)
	g.events = make(chan Event)
	g.motion = make(chan func())

	//握手
	hsPacket := newHandshakePacket(404, addr, port, 2) //构造握手包
	err = g.sendPacket(hsPacket)                       //握手
	if err != nil {
		err = fmt.Errorf("send handshake packect fail: %v", err)
		return
	}

	//登录
	lsPacket := newLoginStartPacket(p.Name)
	err = g.sendPacket(lsPacket) //LoginStart
	if err != nil {
		err = fmt.Errorf("send login start packect fail: %v", err)
		return
	}
	for {
		//Recive Packet
		var pack *pk.Packet
		pack, err = g.recvPacket()
		if err != nil {
			err = fmt.Errorf("recv packet at state Login fail: %v", err)
			return
		}

		//Handle Packet
		switch pack.ID {
		case 0x00: //Disconnect
			s, _ := pk.UnpackString(bytes.NewReader(pack.Data))
			err = fmt.Errorf("connect disconnected by server because: %s", s)
			return
		case 0x01: //Encryption Request
			handleEncryptionRequest(g, pack, p)
		case 0x02: //Login Success
			// uuid, l := pk.UnpackString(pack.Data)
			// name, _ := unpackString(pack.Data[l:])
			return //switches the connection state to PLAY.
		case 0x03: //Set Compression
			threshold, _ := pk.UnpackVarInt(bytes.NewReader(pack.Data))
			g.threshold = int(threshold)
		case 0x04: //Login Plugin Request
			fmt.Println("Waring Login Plugin Request")
		}
	}
}

func (g *Game) recvPacket() (*pk.Packet, error) {
	return pk.RecvPacket(g.reciver, g.threshold > 0)
}

func (g *Game) sendPacket(p *pk.Packet) error {
	_, err := g.sender.Write(p.Pack(g.threshold))
	return err
}

// Auth includes a account
type Auth struct {
	Name string
	UUID string
	AsTk string
}

// 加密请求
func handleEncryptionRequest(g *Game, pack *pk.Packet, auth *Auth) error {
	//创建AES对称加密密钥
	key, encoStream, decoStream := newSymmetricEncryption()

	//解析EncryptionRequest包
	er, err := unpackEncryptionRequest(*pack)
	if err != nil {
		return err
	}
	err = loginAuth(auth.AsTk, auth.Name, auth.UUID, key, *er) //向Mojang验证
	if err != nil {
		return fmt.Errorf("login fail: %v", err)
	}

	// 响应加密请求
	var p *pk.Packet // Encryption Key Response
	p, err = genEncryptionKeyResponse(key, er.PublicKey, er.VerifyToken)
	if err != nil {
		return fmt.Errorf("gen encryption key response fail: %v", err)
	}
	err = g.sendPacket(p)
	if err != nil {
		return err
	}

	//加密连接
	g.reciver = bufio.NewReader(cipher.StreamReader{ //Set reciver for AES
		S: decoStream,
		R: g.conn,
	})
	g.sender = cipher.StreamWriter{
		S: encoStream,
		W: g.conn,
	}
	return nil
}

// Chat send chat message to server
// such as player send message in chat box
// msg can not longger than 256
func (g *Game) Chat(msg string) error {
	data := pk.PackString(msg)
	if len(data) > 256 {
		return fmt.Errorf("message too big")
	}

	pack := pk.Packet{
		ID:   0x02,
		Data: data,
	}

	g.sendChan <- pack
	return nil
}

//GetBlock return the block at (x, y, z)
func (g *Game) GetBlock(x, y, z int) Block {
	bc := make(chan Block)

	g.motion <- func() {
		bc <- g.wd.getBlock(x, y, z)
	}

	return <-bc
}

//GetPlayer return the player
func (g *Game) GetPlayer() Player {
	return g.player
}
