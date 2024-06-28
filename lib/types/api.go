package types

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type Options struct {
	UserDefined map[string]string
}

type Auth struct {
	Type string
	Addr common.Address
	Time int64
	Hash []byte
	Sign []byte
}

type ModelResult struct {
	ModelMeta
	IsPaid bool
}

type SpaceResult struct {
	SpaceMeta
	IsPaid bool
}

type ListModelResult struct {
	Models []ModelResult
}

type ListSpaceResult struct {
	Spaces []SpaceResult
}

type ListGPUResult struct {
	GPUs []GPUMeta
}

type AccountResult struct {
	Name  common.Address
	Value *big.Int
}

type ChalResult struct {
	Epoch  int64
	Random [32]byte
}

type ListPieceResult struct {
	Pieces []PieceReceipt
}

type ListFileResult struct {
	Files []FileReceipt
}

type ListEdgeResult struct {
	Edges []EdgeReceipt
}

type ListReplicaResult struct {
	Replicas []ReplicaCore
}

type ListStatResult struct {
	Stats []StatResult
}

type ListObjectResult struct {
	IsTruncated bool

	NextMarker string

	Objects []ObjectInfo

	Prefixes []string
}

type StatResult struct {
	Value int64
	Time  time.Time
}

func (hr *StatResult) Serialize() ([]byte, error) {
	return cbor.Marshal(hr)
}

func (hr *StatResult) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, hr)
}
