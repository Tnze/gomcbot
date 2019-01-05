package gomcbot

// newHandshakePacket 构造了一个Handshake包
func newHandshakePacket(protocolVersion int, addr string, port int, nextState byte) *packet {
	var hsData []byte                                              //Handshake packet data
	hsData = append(hsData, packVarInt(int32(protocolVersion))...) //Protocol Version
	hsData = append(hsData, packString(addr)...)
	hsData = append(hsData, packUint16(uint16(port))...)
	hsData = append(hsData, nextState)
	pack := packet{
		ID:   0,
		Data: hsData,
	}
	return &pack
}

// newLoginStartPakcket 构造一个LoginStart包
func newLoginStartPacket(userName string) *packet {
	var data []byte
	data = append(data, packString(userName)...)
	pack := packet{
		ID:   0,
		Data: data,
	}
	return &pack
}
