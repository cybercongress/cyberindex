package liquidity

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cybercongress/cyberindex/database"
	"github.com/forbole/juno/v2/modules"
	"github.com/forbole/juno/v2/modules/messages"
	"github.com/forbole/juno/v2/types"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

var (
	_ modules.Module 		= &Module{}
	_ modules.MessageModule = &Module{}
	_ modules.BlockModule   = &Module{}
	_ modules.PeriodicOperationsModule
)

type Module struct {
	messagesParser  messages.MessageAddressesParser
	cdc 		   	codec.Codec
	db              *database.CyberDb
	bankModule      BankModule
	authModule      AuthModule
}

func NewModule(
	messagesParser messages.MessageAddressesParser,
	cdc 		   codec.Codec,
	bankModule 	   BankModule,
	authModule     AuthModule,
	db 			   *database.CyberDb,
) *Module {
	return &Module{
		messagesParser:  messagesParser,
		cdc: 			 cdc,
		db:              db,
		bankModule:      bankModule,
		authModule:      authModule,
	}
}

func (m *Module) Name() string {
	return "liquidity"
}

func (m *Module) HandleBlock(block *coretypes.ResultBlock, results *coretypes.ResultBlockResults, txs []*types.Tx, vals *coretypes.ResultValidators) error {
	err := m.executePoolBatches(block.Block.Height, results.EndBlockEvents, block.Block.Time)
	if err != nil {
		return fmt.Errorf("error while executing pool batches: %s", err)
	}
	return nil
}

