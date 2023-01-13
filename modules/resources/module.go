package resources

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cybercongress/cyberindex/database"
	"github.com/forbole/juno/v3/modules"
	"github.com/forbole/juno/v3/modules/messages"
	"github.com/forbole/juno/v3/types"
)

var _ modules.Module = &Module{}

type Module struct {
	messagesParser messages.MessageAddressesParser
	cdc            codec.Codec
	db             *database.CyberDb
}

func NewModule(
	messagesParser messages.MessageAddressesParser,
	cdc codec.Codec,
	db *database.CyberDb,
) *Module {
	return &Module{
		messagesParser: messagesParser,
		cdc:            cdc,
		db:             db,
	}
}

func (m *Module) Name() string {
	return "resources"
}

func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *types.Tx) error {
	return HandleMsg(tx, msg, m.db)
}
