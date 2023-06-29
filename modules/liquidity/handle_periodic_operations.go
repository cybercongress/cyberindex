package liquidity

import (
	"github.com/cybercongress/cyberindex/database"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "liquidity").Msg("setting up periodic tasks")

	return nil
}

func (m *Module) RunAdditionalOperations() error {
	log.Debug().Str("module", "liquidity").Msg("running additional operation")

	err := m.syncPoolsInfo()
	if err != nil {
		log.Debug().Str("module", "liquidity").Err(err)
		panic(err)
	}

	return nil
}

func (m *Module) syncPoolsInfo() error {
	stmt := `SELECT * FROM pools`
	var rows []database.PoolRow
	err := m.db.Sqlx.Select(&rows, stmt)
	if err != nil {
		return err
	}
	pools, err := m.keeper.GetAllPools(0)
	if err != nil {
		return err
	}
	if len(rows) != len(pools) {
		for _, pool := range pools {
			err := m.db.SavePool(
				pool.Id,
				pool.ReserveAccountAddress,
				pool.Name(),
				pool.ReserveCoinDenoms[0],
				pool.ReserveCoinDenoms[1],
				pool.PoolCoinDenom,
			)
			if err != nil {
				log.Debug().Str("module", "liquidity").Err(err)
			}
		}
	}
	return nil
}
