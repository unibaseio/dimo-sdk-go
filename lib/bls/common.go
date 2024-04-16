package bls

import (
	"fmt"
	"hash"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"

	bls "github.com/consensys/gnark-crypto/ecc/bls12-377"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
	"github.com/consensys/gnark-crypto/ecc/bw6-761/fr/mimc"
)

const (
	UnPadSize = 31
	PadSize   = 32
	PadByte   = 0x01 // incase zero
	LogShard  = 20
	MinShard  = 2
	MaxShard  = 1 << LogShard
	MaxSize   = MaxShard * UnPadSize // per piece
	SlothDec  = 17
	SlothEnc  = "8c920a4fd422fb02d82062c67a1bc3cdee6ce960787878804bcf0f0f0f0f0f1"
)

type Fr = fr.Element
type G1 = bls.G1Affine
type G2 = bls.G2Affine
type GT = bls.GT
type G1Jac = bls.G1Jac
type G2Jac = bls.G2Jac

const (
	FrSize    = fr.Bytes
	G1Size    = bls.SizeOfG1AffineCompressed
	G2Size    = bls.SizeOfG2AffineCompressed
	GTSize    = bls.SizeOfGT
	ProofSize = FrSize + G1Size
)

func NewFieldHash() hash.Hash {
	return mimc.NewMiMC()
}

var _ types.ICommitment = (*Commitment)(nil)

type Commitment struct {
	Value G1
}

func NewCommitment(b []byte) (types.ICommitment, error) {
	var val G1
	if b != nil {
		_, err := val.SetBytes(b)
		if err != nil {
			return nil, err
		}
	}

	return &Commitment{val}, nil
}

func (c *Commitment) Add(ic types.ICommitment) error {
	com, ok := ic.(*Commitment)
	if !ok {
		return fmt.Errorf("wrong commitment")
	}
	c.Value.Add(&c.Value, &com.Value)
	return nil
}

func (c *Commitment) Serialize() []byte {
	cbyte := c.Value.Bytes()
	return cbyte[:]
}

func (c *Commitment) Raw() []byte {
	return c.Value.Marshal()
}

var _ types.IChallenge = (*Challenge)(nil)

type Challenge struct {
	typ        int
	RandomInt  int
	RandomByte []byte
	Sum        G1
	Digests    []G1
}

func NewChallenge(r []byte) types.IChallenge {
	chal := &Challenge{
		RandomByte: r,
		Digests:    make([]G1, 0, 1),
	}
	return chal
}

func NewPointChallenge(val int) types.IChallenge {
	chal := &Challenge{
		RandomInt: val,
		Digests:   make([]G1, 0, 1),
	}
	return chal
}

func (ch *Challenge) Type() int {
	return ch.typ
}

func (ch *Challenge) Add(ic types.ICommitment) error {
	com, ok := ic.(*Commitment)
	if !ok {
		return fmt.Errorf("wrong commitment")
	}

	ch.Sum.Add(&ch.Sum, &com.Value)
	ch.Digests = append(ch.Digests, com.Value)
	return nil
}

func (ch *Challenge) Commitment() types.ICommitment {
	return &Commitment{
		ch.Sum,
	}
}

type MultiProof struct {
	H            G1
	ClaimedValue []Fr
}

var _ types.IProof = (*Proof)(nil)

type Proof struct {
	// H quotient polynomial (f - f(r)))/(x-r)
	H G1

	// ClaimedValue purported value
	ClaimedValue Fr
}

func NewProof(buf []byte) (*Proof, error) {
	pf := Proof{}
	if len(buf) != ProofSize {
		return nil, fmt.Errorf("wrong size: %d, need %d", len(buf), ProofSize)
	}

	pf.ClaimedValue.SetBytes(buf[:FrSize])
	_, err := pf.H.SetBytes(buf[FrSize : FrSize+G1Size])
	if err != nil {
		return nil, err
	}
	return &pf, nil
}

func (pf *Proof) Type() int {
	return 1
}

func (pf *Proof) Add(ip types.IProof) error {
	pr, ok := ip.(*Proof)
	if !ok {
		return fmt.Errorf("wrong proof1")
	}
	pf.ClaimedValue.Add(&pf.ClaimedValue, &pr.ClaimedValue)
	pf.H.Add(&pf.H, &pr.H)
	return nil
}

func (pf *Proof) Serialize() []byte {
	buf := make([]byte, FrSize+G1Size)
	copy(buf[:FrSize], pf.ClaimedValue.Marshal())
	hbyte := pf.H.Bytes()
	copy(buf[FrSize:FrSize+G1Size], hbyte[:])

	return buf
}
