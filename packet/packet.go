package packet

import (
	"bufio"
	"bytes"
	"compress/zlib"
	// "encoding/binary"
	"fmt"
	"io"
	"unsafe"
)

// Packet define a net data package
type Packet struct {
	ID   byte
	Data []byte
}

// Pack 打包一个数据包
func (p *Packet) Pack(threshold int) (pack []byte) {
	data := []byte{p.ID}           //data
	data = append(data, p.Data...) //data

	if threshold > 0 { //是否启用了压缩
		if len(data) > threshold { //是否需要压缩
			Len := len(data)
			VarLen := PackVarInt(int32(Len))
			data = Compress(data)

			pack = append(pack, PackVarInt(int32(len(VarLen)+len(data)))...)
			pack = append(pack, VarLen...)
			pack = append(pack, data...)
		} else {
			pack = append(pack, PackVarInt(int32(len(data)+1))...)
			pack = append(pack, 0x00)
			pack = append(pack, data...)
		}
	} else {
		pack = append(pack, PackVarInt(int32(len(data)))...) //len
		pack = append(pack, data...)
	}

	return
}

// PackString 打包一个字符串
func PackString(s string) (p []byte) {
	byteString := []byte(s)
	p = append(p, PackVarInt(int32(len(byteString)))...) //len
	p = append(p, byteString...)                         //data
	return
}

// PackUint16 打包一个16位无符号整数
func PackUint16(n uint16) (p []byte) {
	p = []byte{
		byte(n >> 8),
		byte(n),
	}
	return
}

// PackUint32 打包一个32位无符号整数
func PackUint32(n uint32) (p []byte) {
	p = []byte{
		byte(n >> 24), byte(n >> 16),
		byte(n >> 8), byte(n),
	}
	return
}

// PackUint64 打包一个64位无符号整数
func PackUint64(n uint64) (p []byte) {
	p = []byte{
		byte(n >> 56), byte(n >> 48), byte(n >> 40), byte(n >> 32),
		byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n),
	}
	return
}

//PackVarInt 打包一个VarInt
func PackVarInt(n int32) (VarInt []byte) {
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

//PackPosition 打包一个位置
func PackPosition(x, y, z int) (p []byte) {
	p = make([]byte, 8)
	position := (int64(x&0x3FFFFFF) << 38) | int64((y&0xFFF)<<26) | int64(z&0x3FFFFFF)
	for i := 7; i >= 0; i-- {
		p[i] = byte(position)
		position >>= 8
	}
	return
}

//PackFloat 打包一个32位浮点数
func PackFloat(f float32) (p []byte) {
	n := *(*uint32)(unsafe.Pointer(&f))
	return PackUint32(n)
}

//PackDouble 打包一个64位浮点数
func PackDouble(d float64) (p []byte) {
	n := *(*uint64)(unsafe.Pointer(&d))
	return PackUint64(n)
}

//PackBoolean 打包一个布尔值
func PackBoolean(b bool) byte {
	if b {
		return 0x01
	}
	return 0x00

}

//UnpackString 读取一个字符串
func UnpackString(b []byte) (s string, len int) {
	l, pre := UnpackVarInt(b)
	len = int(l)
	return string(b[pre : pre+len]), len + pre
}

//UnpackVarInt 读取一个VarInt
func UnpackVarInt(b []byte) (num int32, len int) {
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

//UnpackInt16 读取一个16位有符号整数
func UnpackInt16(b []byte) (num int16) {
	return int16(b[0])<<8 | int16(b[1])
}

//UnpackInt32 读取一个32位有符号整数
func UnpackInt32(b []byte) (num int32) {
	return int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3])
}

//UnpackInt64 读取一个64位有符号整数
func UnpackInt64(b []byte) (num int64) {
	return int64(b[0])<<56 | int64(b[1])<<48 | int64(b[2])<<40 | int64(b[3])<<32 |
		int64(b[4])<<24 | int64(b[5])<<16 | int64(b[6])<<8 | int64(b[7])
}

// UnpackPosition 读取一个位置
func UnpackPosition(b []byte) (x, y, z int) {
	position := UnpackInt64(b)
	x = int(position >> 38)
	y = int((position >> 26) & 0xFFF)
	z = int(position << 38 >> 38)

	//处理负数
	if x >= 1<<25 {
		x -= 1 << 26
	}
	if y >= 1<<11 {
		y -= 1 << 12
	}
	if z >= 1<<25 {
		z -= 1 << 26
	}
	return
}

// UnpackFloat 读取一个单精度浮点数
func UnpackFloat(b []byte) (num float32) {
	n := UnpackInt32(b)
	return *(*float32)(unsafe.Pointer(&n))
}

// UnpackDouble 读取一个双精度浮点数
func UnpackDouble(b []byte) (num float64) {
	n := UnpackInt64(b)
	return *(*float64)(unsafe.Pointer(&n))
}

// RecvPacket recive a packet from server
func RecvPacket(s *bufio.Reader, useZlib bool) (*Packet, error) {
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
		return UnCompress(data)
	}
	return &Packet{
		ID:   data[0],
		Data: data[1:],
	}, nil
}

// UnCompress 读取一个压缩的包
func UnCompress(data []byte) (*Packet, error) {
	sizeUncompressed, len := UnpackVarInt(data)
	uncompressData := make([]byte, sizeUncompressed)
	if sizeUncompressed != 0 { // != 0 means compressed, let's decompress
		b := bytes.NewReader(data[len:])

		r, err := zlib.NewReader(b)

		if err != nil {
			return nil, fmt.Errorf("decompress fail: %v", err)
		}
		_, err = io.ReadFull(r, uncompressData)
		if err != nil {
			return nil, fmt.Errorf("decompress fail: %v", err)
		}
		r.Close()
	} else {
		uncompressData = data[len:]
	}
	return &Packet{
		ID:   uncompressData[0],
		Data: uncompressData[1:],
	}, nil
}

// Compress 压缩数据
func Compress(data []byte) []byte {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(data)
	w.Close()
	return b.Bytes()
}
