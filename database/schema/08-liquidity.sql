CREATE TABLE pools
(
    pool_id                     BIGINT NOT NULL PRIMARY KEY,
    pool_name                   TEXT   NOT NULL,
    address                     TEXT   NOT NULL,
    a_denom                     TEXT   NOT NULL,
    b_denom                     TEXT   NOT NULL,
    pool_denom                  TEXT   NOT NULL
);

CREATE TABLE swaps
(
    id                          SERIAL PRIMARY KEY,
    pool_id                     BIGINT NOT NULL REFERENCES pools (pool_id),
    address                     TEXT   NOT NULL,
    order_price                 TEXT   NOT NULL,
    swap_price                  TEXT   NOT NULL,
    exchanged_offer_coin        COIN   NOT NULL,
    exchanged_demand_coin       COIN   NOT NULL,
    exchanged_offer_coin_fee    COIN   NOT NULL,
    exchanged_demand_coin_fee   COIN   NOT NULL,
    height                      BIGINT NOT NULL REFERENCES block (height),
    timestamp                   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE failed_swaps
(
    id                          SERIAL PRIMARY KEY,
    pool_id                     BIGINT NOT NULL REFERENCES pools (pool_id),
    address                     TEXT   NOT NULL,
    order_price                 TEXT   NOT NULL,
    exchanged_offer_coin        COIN   NOT NULL,
    remaining_offer_coin        COIN   NOT NULL,
    height                      BIGINT NOT NULL REFERENCES block (height),
    timestamp                   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE pools_volumes
(
    pool_id                     BIGINT NOT NULL REFERENCES pools (pool_id),
    volume_a                    BIGINT NOT NULL,
    volume_b                    BIGINT NOT NULL,
    fee_a                       BIGINT NOT NULL,
    fee_b                       BIGINT NOT NULL,
    timestamp                   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE pools_liquidity
(
    pool_id                     BIGINT NOT NULL REFERENCES pools (pool_id),
    liquidity_a                 BIGINT NOT NULL,
    liquidity_b                 BIGINT NOT NULL,
    timestamp                   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE pools_rates
(
    pool_id                     BIGINT NOT NULL REFERENCES pools (pool_id),
    rate                        TEXT   NOT NULL,
    timestamp                   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);