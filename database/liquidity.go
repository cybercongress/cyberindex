package database

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bddbtypes "github.com/forbole/bdjuno/v3/database/types"
	"time"
)

func (db *CyberDb) SavePool(
	poolID uint64,
	address string,
	name string,
	deposit_a sdk.Coin,
	deposit_b sdk.Coin,
	pool_denom string,
) error {
	stmt := `
INSERT INTO pools (pool_id, pool_name, address, deposit_a, deposit_b, pool_denom) 
VALUES ($1, $2, $3, $4, $5, $6)`
	deposit_a_coin := bddbtypes.DbCoin{Amount: deposit_a.Amount.String(), Denom: deposit_a.Denom}
	deposit_a_value, err := deposit_a_coin.Value()
	if err != nil {
		return fmt.Errorf("error while converting coin to dbcoin: %s", err)
	}

	deposit_b_coin := bddbtypes.DbCoin{Amount: deposit_b.Amount.String(), Denom: deposit_b.Denom}
	deposit_b_value, err := deposit_b_coin.Value()
	if err != nil {
		return fmt.Errorf("error while converting coin to dbcoin: %s", err)
	}

	_, err = db.Sql.Exec(
		stmt,
		poolID,
		name,
		address,
		deposit_a_value,
		deposit_b_value,
		pool_denom,
	)
	return err
}

func (db *CyberDb) SaveSwap(
	address string,
	poolID uint64,
	swapPrice sdk.Dec,
	exchangedOfferCoin sdk.Coin,
	exchangedDemandCoin sdk.Coin,
	exchangedOfferCoinFee sdk.Coin,
	exchangedDemandCoinFee sdk.Coin,
	height int64,
) error {
	stmt := `
INSERT INTO swaps (pool_id, address, swap_price, exchanged_offer_coin, exchanged_demand_coin, exchanged_offer_coin_fee, exchanged_demand_coin_fee, height) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	exchangedOfferCoinDb := bddbtypes.DbCoin{Amount: exchangedOfferCoin.Amount.String(), Denom: exchangedOfferCoin.Denom}
	exchangedOfferCoinValue, err := exchangedOfferCoinDb.Value()
	if err != nil {
		return fmt.Errorf("error while converting coin to dbcoin: %s", err)
	}
	exchangedDemandCoinDb := bddbtypes.DbCoin{Amount: exchangedDemandCoin.Amount.String(), Denom: exchangedDemandCoin.Denom}
	exchangedDemandCoinValue, err := exchangedDemandCoinDb.Value()
	if err != nil {
		return fmt.Errorf("error while converting coin to dbcoin: %s", err)
	}
	exchangedOfferCoinFeeDb := bddbtypes.DbCoin{Amount: exchangedOfferCoinFee.Amount.String(), Denom: exchangedOfferCoinFee.Denom}
	exchangedOfferCoinFeeValue, err := exchangedOfferCoinFeeDb.Value()
	if err != nil {
		return fmt.Errorf("error while converting coin to dbcoin: %s", err)
	}
	exchangedDemandCoinFeeDb := bddbtypes.DbCoin{Amount: exchangedDemandCoinFee.Amount.String(), Denom: exchangedDemandCoinFee.Denom}
	exchangedDemandCoinFeeValue, err := exchangedDemandCoinFeeDb.Value()
	if err != nil {
		return fmt.Errorf("error while converting coin to dbcoin: %s", err)
	}
	_, err = db.Sql.Exec(
		stmt,
		poolID,
		address,
		swapPrice.String(),
		exchangedOfferCoinValue,
		exchangedDemandCoinValue,
		exchangedOfferCoinFeeValue,
		exchangedDemandCoinFeeValue,
		height,
	)
	return err
}

func (db *CyberDb) SaveLiquidity(
	poolID uint64,
	liquidityA sdk.Coin,
	liquidityB sdk.Coin,
	timestamp time.Time,
) error {
	stmt := `
INSERT INTO pools_liquidity (pool_id, liquidity_a, liquidity_b, timestamp) 
VALUES ($1, $2, $3, $4)`
	_, err := db.Sql.Exec(
		stmt,
		poolID,
		liquidityA.Amount.String(),
		liquidityB.Amount.String(),
		timestamp,
	)
	return err
}

func (db *CyberDb) SaveVolume(
	poolID uint64,
	volumeA sdk.Coin,
	volumeB sdk.Coin,
	feeA sdk.Coin,
	feeB sdk.Coin,
	timestamp time.Time,
) error {
	stmt := `
INSERT INTO pools_volumes (pool_id, volume_a, volume_b, fee_a, fee_b, timestamp) 
VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Sql.Exec(
		stmt,
		poolID,
		volumeA.Amount.String(),
		volumeB.Amount.String(),
		feeA.Amount.String(),
		feeB.Amount.String(),
		timestamp,
	)
	return err
}

func (db *CyberDb) SaveRate(
	poolID uint64,
	rate sdk.Dec,
	timestamp time.Time,
) error {
	stmt := `
INSERT INTO pools_rates (pool_id, rate, timestamp) 
VALUES ($1, $2, $3)`
	_, err := db.Sql.Exec(
		stmt,
		poolID,
		rate.String(),
		timestamp,
	)
	return err
}

// GetAccountBalance returns the balance of the user having the given address
func (db *CyberDb) GetPoolInfo(poolID uint64) (PoolRowNative, error) {
	stmt := `SELECT * FROM pools WHERE pool_id = $1`

	var rows []PoolRow
	err := db.Sqlx.Select(&rows, stmt, poolID)
	if err != nil {
		return PoolRowNative{}, err
	}

	if len(rows) == 0 {
		return PoolRowNative{}, nil
	}

	return PoolRowNative{
		rows[0].PoolId,
		rows[0].PoolName,
		rows[0].Address,
		rows[0].DepositA.ToCoin(),
		rows[0].DepositB.ToCoin(),
		rows[0].PoolDenom,
	}, nil
}

type PoolRow struct {
	PoolId    int64            `db:"pool_id"`
	PoolName  string           `db:"pool_name"`
	Address   string           `db:"address"`
	DepositA  bddbtypes.DbCoin `db:"deposit_a"`
	DepositB  bddbtypes.DbCoin `db:"deposit_b"`
	PoolDenom string           `db:"pool_denom"`
}

type PoolRowNative struct {
	PoolId    int64    `db:"pool_id"`
	PoolName  string   `db:"pool_name"`
	Address   string   `db:"address"`
	DepositA  sdk.Coin `db:"deposit_a"`
	DepositB  sdk.Coin `db:"deposit_b"`
	PoolDenom string   `db:"pool_denom"`
}
