package remote

import (
	"fmt"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	liquiditykeeper "github.com/cybercongress/cyberindex/modules/liquidity/source"
	"github.com/forbole/juno/v3/node/remote"
	liquiditytypes "github.com/tendermint/liquidity/x/liquidity/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
)

var (
	_ liquiditykeeper.Source = &Source{}
)

type Source struct {
	*remote.Source
	liquidityClient liquiditytypes.QueryClient
}

func NewSource(source *remote.Source, liquidityClient liquiditytypes.QueryClient) *Source {
	return &Source{
		Source:          source,
		liquidityClient: liquidityClient,
	}
}

func (s Source) GetPool(poolID uint64, height int64) (liquiditytypes.Pool, error) {
	header := GetHeightRequestHeader(height)
	res, err := s.liquidityClient.LiquidityPool(s.Ctx, &liquiditytypes.QueryLiquidityPoolRequest{PoolId: poolID}, header)
	if err != nil {
		return liquiditytypes.Pool{}, fmt.Errorf("error while getting pool id: %s", err)
	}

	return res.Pool, nil
}

func (s Source) GetAllPools(height int64) ([]liquiditytypes.Pool, error) {
	header := GetHeightRequestHeader(height)
	res, err := s.liquidityClient.LiquidityPools(s.Ctx, &liquiditytypes.QueryLiquidityPoolsRequest{}, header)
	if err != nil {
		panic(err)
	}

	return res.Pools, err
}

func GetHeightRequestHeader(height int64) grpc.CallOption {
	header := metadata.New(map[string]string{
		grpctypes.GRPCBlockHeightHeader: strconv.FormatInt(height, 10),
	})
	return grpc.Header(&header)
}
