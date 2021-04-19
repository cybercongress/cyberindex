package graph

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	graphtypes "github.com/cybercongress/go-cyber/x/graph/types"
	"google.golang.org/grpc"

	"github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cybercongress/cyberindex/database"
)

var _ modules.Module = &Module{}

type Module struct {
	encodingConfig *params.EncodingConfig
	graphClient    graphtypes.QueryClient
	db             *database.CyberDb
}

func NewModule(encodingConfig *params.EncodingConfig, grpcConnection *grpc.ClientConn, db *database.CyberDb) *Module {
	return &Module{
		encodingConfig: encodingConfig,
		graphClient:    graphtypes.NewQueryClient(grpcConnection),
		db:             db,
	}
}

func (m *Module) Name() string {
	return "graph"
}

func (m *Module) HandleGenesis(_ *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	return nil
}

func (m *Module) HandleBlock(block *tmctypes.ResultBlock, _ []*types.Tx, _ *tmctypes.ResultValidators) error {
	return nil
}

func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *types.Tx) error {
	return HandleMsg(tx, msg, m.encodingConfig.Marshaler, m.db)
}
