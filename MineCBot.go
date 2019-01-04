package gomcbot

import (
	"./authenticate"
	"fmt"
	"net"
)

// Connection is the Object used to access Minecraft server
type Connection struct {
	addr string
	port int
	conn net.Conn
}

// NewConnection connect a server from addr
func NewConnection(addr string, port int) (Connection, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		err = fmt.Errorf("cannot connect the server %q: %v", addr, err)
	}

	return Connection{
		addr: addr,
		port: port,
		conn: conn,
	}, err
}

// handshake make handshake with server
func (c *Connection) handshake(protocolVersion int, nextState byte) error {
	var hsData []byte                                       //Handshake packet data
	hsData = append(hsData, packVarInt(protocolVersion)...) //Protocol Version
	hsData = append(hsData, packString(c.addr)...)
	hsData = append(hsData, packUint16(c.port)...)
	hsData = append(hsData, nextState)
	err := c.sendPacket(Packet{
		ID:   0,
		Data: hsData,
	})
	if err != nil {
		return fmt.Errorf("send handshake packet fail: %v", err)
	}
	return nil
}

// PingAndList chack server status and list online player
func PingAndList(addr string, port int) (string, error) {
	conn, err := NewConnection(addr, 25565) //连接
	if err != nil {
		return "", fmt.Errorf("connect server fail: %v", err)
	}

	err = conn.handshake(404, 1) //握手
	if err != nil {
		return "", fmt.Errorf("handshake fail: %v", err)
	}

	conn.sendPacket(Packet{ID: 0x00, Data: []byte{}}) // 请求状态
	if err != nil {
		return "", fmt.Errorf("send list packect fail: %v", err)
	}
	resp, err := conn.recvPacket()
	if err != nil {
		return "", fmt.Errorf("recv list packect fail: %v", err)
	}
	// fmt.Println(resp)
	s, _ := unpackString(resp.Data)
	return string(s), nil
}

// Login make server auth you
func (c *Connection) Login(auth authenticate.Response) error {
	// 握手
	err := c.handshake(404, 2)
	if err != nil {
		return fmt.Errorf("handshake fail: %v", err)
	}

	// Login Start
	c.sendPacket(Packet{ID: 0x00, Data: packString(auth.SelectedProfile.Name)})
	if err != nil {
		return fmt.Errorf("send list packect fail: %v", err)
	}

	// Encryption Request
	er, err := c.recvPacket()
	if err != nil {
		return fmt.Errorf("recv encryption request fail: %v", err)
	}
	unpackER := unpackEncryptionRequest(er) //解码Encryption Request包
	// Client Auth

	err = loginAuth(auth, unpackER)
	if err != nil {
		return fmt.Errorf("login with auth fail: %v", err)
	}

	// Encryption Key Response
	sharedSecret, _, _ := newSymmetricEncryption()
	ekr, err := genEncryptionKeyResponse(sharedSecret, unpackER.PublicKey, unpackER.VerifyToken)
	if err != nil {
		return fmt.Errorf("gen encryption key response fail: %v", err)
	}
	c.sendPacket(*ekr)

	// Login success
	ls, err := c.recvPacket()
	if err != nil {
		return fmt.Errorf("recv login success fail: %v", err)
	}
	fmt.Println("Ls: ", ls)
	uuid, username := unpackLoginSuccess(ls)
	fmt.Println(uuid, username)
	return nil
}
