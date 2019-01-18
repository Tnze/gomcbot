package packet

import (
	"bytes"
	"testing"
)

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
		p := PackVarInt(v)
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
		p, _ := UnpackVarInt(bytes.NewReader(v))
		if Int[i] != p {
			t.Errorf("unpack \"% x\" should be %d but not %d", v, Int[i], p)
		}
	}
}

func TestPositionPack(t *testing.T) {
	// x (-33554432 to 33554431), y (-2048 to 2047), z (-33554432 to 33554431)

	for x := -33554432; x < 33554432; x += 55443 {
		for y := -2048; y < 2048; y += 48 {
			for z := -33554432; z < 33554432; z += 55443 {
				p := PackPosition(x, y, z)
				x1, y1, z1, _ := UnpackPosition(bytes.NewReader(p))
				if x1 != x || y1 != y || z1 != z {
					t.Errorf("cannot pack (%d, %d, %d)", x, y, z)
				}
			}
		}
	}
}
