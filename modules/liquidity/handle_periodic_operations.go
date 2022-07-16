package liquidity

import (
	"fmt"
	"github.com/forbole/bdjuno/v2/modules/utils"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "liquidity").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(10).Seconds().Do(func() {
		utils.WatchMethod(m.updateRates)
	}); err != nil {
		return fmt.Errorf("error while setting up liquidity periodic operation: %s", err)
	}

	return nil
}

// updateSupply updates the supply of all the tokens
func (m *Module) updateRates() error {
	log.Debug().Str("module", "liqudity").Str("operation", "update rates").
		Msg("updating rates!!")

	//block, err := m.db.GetLastBlock()
	//if err != nil {
	//	return fmt.Errorf("error while getting last block: %s", err)
	//}
	//
	//supply, err := m.keeper.GetSupply(block.Height)
	//if err != nil {
	//	return err
	//}
	//
	//return m.db.SaveSupply(supply, block.Height)
	return nil
}
