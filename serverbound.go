package gomcbot

import (
	pk "github.com/Tnze/gomcbot/packet"
)

// newHandshakePacket 构造了一个Handshake包
func newHandshakePacket(protocolVersion int, addr string, port int, nextState byte) *pk.Packet {
	var hsData []byte                                                 //Handshake packet data
	hsData = append(hsData, pk.PackVarInt(int32(protocolVersion))...) //Protocol Version
	hsData = append(hsData, pk.PackString(addr)...)
	hsData = append(hsData, pk.PackUint16(uint16(port))...)
	hsData = append(hsData, nextState)
	pack := pk.Packet{
		ID:   0,
		Data: hsData,
	}
	return &pack
}

// newLoginStartPakcket 构造一个LoginStart包
func newLoginStartPacket(userName string) *pk.Packet {
	var data []byte
	data = append(data, pk.PackString(userName)...)
	pack := pk.Packet{
		ID:   0,
		Data: data,
	}
	return &pack
}
