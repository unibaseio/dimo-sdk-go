package bls

import (
	"fmt"
	"math/big"

	"github.com/MOSSV2/dimo-sdk-go/lib/lighthash"
)

func SlothEncode(in, v []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}

	var tmp, fr_v Fr
	fr_v.SetBytes(v)
	big_inv, _ := new(big.Int).SetString(SlothEnc, 16)
	for i := 0; i < cnt; i++ {
		tmp.SetBytes(in[PadSize*i : PadSize*(i+1)])
		tmp.Add(&tmp, &fr_v)
		tmp.Exp(tmp, big_inv)
		copy(in[PadSize*i:PadSize*(i+1)], tmp.Marshal())
	}
	return nil
}

func SlothDecode(in, v []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}
	var tmp, fr_v Fr
	fr_v.SetBytes(v)
	big_e := big.NewInt(SlothDec)
	for i := 0; i < cnt; i++ {
		tmp.SetBytes(in[PadSize*i : PadSize*(i+1)])
		tmp.Exp(tmp, big_e)
		tmp.Sub(&tmp, &fr_v)
		copy(in[PadSize*i:PadSize*(i+1)], tmp.Marshal())
	}
	return nil
}

// 32MB: 10.5s- 12s
func SlothEncodeV2(in, v []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}

	h := lighthash.New()
	var tmp, tmpv, fr_v Fr
	fr_v.SetBytes(v)
	big_inv, _ := new(big.Int).SetString(SlothEnc, 16)
	for i := 0; i < cnt; i++ {
		h.Reset()
		h.Write(v)
		tmp.SetInt64(int64(i))
		h.Write(tmp.Marshal())
		tmpv.SetBytes(h.Sum(nil))
		tmp.SetBytes(in[PadSize*i : PadSize*(i+1)])
		tmp.Add(&tmp, &tmpv)
		tmp.Exp(tmp, big_inv)
		copy(in[PadSize*i:PadSize*(i+1)], tmp.Marshal())
	}
	return nil
}

// 32MB: 350ms -> 1.8s
func SlothDecodeV2(in, v []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}

	h := lighthash.New()
	var tmp, tmpv, fr_v Fr
	fr_v.SetBytes(v)
	big_e := big.NewInt(SlothDec)
	for i := 0; i < cnt; i++ {
		h.Reset()
		h.Write(v)
		tmp.SetInt64(int64(i))
		h.Write(tmp.Marshal())
		tmpv.SetBytes(h.Sum(nil))

		tmp.SetBytes(in[PadSize*i : PadSize*(i+1)])
		tmp.Exp(tmp, big_e)
		tmp.Sub(&tmp, &tmpv)
		copy(in[PadSize*i:PadSize*(i+1)], tmp.Marshal())
	}
	return nil
}
