package gomcbot

import "testing"

func TestPacket(t *testing.T) {
	p := Packet{
		Data: []byte{},
	}
	t.Log(p.pack())
}

func TestPacketHandshokePacket(t *testing.T) {
	var data []byte
	data = append(data, packString("localhost")...)
	data = append(data, packUint16(25565)...)
	data = append(data, 1)
	p := Packet{
		ID:   0,
		Data: data,
	}
	t.Log(p.pack()) //Should be [15 0 0 9 108 111 99 97 108 104 111 115 116 99 221 1]
}

func TestUnpackInt(t *testing.T) {
	if res, _ := unpackVarInt([]byte{0x81, 0x04}); res != 0x0201 {
		t.Errorf("0x81 0x04 should be 0x84 but not %x", res)
	}
}
