package bls

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

func Encrypt(in, v []byte) error {
	if len(v) != 32 {
		return fmt.Errorf("invalid key")
	}

	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}

	var fr_v, fr_e, fr_c, fr_tmp Fr
	nv := sha256.Sum256(v[:20])
	fr_v.SetBytes(nv[:])
	fr_big := new(big.Int)
	fr_v.BigInt(fr_big)
	fr_e.Exp(fr_v, fr_big)
	fr_c.SetBytes(v)

	for i := 0; i < cnt; i++ {
		fr_tmp.SetBytes(in[PadSize*i : PadSize*(i+1)])
		fr_e.Mul(&fr_e, &fr_v)
		fr_tmp.Add(&fr_tmp, &fr_e)
		fr_tmp.Add(&fr_tmp, &fr_c)
		copy(in[PadSize*i:PadSize*(i+1)], fr_tmp.Marshal())
	}

	return nil
}

func Decrypt(in, v []byte) error {
	if len(v) != 32 {
		return fmt.Errorf("invalid key")
	}

	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}

	var fr_v, fr_e, fr_c, fr_tmp Fr
	nv := sha256.Sum256(v[:20])
	fr_v.SetBytes(nv[:])
	fr_big := new(big.Int)
	fr_v.BigInt(fr_big)
	fr_e.Exp(fr_v, fr_big)
	fr_c.SetBytes(v)

	for i := 0; i < cnt; i++ {
		fr_tmp.SetBytes(in[PadSize*i : PadSize*(i+1)])
		fr_e.Mul(&fr_e, &fr_v)
		fr_tmp.Sub(&fr_tmp, &fr_e)
		fr_tmp.Sub(&fr_tmp, &fr_c)
		copy(in[PadSize*i:PadSize*(i+1)], fr_tmp.Marshal())
	}
	return nil
}
