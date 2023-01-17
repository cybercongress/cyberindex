package source

import (
	"github.com/tendermint/liquidity/x/liquidity/types"
)

type Source interface {
	GetPool(poolID uint64, height int64) (types.Pool, error)
}
