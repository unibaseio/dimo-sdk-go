package types

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type GPUDetail struct {
	Type   string
	Memory string
}

type GPUCore struct {
	GPUDetail
	Name  string
	Owner common.Address
}

func (mm *GPUCore) Serialize() ([]byte, error) {
	return cbor.Marshal(mm)
}

func (mm *GPUCore) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, mm)
}

type GPUMeta struct {
	GPUCore
	OnChain bool
	Active  bool
	InUse   bool
}

func (gm *GPUMeta) Serialize() ([]byte, error) {
	return cbor.Marshal(gm)
}

func (gm *GPUMeta) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, gm)
}

type IGPU interface {
	Submit(ctx context.Context, addr common.Address, gc GPUCore) (GPUMeta, error)
	Get(ctx context.Context, gn string) (GPUMeta, error)
	List(ctx context.Context, addr common.Address, opt Options) ([]GPUMeta, error)
}
