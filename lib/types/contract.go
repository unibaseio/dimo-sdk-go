package types

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type IContract interface {
	BalanceOf(ctx context.Context, addr common.Address) *big.Int
	Recharge(ctx context.Context, addr common.Address, val *big.Int) error
}

type EpochInfo struct {
	Epoch  uint64
	Height *big.Int
	Seed   [32]byte
}

func (ei *EpochInfo) Serialize() ([]byte, error) {
	return cbor.Marshal(ei)
}

func (ei *EpochInfo) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, ei)
}

type RSChalInChain struct {
	Store   common.Address
	Replica uint64
}

func (ei *RSChalInChain) Serialize() ([]byte, error) {
	return cbor.Marshal(ei)
}

func (ei *RSChalInChain) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, ei)
}

type EPChalInChain struct {
	Store  common.Address
	Epoch  uint64
	Round  uint8
	QIndex uint8
	Coms   [][]byte
}

type ReplicaInChain struct {
	Name     string
	Serial   uint64
	Piece    uint64
	Index    uint8
	StoredOn common.Address
	Witness  ReplicaWitness
}

type EProofInChain struct {
	Epoch uint64
	Store common.Address
	Hash  []byte
	Sum   []byte
	H     []byte
	Value []byte
}

func (ei *EProofInChain) Serialize() ([]byte, error) {
	return cbor.Marshal(ei)
}

func (ei *EProofInChain) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, ei)
}
