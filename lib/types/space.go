package types

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type SpaceStat struct {
	Serial uint64
	Stat   uint8     // 0 init, 1. active
	Last   time.Time // heartbeat
}

type SpaceMeta struct {
	SpaceStat
	Name     string // hash(ModelName|Owner|Runner)
	Model    string
	GPU      string
	Root     string
	Runtime  string // docker image name or docker file hash
	Price    *big.Int
	Owner    common.Address
	Runner   common.Address
	Creation time.Time
}

func (msm *SpaceMeta) Serialize() ([]byte, error) {
	return cbor.Marshal(msm)
}

func (msm *SpaceMeta) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, msm)
}

type ISpace interface {
	Submit(ctx context.Context, addr common.Address, ms SpaceMeta) (SpaceMeta, error)
	Confirm(ctx context.Context, addr common.Address, ms, root string, pf []byte) error
	Update(ctx context.Context, addr common.Address, name string) error
	SetRuntime(ctx context.Context, addr common.Address, name, runtime string) error
	Get(ctx context.Context, addr common.Address, name string) (SpaceResult, error)
	List(ctx context.Context, addr common.Address, opt Options) ([]SpaceResult, error)
}

// save in data
type SpaceObject struct {
	Space    string
	Creation int64
	Prompt   []byte
	Result   []byte
}

func (pm *SpaceObject) Serialize() ([]byte, error) {
	return cbor.Marshal(pm)
}

func (pm *SpaceObject) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, pm)
}
