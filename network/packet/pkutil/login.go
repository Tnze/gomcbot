package pkutil

import (
	"bytes"

	pk "github.com/Tnze/gomcbot/network/packet"
)

// NewLoginStartPacket 构造一个LoginStart包
func NewLoginStartPacket(userName string) pk.Packet {
	return pk.Marshal(0, pk.String(userName))
}

type EncryptionRequest struct {
	ServerID    string
	PublicKey   []byte
	VerifyToken []byte
}

func UnpackEncryptionRequest(p pk.Packet) (*EncryptionRequest, error) {
	r := bytes.NewReader(p.Data)
	var serverID pk.String
	if err := serverID.Decode(r); err != nil {
		return nil, err
	}

	var publicKeyLength, verifyTokenLength pk.VarInt

	if err := publicKeyLength.Decode(r); err != nil {
		return nil, err
	}
	publicKey, err := pk.ReadNBytes(r, int(publicKeyLength))
	if err != nil {
		return nil, err
	}

	if err := verifyTokenLength.Decode(r); err != nil {
		return nil, err
	}
	verifyToken, err := pk.ReadNBytes(r, int(verifyTokenLength))
	if err != nil {
		return nil, err
	}

	er := EncryptionRequest{
		ServerID:    string(serverID),
		PublicKey:   publicKey,
		VerifyToken: verifyToken,
	}
	return &er, nil
}
