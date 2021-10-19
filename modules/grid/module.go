package grid

import (
	"github.com/forbole/juno/v2/modules"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v2/modules/messages"
	"github.com/forbole/juno/v2/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cybercongress/cyberindex/database"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
)

type Module struct {
	messagesParser messages.MessageAddressesParser
	cdc 		   codec.Codec
	db             *database.CyberDb
}

func NewModule(
	messagesParser messages.MessageAddressesParser,
	cdc 		   codec.Codec,
	db 			   *database.CyberDb,
) *Module {
	return &Module{
		messagesParser: messagesParser,
		cdc: 			cdc,
		db:             db,
	}
}

func (m *Module) Name() string {
	return "grid"
}

func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *types.Tx) error {
	return HandleMsg(tx, msg, m.db)
}
