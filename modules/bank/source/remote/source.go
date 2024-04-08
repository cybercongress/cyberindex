package remote

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/forbole/juno/v3/node/remote"

	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	bankkeeper "github.com/cybercongress/cyberindex/v2/modules/bank/source"
	"github.com/forbole/bdjuno/v3/types"
)

var (
	_ bankkeeper.Source = &Source{}
)

type Source struct {
	*remote.Source
	bankClient banktypes.QueryClient
}

// NewSource builds a new Source instance
func NewSource(source *remote.Source, bankClient banktypes.QueryClient) *Source {
	return &Source{
		Source:     source,
		bankClient: bankClient,
	}
}

// GetBalances implements bankkeeper.Source
func (s Source) GetBalances(addresses []string, height int64) ([]types.AccountBalance, error) {
	header := GetHeightRequestHeader(height)

	var balances []types.AccountBalance
	for _, address := range addresses {
		balRes, err := s.bankClient.AllBalances(s.Ctx, &banktypes.QueryAllBalancesRequest{Address: address}, header)
		if err != nil {
			return nil, fmt.Errorf("error while getting all balances: %s", err)
		}

		balances = append(balances, types.NewAccountBalance(
			address,
			balRes.Balances,
			height,
		))
	}

	return balances, nil
}

// GetSupply implements bankkeeper.Source
func (s Source) GetSupply(height int64) (sdk.Coins, error) {
	header := GetHeightRequestHeader(height)
	res, err := s.bankClient.TotalSupply(s.Ctx, &banktypes.QueryTotalSupplyRequest{}, header)
	if err != nil {
		return nil, fmt.Errorf("error while getting total supply: %s", err)
	}

	return res.Supply, nil
}

// TODO move outside to utils
func GetHeightRequestHeader(height int64) grpc.CallOption {
	header := metadata.New(map[string]string{
		grpctypes.GRPCBlockHeightHeader: strconv.FormatInt(height, 10),
	})
	return grpc.Header(&header)
}
