package types

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type ModelCore struct {
	Name     string
	Size     uint64
	Creation time.Time
}

func (mm *ModelCore) Serialize() ([]byte, error) {
	return cbor.Marshal(mm)
}

func (mm *ModelCore) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, mm)
}

type ModelMeta struct {
	ModelCore
	OnChain bool
	Score   uint64
	Price   *big.Int
	Owner   common.Address
	Type    string // stable-diffusion or others
	Hash    string
	Count   int
	Files   map[string]string // name->content; if content is sha256, download it or put it
}

func (mmr *ModelMeta) Serialize() ([]byte, error) {
	return cbor.Marshal(mmr)
}

func (mmr *ModelMeta) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, mmr)
}

type IModel interface {
	SubmitModel(ctx context.Context, addr common.Address, pm ModelMeta) (ModelMeta, error)
	BuyModel(ctx context.Context, addr common.Address, modelName string) error
	GetModel(ctx context.Context, addr common.Address, mr string) (ModelResult, error)
	ListModel(ctx context.Context, addr common.Address, opt Options) ([]ModelResult, error)
}

type RepoCore struct {
	Type     string
	Name     string
	Revision string
}

func (rc *RepoCore) ID() string {
	if rc.Revision == "" {
		return rc.Name
	} else {
		return rc.Name + "_" + rc.Revision
	}
}

func (rc *RepoCore) Hash() string {
	h := sha256.New()
	h.Write([]byte(rc.Name))
	h.Write([]byte(rc.Revision))
	rcname := hex.EncodeToString(h.Sum(nil))
	return rcname[:12]
}

func (rc *RepoCore) Serialize() ([]byte, error) {
	return cbor.Marshal(rc)
}

func (rc *RepoCore) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, rc)
}

type RepoMeta struct {
	RepoCore
	State int // 0: downloading, 1: fail,2: downloaded; 3: uploaded to stream;
}

func (rm *RepoMeta) Serialize() ([]byte, error) {
	return cbor.Marshal(rm)
}

func (rm *RepoMeta) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, rm)
}
