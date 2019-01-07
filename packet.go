package gomcbot

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
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
func recvPacket(s *bufio.Reader, useZlib bool) (*packet, error) {
	var len int
	for i := 0; i < 5; i++ { //读数据前的长度标记
		b, err := s.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("read len of packet fail: %v", err)
		}
		len |= (int(b&0x7F) << uint(7*i))
		if b&0x80 == 0 {
			break
		}
	}

	data := make([]byte, len) //读包内容
	var err error
	for i := 0; i < len; i++ {
		data[i], err = s.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("read content of packet fail: %v", err)
		}
	}

	//解压数据
	if useZlib {
		return unCompress(data)
	}

	return &packet{
		ID:   data[0],
		Data: data[1:],
	}, nil
}

//读取一个压缩的包
func unCompress(data []byte) (*packet, error) {
	sizeUncompressed, len := unpackVarInt(data)
	uncompressData := make([]byte, sizeUncompressed)
	if sizeUncompressed != 0 { // != 0 means compressed, let's decompress
		b := bytes.NewReader(data)

		r, err := zlib.NewReader(b)
		defer r.Close()
		if err != nil {
			return nil, fmt.Errorf("decompress fail: %v", err)
		}
		_, err = io.ReadFull(r, uncompressData)
		if err != nil {
			return nil, fmt.Errorf("decompress fail: %v", err)
		}
	} else {
		uncompressData = data[len:]
	}
	return &packet{
		ID:   uncompressData[0],
		Data: uncompressData[1:],
	}, nil
}
