package gomcbot

import (
	pk "./packet"
	"bufio"
	"crypto/cipher"
	"fmt"
	"io"
	"net"
)

// Game is the Object used to access Minecraft server
type Game struct {
	addr      string
	port      int
	conn      net.Conn
	reciver   *bufio.Reader
	sender    io.Writer
	threshold int
	Info      PlayerInfo
	abilities PlayerAbilities
	settings  Settings
	player    Player
	sendChan  chan pk.Packet //be used when HandleGame
	world     World          //The Level
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
	s, _ := pk.UnpackString(recv.Data)
	return string(s), nil
}

// JoinServer connect a server from addr
func (p *Auth) JoinServer(addr string, port int) (g *Game, err error) {
	//连接
	g = new(Game)
	g.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		err = fmt.Errorf("cannot connect the server %q: %v", addr, err)
		return
	}
	g.settings = DefaultSettings //默认设置
	g.reciver = bufio.NewReader(g.conn)
	g.sender = g.conn
	g.world.Entities = make(map[int32]Entity)

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
			s, _ := pk.UnpackString(pack.Data)
			err = fmt.Errorf("connect disconnected by server because: %s", s)
			return
		case 0x01: //Encryption Request
			key, encoStream, decoStream := newSymmetricEncryption() //创建AES对称加密密钥
			//This still not work
			er := unpackEncryptionRequest(*pack) //从服务器接收EncryptionRequest
			fmt.Println(er)
			err = loginAuth(p.AsTk, p.Name, p.UUID, key, er) //向Mojang验证
			if err != nil {
				err = fmt.Errorf("login fail: %v", err)
				return
			}

			var p *pk.Packet
			p, err = genEncryptionKeyResponse(key, er.PublicKey, er.VerifyToken)
			if err != nil {
				err = fmt.Errorf("gen encryption key response fail: %v", err)
				return
			}
			err = g.sendPacket(p) // Encryption Key Response
			if err != nil {
				return
			}

			g.reciver = bufio.NewReader(cipher.StreamReader{ //Set reciver for AES
				S: decoStream,
				R: g.conn,
			})
			g.sender = cipher.StreamWriter{
				S: encoStream,
				W: g.conn,
			}
		case 0x02: //Login Success
			// uuid, l := unpackString(pack.Data)
			// UserName, _ := unpackString(pack.Data[l:])
			fmt.Println("Login success")
			return //switches the connection state to PLAY.
		case 0x03: //Set Compression
			threshold, _ := pk.UnpackVarInt(pack.Data)
			g.threshold = int(threshold)
		case 0x04: //Login Plugin Request
			fmt.Println("Waring Login Plugin Request")
		}
	}
}
