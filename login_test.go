package gomcbot

import (
	// "github.com/JoshuaDoes/go-yggdrasil"
	"testing"
)

func TestUnpackER(t *testing.T) {
	p := packet{Data: []byte{0, 162, 1, 48, 129, 159, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 1, 5, 0, 3, 129, 141, 0, 48, 129, 137, 2, 129, 129, 0, 133, 14, 220, 58, 255, 89, 147, 202, 197, 64, 239, 139, 25, 117, 190, 134, 140, 252, 199, 227, 220, 210, 74, 142, 223, 86, 114, 201, 44, 125, 51, 250, 65, 158, 72, 209, 175, 27, 132, 193, 174, 86, 250, 171, 19, 142, 81, 111, 152, 83, 130, 105, 230, 59, 67, 241, 118, 62, 203, 11, 195, 15, 127, 158, 42, 9, 35, 215, 240, 163, 163, 118, 44, 15, 48, 121, 224, 75, 195, 146, 24, 143, 76, 87, 235, 11, 147, 164, 219, 59, 220, 229, 31, 166, 203, 12, 30, 227, 70, 115, 38, 3, 17, 32, 118, 207, 167, 250, 29, 61, 242, 208, 83, 252, 29, 56, 6, 254, 246, 177, 39, 138, 132, 91, 100, 71, 116, 33, 2, 3, 1, 0, 1, 4, 157, 172, 11, 128}}
	t.Log(p.Data)
	i := 0
	serverID, length := unpackString(p.Data[i:])
	t.Log(serverID, length)
	i += length
	t.Log(i, p.Data[i:])
	publicKeyLength, length := unpackVarInt(p.Data[i:])
	t.Log(publicKeyLength, length)
	i += length
	publicKey := p.Data[i : i+int(publicKeyLength)]
	i += int(publicKeyLength)
	t.Log(publicKey)
	t.Log(i, p.Data[i:])
	verifyTokenLength, length := unpackVarInt(p.Data[i:])
	t.Log(length, verifyTokenLength)
	i += length
	verifyToken := p.Data[i : i+int(verifyTokenLength)]
	t.Log(verifyToken)
}
