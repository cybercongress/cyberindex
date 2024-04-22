package source

import (
	"github.com/gravity-devs/liquidity/x/liquidity/types"
)

type Source interface {
	GetPool(poolID uint64, height int64) (types.Pool, error)
}
