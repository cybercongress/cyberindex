CREATE MATERIALIZED VIEW honest_pre_commits AS (
    SELECT
        validator_address,
        count(*) as pre_commits
    FROM (
        SELECT
            RANK () OVER (
                PARTITION BY validator_address, height
                ORDER BY height
                ) pre_commit_rank,
                validator_address
        FROM pre_commit
    ) t
    WHERE pre_commit_rank = 1
    GROUP BY validator_address
);


CREATE OR REPLACE VIEW pre_commits_total AS (
    SELECT
        validator.consensus_address,
        t.pre_commits
    FROM (
        SELECT
            honest_pre_commits.validator_address,
            honest_pre_commits.pre_commits * 100.0 / (max_height - min_height) AS pre_commits
        FROM (
            SELECT
                validator_address,
                max(height) AS max_height,
                min(height) AS min_height
            FROM pre_commit
            GROUP BY validator_address
        ) t
        LEFT JOIN honest_pre_commits ON honest_pre_commits.validator_address = t.validator_address
    ) t
    LEFT JOIN validator ON validator.consensus_address = t.validator_address
    ORDER BY pre_commits DESC
);