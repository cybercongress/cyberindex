CREATE TABLE old_precommits
(
    consensus_pubkey  TEXT NOT NULL UNIQUE /* Validator consensus public key */
    consensus_address TEXT NOT NULL PRIMARY KEY, /* Validator consensus address */
    precommits NUMERIC
);

CREATE VIEW _transaction as (
    SELECT *,
    involved_accounts_addresses[1] subject1,
    involved_accounts_addresses[2] subject2
    FROM (
        SELECT *
        FROM transaction
        LEFT JOIN (
            SELECT * 
            FROM message
        ) as msg ON transaction.hash = msg.transaction_hash
    ) _tx
);

CREATE VIEW pre_commits_view as (
    SELECT 
        consensus_pubkey,
        consensus_address,
        SUM(precommits) AS precommits
    FROM (
        SELECT 
            validator.consensus_pubkey,
            validator.consensus_address,
            pr_cnt.precommits
        FROM (
            SELECT
                validator_address,
                count(*) AS precommits
            FROM 
                pre_commit
            GROUP BY validator_address
        ) as pr_cnt
        LEFT JOIN validator ON (
            validator.consensus_address = pr_cnt.validator_address
        )
        UNION
        SELECT * FROM old_precommits
    ) t
    GROUP BY consensus_pubkey, consensus_address
);

CREATE VIEW pre_commits_rewards_view as (
    SELECT
        t.consensus_pubkey,
        t.precommits,
        t.sum_precommits,
        max(height) + 200000 as max_block,
        ((((max(height) + 200000 + 257620) / 50000) + 1) * (cast(precommits as decimal) / sum_precommits)) * 5000000000 as pre_commit_rewards
    FROM (
        SELECT
            consensus_pubkey,
            precommits,
            b.sum_precommits
        FROM
            pre_commits_view
        CROSS JOIN
        (
            SELECT sum(precommits) sum_precommits FROM pre_commits_view
        ) as b
    ) as t, block
    GROUP BY consensus_pubkey, precommits, sum_precommits
    ORDER BY precommits DESC
);

CREATE VIEW _uptime_temp AS (
    SELECT 
        t.validator_address,
        count(t.validator_address) AS pre_commits
    FROM (
        SELECT 
            temp.validator_address,
            temp.height,
            temp.rank
            FROM ( 
                SELECT 
                    pre_commit.validator_address,
                    pre_commit.height,
                    rank() OVER (PARTITION BY pre_commit.validator_address, pre_commit.height ORDER BY pre_commit.height DESC) AS rank
                FROM pre_commit
            ) temp
            WHERE (
                temp.height >= (SELECT max(block.height) AS max
                   FROM block) - 50000 AND temp.rank = 1)
    ) t
  GROUP BY t.validator_address
);

CREATE VIEW uptime AS (
    SELECT
        consensus_address,
        consensus_pubkey,
        _uptime_temp.pre_commits,
        cast(_uptime_temp.pre_commits as decimal) / 50000 AS uptime
    FROM
        validator
    LEFT JOIN _uptime_temp ON validator.consensus_address = _uptime_temp.validator_address
);