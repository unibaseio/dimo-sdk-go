package types

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type PromptMeta struct {
	Space  string
	Serial uint64 // second
	Prompt string // needle meta
}

func (pm *PromptMeta) Serialize() ([]byte, error) {
	return cbor.Marshal(pm)
}

func (pm *PromptMeta) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, pm)
}

type IPrompt interface {
	SubmitPrompt(ctx context.Context, addr common.Address, pm PromptMeta) (PromptMeta, error)
	GetPrompt(ctx context.Context, sn string, serial uint64) (PromptMeta, error)
	ListPrompt(ctx context.Context, addr common.Address, opt Options) ([]PromptMeta, error)
}

type Text2ImageInference struct {
	Prompt   string
	Negative string
	Width    int
	Height   int
	Guidance int
	Steps    int
	Seed     *big.Int
}
