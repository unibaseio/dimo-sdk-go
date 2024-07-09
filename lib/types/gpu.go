package types

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type GPUMeta struct {
	GPUCore
	Active bool
	InUse  bool
}

type GPUCore struct {
	Name   string
	Serial uint64
	Owner  common.Address
	Type   string
	Memory string
}

func (gm *GPUCore) Serialize() ([]byte, error) {
	return cbor.Marshal(gm)
}

func (gm *GPUCore) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, gm)
}

type IGPU interface {
	Submit(ctx context.Context, addr common.Address, gc GPUCore) error
	Get(ctx context.Context, gn string) (GPUMeta, error)
	List(ctx context.Context, addr common.Address, opt Options) ([]GPUMeta, error)
}
