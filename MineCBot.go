package gomcbot

import (
	"bufio"
	"crypto/cipher"
	"fmt"
	"net"
)

// Game is the Object used to access Minecraft server
type Game struct {
	addr      string
	port      int
	conn      net.Conn
	reciver   *bufio.Reader
	threshold int32
}

func (g *Game) recvPacket() (*packet, error) {
	return recvPacket(g.reciver, g.threshold > 0)
}

// Player includes a account
type Player struct {
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
	_, err = conn.Write(hsPacket.pack())
	if err != nil {
		return "", fmt.Errorf("send handshake packect fail: %v", err)
	}

	//请求服务器状态
	reqPacket := packet{
		ID:   0,
		Data: []byte{},
	}
	_, err = conn.Write(reqPacket.pack())
	if err != nil {
		return "", fmt.Errorf("send list packect fail: %v", err)
	}

	//服务器返回状态
	recv, err := recvPacket(bufio.NewReader(conn), false)
	if err != nil {
		return "", fmt.Errorf("recv list packect fail: %v", err)
	}
	s, _ := unpackString(recv.Data)
	return string(s), nil
}

// JoinServer connect a server from addr
func (p *Player) JoinServer(addr string, port int) (g *Game, err error) {
	//连接
	g = new(Game)
	g.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		err = fmt.Errorf("cannot connect the server %q: %v", addr, err)
		return
	}
	g.reciver = bufio.NewReader(g.conn)

	//握手
	hsPacket := newHandshakePacket(404, addr, port, 2) //构造握手包
	_, err = g.conn.Write(hsPacket.pack())             //握手
	if err != nil {
		err = fmt.Errorf("send handshake packect fail: %v", err)
		return
	}

	//登录
	lsPacket := newLoginStartPacket(p.Name)
	_, err = g.conn.Write(lsPacket.pack()) //LoginStart
	if err != nil {
		err = fmt.Errorf("send login start packect fail: %v", err)
		return
	}
	for {
		//Recive Packet
		var pack *packet
		pack, err = g.recvPacket()
		if err != nil {
			err = fmt.Errorf("recv packet at state Login fail: %v", err)
			return
		}

		//Handle Packet
		switch pack.ID {
		case 0x00: //Disconnect
			s, _ := unpackString(pack.Data)
			err = fmt.Errorf("connect disconnected by server because: %s", s)
			return
		case 0x01: //Encryption Request
			//This still not work
			er := unpackEncryptionRequest(*pack) //从服务器接收EncryptionRequest
			err = loginAuth(p.AsTk, p.UUID, er)  //向Mojang验证
			if err != nil {
				err = fmt.Errorf("login auth fail: %v", err)
				return
			}

			key, _, decoStream := newSymmetricEncryption()
			var p *packet
			p, err = genEncryptionKeyResponse(key, er.PublicKey, er.VerifyToken)
			if err != nil {
				err = fmt.Errorf("gen encryption key response fail: %v", err)
				return
			}
			g.conn.Write(p.pack()) // Encryption Key Response

			g.reciver = bufio.NewReader(cipher.StreamReader{ //Set reciver for AES
				S: decoStream,
				R: g.conn,
			})
		case 0x02: //Login Success
			// uuid, l := unpackString(pack.Data)
			// UserName, _ := unpackString(pack.Data[l:])
			fmt.Println("Login success")
			return //switches the connection state to PLAY.
		case 0x03: //Set Compression
			g.threshold, _ = unpackVarInt(pack.Data)
		case 0x04: //Login Plugin Request
			fmt.Println("Waring Login Plugin Request")
		}
	}
}
