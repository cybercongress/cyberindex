CREATE TABLE cyberlinks
(
    id                  SERIAL PRIMARY KEY,
    object_from         VARCHAR(256)                NOT NULL,
    object_to           VARCHAR(256)                NOT NULL,
    subject             TEXT                        NOT NULL REFERENCES account (address),
    timestamp           TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    height              BIGINT                      NOT NULL REFERENCES block (height),
    transaction_hash    TEXT                        NOT NULL REFERENCES transaction (hash)
);

CREATE TABLE particles
(
    id                  SERIAL PRIMARY KEY,
    object              VARCHAR(256)                NOT NULL UNIQUE,
    subject             TEXT                        NOT NULL REFERENCES account (address),
    timestamp           TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    height              BIGINT                      NOT NULL REFERENCES block (height),
    transaction_hash    TEXT                        NOT NULL REFERENCES transaction (hash)
)