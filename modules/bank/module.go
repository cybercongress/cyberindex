package bank

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cybercongress/cyberindex/database"
	"github.com/cybercongress/cyberindex/modules/bank/source"

	junomessages "github.com/forbole/juno/v3/modules/messages"

	"github.com/forbole/juno/v3/modules"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.BlockModule   = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/bank module
type Module struct {
	messageParser junomessages.MessageAddressesParser
	cdc           codec.Codec
	keeper        source.Source
	db            *database.CyberDb
}

// NewModule returns a new Module instance
func NewModule(
	messageParser junomessages.MessageAddressesParser,
	cdc codec.Codec,
	keeper source.Source,
	db *database.CyberDb,
) *Module {
	return &Module{
		messageParser: messageParser,
		cdc:           cdc,
		keeper:        keeper,
		db:            db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "bank"
}
