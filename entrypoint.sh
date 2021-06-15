#! /bin/bash

if [ ! -d "/root/.cyberindex/" ]
then
    export $(cat .env)
    mkdir /root/.cyberindex
    echo -e "[cosmos]\n  prefix = \"cyber\"\n  modules = [\n    \"modules\",\n    \"auth\",\n    \"bank\",\n    \"messages\",\n    \"graph\"\n  ]\n\n[database]\n  host = \"$POSTGRES_DB_HOST\"\n  name = \"$POSTGRES_DB_NAME\"\n  user = \"$POSTGRES_USER_NAME\"\n  password = \"$POSTGRES_DB_PASSWORD\"\n  port = $POSTGRES_DB_PORT\n  schema = \"public\"\n  ssl_mode = \"$POSTGRES_SSL_MODE\"\n\n[grpc]\n  address = \"$GRPC_URL\"\n  insecure = true\n\n[logging]\n  format = \"json\"\n  level = \"debug\"\n\n[parsing]\n  fast_sync = false\n  listen_new_blocks = true\n  parse_old_blocks = true\n  start_height = 0\n  workers = $JUNO_WORKERS\n  parse_genesis = true\n\n[rpc]\n  address = \"$RPC_URL\"" >> /root/.cyberindex/config.toml
fi

if [ ! -f "/root/.cyberindex/config.toml" ]
then
    export $(cat .env)
    echo -e "[cosmos]\n  prefix = \"cyber\"\n  modules = [\n    \"modules\",\n    \"auth\",\n    \"bank\",\n    \"messages\",\n    \"graph\"\n  ]\n\n[database]\n  host = \"$POSTGRES_DB_HOST\"\n  name = \"$POSTGRES_DB_NAME\"\n  user = \"$POSTGRES_USER_NAME\"\n  password = \"$POSTGRES_DB_PASSWORD\"\n  port = $POSTGRES_DB_PORT\n  schema = \"public\"\n  ssl_mode = \"$POSTGRES_SSL_MODE\"\n\n[grpc]\n  address = \"$GRPC_URL\"\n  insecure = true\n\n[logging]\n  format = \"json\"\n  level = \"debug\"\n\n[parsing]\n  fast_sync = false\n  listen_new_blocks = true\n  parse_old_blocks = true\n  start_height = 0\n  workers = $JUNO_WORKERS\n  parse_genesis = true\n\n[rpc]\n  address = \"$RPC_URL\"" >> /root/.cyberindex/config.toml
fi

exec "$@"
