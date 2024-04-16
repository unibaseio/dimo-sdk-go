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
	rest := cnt*UnPadSize - len(in)
	out[PadSize*(cnt-1)] = PadByte
	copy(out[PadSize*(cnt-1)+1+rest:PadSize*cnt], in[UnPadSize*(cnt-1):])

	// one element is not enough, padding it
	if cnt == 1 {
		padding := make([]byte, PadSize)
		padding[0] = PadByte
		out = append(out, padding...)
	}

	return out
}

func Unpad(in []byte, size int) ([]byte, error) {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return nil, fmt.Errorf("invalid unpad length, should align to %d", PadSize)
	}
	if size > PadSize*cnt {
		return nil, fmt.Errorf("invalid unpad size, larger than len(in)")
	}

	unpadcnt := 1 + (size-1)/UnPadSize
	out := make([]byte, size)
	for i := 0; i < unpadcnt-1; i++ {
		copy(out[UnPadSize*i:UnPadSize*(i+1)], in[PadSize*i+1:PadSize*(i+1)])
	}
	rest := unpadcnt*UnPadSize - size
	copy(out[UnPadSize*(unpadcnt-1):], in[PadSize*(unpadcnt-1)+1+rest:PadSize*unpadcnt])
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
