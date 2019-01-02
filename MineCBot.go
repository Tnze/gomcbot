package minecbot

import (
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
func (c *Connection) handshake(nextState int) error {
	var hsData []byte             //Handshake packet data
	hsData = append(hsData, 0x00) //Protocol Version
	hsData = append(hsData, packString(c.addr)...)
	hsData = append(hsData, packPort(c.port)...)
	hsData = append(hsData, 1)
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

	err = conn.handshake(1) //握手
	if err != nil {
		return "", fmt.Errorf("handshake fail: %v", err)
	}

	conn.sendPacket(Packet{
		Data: []byte{},
	})
	if err != nil {
		return "", fmt.Errorf("send list packect fail: %v", err)
	}
	resp, err := conn.recvPacket()
	if err != nil {
		return "", fmt.Errorf("send list packect fail: %v", err)
	}
	return string(unpackString(resp.Data)), nil
}
