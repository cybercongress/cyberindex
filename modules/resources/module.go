package resources

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	resourcestypes "github.com/cybercongress/go-cyber/x/resources/types"
	"github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/modules/messages"
	"github.com/desmos-labs/juno/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cybercongress/cyberindex/database"
)

var _ modules.Module = &Module{}

type Module struct {
	messagesParser  messages.MessageAddressesParser
	encodingConfig  *params.EncodingConfig
	resourcesClient resourcestypes.QueryClient
	db              *database.CyberDb
}

func NewModule(
	messagesParser messages.MessageAddressesParser,
	resourcesClient resourcestypes.QueryClient,
	encodingConfig *params.EncodingConfig,
	db *database.CyberDb,
) *Module {
	return &Module{
		messagesParser:  messagesParser,
		encodingConfig:  encodingConfig,
		resourcesClient: resourcesClient,
		db:              db,
	}
}

func (m *Module) Name() string {
	return "resources"
}

func (m *Module) DownloadState(height int64) error {
	return nil
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
