package types

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

type IProxy interface {
	Get(ctx context.Context, addr common.Address) (common.Address, error)
	GetProxy(ctx context.Context, addr common.Address) common.Address
	Update(ctx context.Context, addr, paddr common.Address) error
}
