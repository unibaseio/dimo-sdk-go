package lighthash

import (
	"errors"
	"hash"
	"math/big"
	"sync"

	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"golang.org/x/crypto/sha3"
)

const (
	mimcNbRounds = 2
	seed         = "seed"   // seed to derive the constants
	BlockSize    = fr.Bytes // BlockSize size that mimc consumes
)

type Fr = fr.Element

var (
	mimcConstants [mimcNbRounds]Fr
	once          sync.Once
)

type digest struct {
	h    Fr
	data []Fr
}

func GetConstants() []big.Int {
	once.Do(initConstants) // init constants
	res := make([]big.Int, mimcNbRounds)
	for i := 0; i < mimcNbRounds; i++ {
		mimcConstants[i].BigInt(&res[i])
	}
	return res
}

func New() hash.Hash {
	d := new(digest)
	d.Reset()
	return d
}

func (d *digest) Reset() {
	d.data = d.data[:0]
	d.h = Fr{0, 0, 0, 0}
}

func (d *digest) Sum(b []byte) []byte {
	buffer := d.checksum()
	d.data = nil // flush the data already hashed
	hash := buffer.Bytes()
	b = append(b, hash[:]...)
	return b
}

func (d *digest) Size() int {
	return BlockSize
}

func (d *digest) BlockSize() int {
	return BlockSize
}

func (d *digest) Write(p []byte) (int, error) {

	var start int
	for start = 0; start < len(p); start += BlockSize {
		if elem, err := fr.BigEndian.Element((*[BlockSize]byte)(p[start : start+BlockSize])); err == nil {
			d.data = append(d.data, elem)
		} else {
			return 0, err
		}
	}

	if start != len(p) {
		return 0, errors.New("invalid input length: must represent a list of field elements, expects a []byte of len m*BlockSize")
	}
	return len(p), nil
}

func (d *digest) checksum() Fr {
	for i := range d.data {
		r := d.encrypt(d.data[i])
		d.h.Add(&r, &d.h).Add(&d.h, &d.data[i])
	}

	return d.h
}

func (d *digest) encrypt(m Fr) Fr {
	once.Do(initConstants) // init constants

	var tmp Fr
	for i := 0; i < mimcNbRounds; i++ {
		// m = (m+k+c)^**5
		tmp.Add(&m, &d.h).Add(&tmp, &mimcConstants[i])
		m.Square(&tmp).
			Square(&m).
			Mul(&m, &tmp)
	}
	m.Add(&m, &d.h)
	return m
}

func Sum(msg []byte) ([]byte, error) {
	var d digest
	if _, err := d.Write(msg); err != nil {
		return nil, err
	}
	h := d.checksum()
	bytes := h.Bytes()
	return bytes[:], nil
}

func initConstants() {
	bseed := ([]byte)(seed)

	hash := sha3.NewLegacyKeccak256()
	_, _ = hash.Write(bseed)
	rnd := hash.Sum(nil) // pre hash before use
	hash.Reset()
	_, _ = hash.Write(rnd)

	for i := 0; i < mimcNbRounds; i++ {
		rnd = hash.Sum(nil)
		mimcConstants[i].SetBytes(rnd)
		//log.Println(mimcConstants[i].Text(16))
		hash.Reset()
		_, _ = hash.Write(rnd)
	}
}

func (d *digest) WriteString(rawBytes []byte) {
	if elems, err := fr.Hash(rawBytes, []byte("string:"), 1); err != nil {
		panic(err)
	} else {
		d.data = append(d.data, elems[0])
	}
}

func HashHint(inputs []*big.Int) (*big.Int, error) {
	mod, _ := new(big.Int).SetString("73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001", 16)
	param, _ := new(big.Int).SetString("1dbfc7763d69ca7d15701422f37bc6692bd01ebc4da42360f81f9adb4a91b01a", 16)
	param1, _ := new(big.Int).SetString("4fd2cddd334dab1c4005161c290f25a0e18d4175ecfa898b17095d8ec2dd344a", 16)

	h := new(big.Int)

	for i := 0; i < len(inputs); i++ {
		tmp := new(big.Int).Add(inputs[i], h)
		tmp.Add(tmp, param)
		tmp.Exp(tmp, big.NewInt(5), mod)

		tmp.Add(tmp, h)
		tmp.Add(tmp, param1)
		tmp.Exp(tmp, big.NewInt(5), mod)

		tmp.Add(tmp, h)

		h.Add(h, tmp)
		h.Add(h, inputs[i])
	}
	h.Mod(h, mod)
	return h, nil
}
