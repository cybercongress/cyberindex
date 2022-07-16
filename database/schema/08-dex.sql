CREATE TABLE swaps
(
    pool_id                     BIGINT NOT NULL PRIMARY KEY,
    address                     TEXT   NOT NULL,
    swap_price                  TEXT   NOT NULL,
    exchanged_offer_coin        COIN   NOT NULL,
    exchanged_demand_coin       COIN   NOT NULL,
    exchanged_offer_coin_fee    COIN   NOT NULL,
    exchanged_demand_coin_fee   COIN   NOT NULL,
    height                      BIGINT NOT NULL REFERENCES block (height)
);

CREATE TABLE pools
(
    pool_id                     BIGINT NOT NULL PRIMARY KEY,
    pool_name                   TEXT   NOT NULL,
    address                     TEXT   NOT NULL,
    deposit_a                   COIN   NOT NULL,
    deposit_b                   COIN   NOT NULL,
    pool_denom                  TEXT   NOT NULL
);

CREATE TABLE pools_volumes
(
    pool_id                     BIGINT NOT NULL,
    volume_a                    BIGINT NOT NULL,
    volume_b                    BIGINT NOT NULL,
    fee_a                       BIGINT NOT NULL,
    fee_b                       BIGINT NOT NULL,
    timestamp                   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE pools_liquidity
(
    pool_id                     BIGINT NOT NULL,
    liquidity_a                 BIGINT NOT NULL,
    liquidity_b                 BIGINT NOT NULL,
    timestamp                   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE pools_rates
(
    pool_id                     BIGINT NOT NULL,
    rate                        TEXT   NOT NULL,
    timestamp                   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);