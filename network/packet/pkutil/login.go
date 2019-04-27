package pkutil

import (
	"bytes"

	pk "github.com/Tnze/gomcbot/network/packet"
)

// NewHandshakePacket 构造了一个Handshake包
func NewHandshakePacket(protocolVersion int, addr string, port int, nextState byte) *pk.Packet {
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

// NewLoginStartPacket 构造一个LoginStart包
func NewLoginStartPacket(userName string) *pk.Packet {
	var data []byte
	data = append(data, pk.PackString(userName)...)
	pack := pk.Packet{
		ID:   0,
		Data: data,
	}
	return &pack
}

type encryptionRequest struct {
	ServerID    string
	PublicKey   []byte
	VerifyToken []byte
}

func unpackEncryptionRequest(p pk.Packet) (*encryptionRequest, error) {
	r := bytes.NewReader(p.Data)
	serverID, err := pk.UnpackString(r)
	if err != nil {
		return nil, err
	}
	publicKeyLength, err := pk.UnpackVarInt(r)
	if err != nil {
		return nil, err
	}
	publicKey, err := pk.ReadNBytes(r, int(publicKeyLength))
	if err != nil {
		return nil, err
	}
	verifyTokenLength, err := pk.UnpackVarInt(r)
	if err != nil {
		return nil, err
	}
	verifyToken, err := pk.ReadNBytes(r, int(verifyTokenLength))
	if err != nil {
		return nil, err
	}

	er := encryptionRequest{
		ServerID:    serverID,
		PublicKey:   publicKey,
		VerifyToken: verifyToken,
	}
	return &er, nil
}
