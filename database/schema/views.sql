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
);

CREATE VIEW pre_commits_rewards_view as (
    SELECT
        t.consensus_pubkey,
        t.precommits,
        t.sum_precommits,
        max(height) as max_block,
        (((max(height) / 50000) + 1) * (cast(precommits as decimal) / sum_precommits)) * 5000000000 as pre_commit_rewards
    FROM (
        SELECT
            consensus_pubkey,
            precommits,
            b.sum_precommits
        FROM
            pre_commits_view
        CROSS JOIN
        (
            SELECT count('timestamp') sum_precommits FROM pre_commit
        ) as b
    ) as t, block
    GROUP BY consensus_pubkey, precommits, sum_precommits
    ORDER BY precommits DESC
)