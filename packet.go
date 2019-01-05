package gomcbot

import (
	"fmt"
	"io"
	// "io/ioutil"
	"net"
)

// Packet define a net data package
type packet struct {
	ID   byte
	Data []byte
}

func (p *packet) pack() (pack []byte) {
	pack = append(pack, packVarInt(int32(len(p.Data)+1))...) //len
	pack = append(pack, p.ID)                                //data
	pack = append(pack, p.Data...)                           //data
	return
}

func packString(s string) (ps []byte) {
	byteString := []byte(s)
	ps = append(ps, packVarInt(int32(len(byteString)))...) //len
	ps = append(ps, byteString...)                         //data
	return
}

func packUint16(n uint16) []byte {
	return []byte{
		byte(n >> 8),
		byte(n & 0xFF),
	}
}

func packVarInt(n int32) (VarInt []byte) {
	num := uint32(n)
	for {
		b := num & 0x7F
		num >>= 7
		if num != 0 {
			b |= 0x80
		}
		VarInt = append(VarInt, byte(b))
		if num == 0 {
			break
		}
	}
	return
}

func unpackString(b []byte) (s string, len int) {
	l, pre := unpackVarInt(b)
	len = int(l)
	// fmt.Println(b)
	// fmt.Println(len, pre)
	return string(b[pre : pre+len]), len + pre
}

func unpackVarInt(b []byte) (num int32, len int) {
	var n uint
	for i := 0; i < 5; i++ { //读数据前的长度标记
		n |= (uint(b[i]&0x7F) << uint(7*i))
		len++
		if b[i]&0x80 == 0 {
			break
		}
	}
	num = int32(n) //这里要把超过int32的负数溢出
	return
}

// recvPacket recive a packet from server
func recvPacket(conn net.Conn) ([]byte, error) {
	data := make([]byte, 1) //len=1
	len := 0
	for i := 0; i < 5; i++ { //读数据前的长度标记
		if _, err := io.ReadFull(conn, data); err != nil {
			return nil, fmt.Errorf("read len of packet fail: %v", err)
		}
		len |= (int(data[0]&0x7F) << uint(7*i))
		if data[0]&0x80 == 0 {
			break
		}
	}
	data = make([]byte, len)
	_, err := io.ReadAtLeast(conn, data, len) //读取数据
	if err != nil {
		return nil, fmt.Errorf("read packet data fail: %v", err)
	}

	return data, nil
}
