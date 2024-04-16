package contract

import (
	"encoding/hex"
	"fmt"

	"github.com/MOSSV2/dimo-sdk-go/lib/bls"
	"github.com/ethereum/go-ethereum/common"
)

const (
	mUncompressed byte = 0b000 << 5
)

func StringToSolByte(s string) ([]byte, error) {
	gb, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return ByteToSolByte(gb)
}

func ByteToSolByte(gb []byte) ([]byte, error) {
	switch len(gb) {
	case 32:
		return gb, nil
	case 48:
		var g1 bls.G1
		_, err := g1.SetBytes(gb)
		if err != nil {
			return nil, err
		}
		return G1ToSolByte(g1), nil
	case 96:
		var g2 bls.G2
		_, err := g2.SetBytes(gb)
		if err != nil {
			return nil, err
		}
		return G2ToSolByte(g2), nil
	}
	return nil, fmt.Errorf("invalid length")
}

func G1ToSolByte(g1 bls.G1) []byte {
	x := common.LeftPadBytes(g1.X.Marshal(), 64)
	y := common.LeftPadBytes(g1.Y.Marshal(), 64)
	x = append(x, y...)
	return x
}

func G2ToSolByte(g2 bls.G2) []byte {
	x := common.LeftPadBytes(g2.X.A1.Marshal(), 64)
	x0 := common.LeftPadBytes(g2.X.A0.Marshal(), 64)
	x = append(x, x0...)
	y1 := common.LeftPadBytes(g2.Y.A1.Marshal(), 64)
	x = append(x, y1...)
	y0 := common.LeftPadBytes(g2.Y.A0.Marshal(), 64)
	x = append(x, y0...)
	return x
}

func SolByteToG1(gb []byte) (bls.G1, error) {
	var g1 bls.G1
	if len(gb) != 128 {
		return g1, fmt.Errorf("invalid length")
	}
	ngb := make([]byte, bls.G1Size*2)
	copy(ngb, gb[16:64])
	copy(ngb[48:], gb[80:128])
	ngb[0] |= mUncompressed
	_, err := g1.SetBytes(ngb)
	return g1, err
}

func SolByteToG2(gb []byte) (bls.G2, error) {
	var g2 bls.G2
	if len(gb) != 256 {
		return g2, fmt.Errorf("invalid length")
	}
	ngb := make([]byte, bls.G2Size*2)
	copy(ngb, gb[16:64])
	copy(ngb[48:], gb[80:128])
	copy(ngb[96:], gb[144:192])
	copy(ngb[144:], gb[208:256])

	ngb[0] |= mUncompressed
	_, err := g2.SetBytes(ngb)
	return g2, err
}

func SolByteToString(gb []byte) (string, error) {
	switch len(gb) {
	case 32:
		return hex.EncodeToString(gb), nil
	case 128:
		g1, err := SolByteToG1(gb)
		if err != nil {
			return "", err
		}
		ngb := g1.Bytes()
		return hex.EncodeToString(ngb[:]), nil
	case 256:
		g2, err := SolByteToG2(gb)
		if err != nil {
			return "", err
		}
		ngb := g2.Bytes()
		return hex.EncodeToString(ngb[:]), nil
	}
	return "", fmt.Errorf("invalid length")
}
