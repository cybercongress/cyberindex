#! /bin/bash

if [ ! -d "/root/.cyberindex/" ]
then
    export $(cat .env)
    mkdir /root/.cyberindex
    echo -e "chain:\n    bech32_prefix: $CHAIN_PREFIX\n    modules:\n        - modules\n        - auth\n        - bank\n        - messages\n        - graph\n        - grid\n        - wasm\n        - liquidity\n        - resources\nnode:\n    type: \"remote\"\n    config:\n        rpc:\n            client_name: \"cyberindex\"\n            max_connections: 10\n            address: \"$RPC_URL\"\n        grpc:\n            insecure: true\n            address: \"$GRPC_URL\"\ndatabase:\n    host: $POSTGRES_DB_HOST\n    max_idle_connections: 1\n    max_open_connections: 1\n    name: $POSTGRES_DB_NAME\n    password: \"$POSTGRES_DB_PASSWORD\"\n    port: $POSTGRES_DB_PORT\n    schema: public\n    partition_size: 1000000\n    partition_batch: 1000\n    user: $POSTGRES_USER_NAME\nlogging:\n    format: text\n    level: debug\n\nparsing:\n    listen_new_blocks: true\n    parse_genesis: true\n    parse_old_blocks: true\n    start_height: $START_HEIGHT\n    workers: $JUNO_WORKERS" >> /root/.cyberindex/config.yaml
fi

if [ -f "/root/.cyberindex/config.yaml" ]
then
    export $(cat .env)
    rm /root/.cyberindex/config.yaml
    echo -e "chain:\n    bech32_prefix: $CHAIN_PREFIX\n    modules:\n        - modules\n        - auth\n        - bank\n        - messages\n        - graph\n        - grid\n        - wasm\n        - liquidity\n        - resources\nnode:\n    type: \"remote\"\n    config:\n        rpc:\n            client_name: \"cyberindex\"\n            max_connections: 10\n            address: \"$RPC_URL\"\n        grpc:\n            insecure: true\n            address: \"$GRPC_URL\"\ndatabase:\n    host: $POSTGRES_DB_HOST\n    max_idle_connections: 1\n    max_open_connections: 1\n    name: $POSTGRES_DB_NAME\n    password: \"$POSTGRES_DB_PASSWORD\"\n    port: $POSTGRES_DB_PORT\n    schema: public\n    partition_size: 1000000\n    partition_batch: 1000\n    user: $POSTGRES_USER_NAME\nlogging:\n    format: text\n    level: debug\n\nparsing:\n    listen_new_blocks: true\n    parse_genesis: true\n    parse_old_blocks: true\n    start_height: $START_HEIGHT\n    workers: $JUNO_WORKERS" > /root/.cyberindex/config.yaml
fi

if [ ! -f "/root/.cyberindex/config.yaml" ]
then
    export $(cat .env)
    echo -e "chain:\n    bech32_prefix: $CHAIN_PREFIX\n    modules:\n        - modules\n        - auth\n        - bank\n        - messages\n        - graph\n        - grid\n        - wasm\n        - liquidity\n        - resources\nnode:\n    type: \"remote\"\n    config:\n        rpc:\n            client_name: \"cyberindex\"\n            max_connections: 10\n            address: \"$RPC_URL\"\n        grpc:\n            insecure: true\n            address: \"$GRPC_URL\"\ndatabase:\n    host: $POSTGRES_DB_HOST\n    max_idle_connections: 1\n    max_open_connections: 1\n    name: $POSTGRES_DB_NAME\n    password: \"$POSTGRES_DB_PASSWORD\"\n    port: $POSTGRES_DB_PORT\n    schema: public\n    partition_size: 1000000\n    partition_batch: 1000\n    user: $POSTGRES_USER_NAME\nlogging:\n    format: text\n    level: debug\n\nparsing:\n    listen_new_blocks: true\n    parse_genesis: true\n    parse_old_blocks: true\n    start_height: $START_HEIGHT\n    workers: $JUNO_WORKERS" >> /root/.cyberindex/config.yaml
fi

exec "$@"
