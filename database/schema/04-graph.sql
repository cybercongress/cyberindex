CREATE TABLE cyberlinks
(
    id                  SERIAL PRIMARY KEY,
    particle_from       VARCHAR(256)                NOT NULL,
    particle_to         VARCHAR(256)                NOT NULL,
    neuron              TEXT                        NOT NULL REFERENCES account (address),
    timestamp           TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    height              BIGINT                      NOT NULL REFERENCES block (height),
    transaction_hash    TEXT                        NOT NULL REFERENCES transaction (hash)
);

CREATE TABLE particles
(
    id                  SERIAL PRIMARY KEY,
    particle            VARCHAR(256)                NOT NULL UNIQUE,
    neuron              TEXT                        NOT NULL REFERENCES account (address),
    timestamp           TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    height              BIGINT                      NOT NULL REFERENCES block (height),
    transaction_hash    TEXT                        NOT NULL REFERENCES transaction (hash)
)