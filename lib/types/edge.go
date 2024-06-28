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
}

func (em *EdgeMeta) Serialize() ([]byte, error) {
	return cbor.Marshal(em)
}

func (em *EdgeMeta) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, em)
}

type IEdge interface {
	Register(context.Context, EdgeMeta) error
	GetEdge(context.Context, common.Address) (EdgeReceipt, error)
	ListEdge(context.Context, Options) ([]EdgeReceipt, error)
}
