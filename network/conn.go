package network

import (
	"bufio"
	"io"
	"net"

	pk "github.com/Tnze/gomcbot/network/packet"
)

type Conn struct {
	socket net.Conn
	io.ByteReader
	io.Writer

	threshold int
}

func DialMC(addr string) (conn *Conn, err error) {
	conn = new(Conn)
	conn.socket, err = net.Dial("tcp", addr)
	if err != nil {
		return
	}

	conn.ByteReader = bufio.NewReader(conn.socket)
	conn.Writer = conn.socket

	return
}

func (c *Conn) ReadPacket() (pk.Packet, error) {
	pk, err := pk.RecvPacket(c.ByteReader, c.threshold > 0)
	return *pk, err
}

func (c *Conn) WritePacket(p pk.Packet) error {
	_, err := c.Write(p.Pack(c.threshold))
	return err
}
