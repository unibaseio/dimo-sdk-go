package bls

import (
	"fmt"
)

// treat in as big endian here
func Pad(in []byte) []byte {
	cnt := 1 + (len(in)-1)/UnPadSize
	out := make([]byte, cnt*PadSize)
	for i := 0; i < cnt-1; i++ {
		out[PadSize*i] = PadByte // incase zero
		copy(out[PadSize*i+1:PadSize*(i+1)], in[UnPadSize*i:UnPadSize*(i+1)])
	}

	// handle last one
	out[PadSize*(cnt-1)] = PadByte
	copy(out[PadSize*(cnt-1)+1:PadSize*cnt], in[UnPadSize*(cnt-1):])

	return out
}

func Unpad(in []byte) ([]byte, error) {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return nil, fmt.Errorf("invalid unpad length %d , should align to %d", len(in), PadSize)
	}

	out := make([]byte, UnPadSize*cnt)
	for i := 0; i < cnt; i++ {
		copy(out[UnPadSize*i:UnPadSize*(i+1)], in[PadSize*i+1:PadSize*(i+1)])
	}
	return out, nil
}

func Split(slen int, data []byte) []Fr {
	switch slen {
	case 31:
		data = Pad(data)
	default:
		if len(data)/32*32 != len(data) {
			panic("data length should align to 32")
		}
	}

	num := len(data) / 32
	atom := make([]Fr, num)
	for i := 0; i < num; i++ {
		atom[i].SetBytes(data[32*i : 32*(i+1)])
	}
	return atom
}
