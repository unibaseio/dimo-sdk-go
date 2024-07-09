package types

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

const (
	StoreType     = "store"
	StreamType    = "stream"
	ValidatorType = "validator"
	ComputeType   = "compute"
	GatewayType   = "gateway"
	ClientType    = "client"
	HubType       = "hub"
)

type EdgeReceipt struct {
	EdgeMeta
	OnChain bool
	Revenue *big.Int
	Last    time.Time
}

type EdgeMeta struct {
	Type      string // "store", "compute"
	Name      common.Address
	PublicKey []byte
	ExposeURL string
	Hardware  HardwareInfo
}

func (em *EdgeMeta) Serialize() ([]byte, error) {
	return cbor.Marshal(em)
}

func (em *EdgeMeta) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, em)
}

type HardwareInfo struct {
	CPU    string
	Memory string
}

func (gm *HardwareInfo) Serialize() ([]byte, error) {
	return cbor.Marshal(gm)
}

func (gm *HardwareInfo) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, gm)
}

type IEdge interface {
	Register(context.Context, EdgeMeta) error
	GetEdge(context.Context, common.Address) (EdgeReceipt, error)
	ListEdge(context.Context, Options) ([]EdgeReceipt, error)
}
