package gomcbot

import "testing"

// 0	0x00	0
// 1	0x01	1
// 2	0x02	2
// 127	0x7f	127
// 128	0x80 0x01	128 1
// 255	0xff 0x01	255 1
// 2147483647	0xff 0xff 0xff 0xff 0x07	255 255 255 255 7
// -1	0xff 0xff 0xff 0xff 0x0f	255 255 255 255 15
// -2147483648	0x80 0x80 0x80 0x80 0x08	128 128 128 128 8

var Int = []int32{0, 1, 2, 127, 128, 255, 2147483647, -1, -2147483648}

var VarInt = [][]byte{
	[]byte{0x00},
	[]byte{0x01},
	[]byte{0x02},
	[]byte{0x7f},
	[]byte{0x80, 0x01},
	[]byte{0xff, 0x01},
	[]byte{0xff, 0xff, 0xff, 0xff, 0x07},
	[]byte{0xff, 0xff, 0xff, 0xff, 0x0f},
	[]byte{0x80, 0x80, 0x80, 0x80, 0x08},
}

func TestPackInt(t *testing.T) {
	for i, v := range Int {
		p := packVarInt(v)
		if !bytesEqual(p, VarInt[i]) {
			t.Errorf("pack int %d should be \"% x\" but not \"% x\"", v, VarInt[i], p)
		}
	}
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestUnpackInt(t *testing.T) {
	for i, v := range VarInt {
		p, _ := unpackVarInt(v)
		if Int[i] != p {
			t.Errorf("unpack \"% x\" should be %d but not %d", v, Int[i], p)
		}
	}
}
