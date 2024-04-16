package types

import (
	"context"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type NFTMeta struct {
	Name   string
	Size   uint64
	Type   string // png for image
	Prompt uint64
	Model  string
	Space  string
	Price  *big.Int
	Owner  common.Address
}

func (nm *NFTMeta) Serialize() ([]byte, error) {
	return cbor.Marshal(nm)
}

func (nm *NFTMeta) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, nm)
}

type INFT interface {
	UploadNFT(ctx context.Context, addr common.Address, nm NFTMeta, r io.Reader) (NFTMeta, error)
	BuyNFT(ctx context.Context, addr common.Address, nftName string) error
	GetNFT(ctx context.Context, nftName string) (NFTMeta, error)
	ListNFT(ctx context.Context, addr common.Address, opt Options) ([]NFTMeta, error)
	GetNFTData(ctx context.Context, addr common.Address, nftName string, w io.Writer) (NFTMeta, error)
}
