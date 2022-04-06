#! /bin/bash

echo "Please put use your onw passwords in .env file. Do you want to proceed with current .env (y / n)?"
read -r  CONT
echo
if [ "$CONT" = "y" ]
then
    # temporeraly import variables
    export $(cat .env)

    # get static gift and genesis tables

    if [ ! -f "database/schema/genesis.csv" ]
    then
        wget -O database/schema/genesis.csv https://gateway.ipfs.cybernode.ai/ipfs/QmWxvLnFZDJUrjTjNDt4BfanzncdbzTMfSQmkNAACQ8ZaF
    fi

    if [ ! -f "database/schema/cyber_gift.csv" ]
    then
        wget -O database/schema/cyber_gift.csv https://gateway.ipfs.cybernode.ai/ipfs/QmQC1WRfAfp6zDdbaVrYTC4qmJr1uMCo9LmDByePZ9TFEy
    fi

    # build cyberindexer and run it in container
    docker build -t cyberindex:latest .

    # run postgres and hasura in containers
    docker-compose up -d postgres 
    sleep 10

    docker-compose up -d graphql-engine 
    sleep 10

    # init database with basic tables
    docker exec -ti cyberindex_postgres psql -f /root/schema/00-cosmos.sql -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME
    docker exec -ti cyberindex_postgres psql -f /root/schema/01-auth.sql -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME
    docker exec -ti cyberindex_postgres psql -f /root/schema/02-bank.sql -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME
    docker exec -ti cyberindex_postgres psql -f /root/schema/03-modules.sql -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME
    docker exec -ti cyberindex_postgres psql -f /root/schema/04-graph.sql -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME
    docker exec -ti cyberindex_postgres psql -f /root/schema/05-grid.sql -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME
    docker exec -ti cyberindex_postgres psql -f /root/schema/06-resources.sql -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME
    docker exec -ti cyberindex_postgres psql -f /root/schema/07-wasm.sql -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME

    # init additional views and tables
    docker exec -ti cyberindex_postgres psql -f /root/schema/views.sql -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME
    docker exec -ti cyberindex_postgres psql -f /root/schema/delegation_strategy.sql -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME

    # copy genesis and cyber_gift from csv to table
    docker exec -ti cyberindex_postgres psql -c "\copy genesis FROM /root/schema/genesis.csv with csv HEADER" -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME
    docker exec -ti cyberindex_postgres psql -c "\copy cyber_gift FROM /root/schema/cyber_gift.csv with csv HEADER" -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME

    docker run -d --name cyberindex --network="host" -v $HOME/.cyberindex:/root/.cyberindex cyberindex:latest

    # add cronjob to refresh tables for stats

    croncmd="docker exec -t cyberindex_postgres psql -c \"REFRESH MATERIALIZED VIEW CONCURRENTLY txs_ranked\" -d cyberindex -U cyber"
    cronjob="*/5 * * * * $croncmd"
    ( crontab -l | grep -v -F "$croncmd" ; echo "$cronjob" ) | crontab -

    croncmd="docker exec -t cyberindex_postgres psql -c \"REFRESH MATERIALIZED VIEW CONCURRENTLY honest_pre_commits\" -d cyberindex -U cyber"
    cronjob="*/30 * * * * $croncmd"
    ( crontab -l | grep -v -F "$croncmd" ; echo "$cronjob" ) | crontab -

else
    echo "Done."
fi