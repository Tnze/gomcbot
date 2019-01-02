package minecbot

import (
	"fmt"
	"io"
)

// Packet define a net data package
type Packet struct {
	ID   byte
	Data []byte
}

func (p *Packet) pack() (pack []byte) {
	pack = append(pack, packInt(len(p.Data)+1)...) //len
	pack = append(pack, p.ID)                      //data
	pack = append(pack, p.Data...)                 //data
	return
}

func packString(s string) (ps []byte) {
	byteString := []byte(s)
	ps = append(ps, packInt(len(byteString))...) //len
	ps = append(ps, byteString...)               //data
	return
}

func unpackString(b []byte) string {
	len, pre := unpackInt(b)
	// fmt.Println(b)
	// fmt.Println(len, pre)
	return string(b[pre : len+1])
}

func packPort(n int) []byte {
	return []byte{
		byte(n >> 8),
		byte(n & 0xFF),
	}
}

func packInt(n int) (VarInt []byte) {
	for n != 0 {
		b := n & 0x7F
		n >>= 7
		if n > 0 {
			b |= 0x80
		}
		VarInt = append(VarInt, byte(b))
	}
	return
}

func unpackInt(b []byte) (n, len int) {
	for i := 0; i < 5; i++ { //读数据前的长度标记
		n |= (int(b[i]&0x7F) << uint(7*i))
		len++
		if b[i]&0x80 == 0 {
			break
		}
	}
	return
}

// sendPacket send a packet to server
func (c *Connection) sendPacket(p Packet) error {
	// fmt.Println(p.pack())
	_, err := c.conn.Write(p.pack())
	if err != nil {
		err = fmt.Errorf("cannot send packet to ")
	}
	return err
}

// recvPacket recive a packet from server
func (c *Connection) recvPacket() (Packet, error) {
	data := make([]byte, 1) //len=1
	len := 0
	for i := 0; i < 5; i++ { //读数据前的长度标记
		if _, err := io.ReadFull(c.conn, data); err != nil {
			return Packet{}, fmt.Errorf("read len of packet fail: %v", err)
		}
		len |= (int(data[0]&0x7F) << uint(7*i))
		if data[0]&0x80 == 0 {
			break
		}
	}
	data = make([]byte, len)
	_, err := io.ReadAtLeast(c.conn, data, len) //读取数据
	if err != nil {
		return Packet{}, fmt.Errorf("read packet data fail: %v", err)
	}

	return Packet{
		ID:   data[0],
		Data: data[1:],
	}, nil
}
