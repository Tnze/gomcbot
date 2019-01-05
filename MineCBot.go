package gomcbot

import (
	// "./authenticate"
	"bytes"
	"compress/zlib"
	"crypto/aes"
	// "crypto/cipher"
	"fmt"
	"io"
	"io/ioutil"
	"net"
)

// Game is the Object used to access Minecraft server
type Game struct {
	addr       string
	port       int
	conn       net.Conn
	recvPacket func() (*packet, error)
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
		err = fmt.Errorf("cannot connect the server %q: %v", addr, err)
	}

	hsPacket := newHandshakePacket(404, addr, port, 1) //构造握手包
	_, err = conn.Write(hsPacket.pack())               //握手
	if err != nil {
		return "", fmt.Errorf("send handshake packect fail: %v", err)
	}
	reqPacket := packet{ //构造请求服务器状态包
		ID:   0,
		Data: []byte{},
	}
	_, err = conn.Write(reqPacket.pack()) //请求
	if err != nil {
		return "", fmt.Errorf("send list packect fail: %v", err)
	}
	data, err := recvPacket(conn)
	if err != nil {
		return "", fmt.Errorf("recv list packect fail: %v", err)
	}
	// fmt.Println(resp)
	s, _ := unpackString(data[1:])
	return string(s), nil
}

// JoinServer connect a server from addr
func (p *Player) JoinServer(addr string, port int) (g *Game, err error) {
	g = new(Game)
	g.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		err = fmt.Errorf("cannot connect the server %q: %v", addr, err)
		return
	}

	g.recvPacket = func() (*packet, error) { //默认接收未压缩的数据
		data, err := recvPacket(g.conn)
		if err != nil {
			return nil, err
		}
		return &packet{
			ID:   data[0],
			Data: data[1:],
		}, nil
	}

	// State Handshake

	//握手
	hsPacket := newHandshakePacket(404, addr, port, 2) //构造握手包
	_, err = g.conn.Write(hsPacket.pack())             //握手
	if err != nil {
		err = fmt.Errorf("send handshake packect fail: %v", err)
		return
	}

	// State Login
	lsPacket := newLoginStartPacket(p.Name)
	_, err = g.conn.Write(lsPacket.pack()) //LoginStart
	if err != nil {
		err = fmt.Errorf("send login start packect fail: %v", err)
		return
	}
	for {
		var pack *packet
		pack, err = g.recvPacket()
		if err != nil {
			err = fmt.Errorf("recv packet at state Login fail: %v", err)
			return
		}
		switch pack.ID {
		case 0x00: //Disconnect
			s, _ := unpackString(pack.Data)
			err = fmt.Errorf("connect disconnected by server because: %s", s)
		case 0x01: //Encryption Request
			er := unpackEncryptionRequest(*pack) //从服务器接收EncryptionRequest
			err = loginAuth(p.AsTk, p.UUID, er)  //向Mojang验证
			if err != nil {
				err = fmt.Errorf("login auth fail: %v", err)
				return
			}
			key, encodeStream := newSymmetricEncryption()
			var erp *packet
			erp, err = genEncryptionKeyResponse(key, er.PublicKey, er.VerifyToken)
			if err != nil {
				err = fmt.Errorf("gen encryption key response fail: %v", err)
				return
			}
			g.conn.Write(erp.pack()) // Encryption Key Response

			data := make([]byte, aes.BlockSize)
			_, err = io.ReadAtLeast(g.conn, data, aes.BlockSize)
			if err != nil {
				err = fmt.Errorf("read fail: %v", err)
				return
			}
			decodeData := make([]byte, aes.BlockSize)
			encodeStream.XORKeyStream(decodeData, data)

			fmt.Println(decodeData)
		case 0x02: //Login Success
			uuid, l := unpackString(pack.Data)
			fmt.Println(uuid)
			UserName, _ := unpackString(pack.Data[l:])
			fmt.Println(UserName)
			fmt.Println("Login success")
			return //switches the connection state to PLAY.
		case 0x03: //Set Compression
			g.recvPacket = func() (*packet, error) { //启用压缩
				data, err := recvPacket(g.conn)
				if err != nil {
					return nil, err
				}

				DataLength, m := unpackVarInt(data)
				if DataLength != 0 {
					buff := bytes.NewReader(data[m:])
					r, err := zlib.NewReader(buff)
					if err != nil {
						return nil, fmt.Errorf("new zlib error: %v", err)
					}
					d, err := ioutil.ReadAll(r)
					if err != nil {
						return nil, fmt.Errorf("read zlib error: %v", err)
					}
					r.Close()
					return &packet{
						ID:   d[0],
						Data: d[1:],
					}, nil
				}
				return &packet{
					ID:   data[m],
					Data: data[m+1:],
				}, err
			}
		case 0x04: //Login Plugin Request
			fmt.Println("Waring Login Plugin Request")
		}
	}
}
