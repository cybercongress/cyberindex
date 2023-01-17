#! /bin/bash

echo "Please put use your onw passwords in .env file. Do you want to proceed with current .env (y / n)?"
read -r  CONT
echo
if [ "$CONT" = "y" ]
then
    # temporeraly import variables
    export $(cat .env)

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
    docker exec -ti cyberindex_postgres psql -f /root/schema/08-dex.sql -d $POSTGRES_DB_NAME -U $POSTGRES_USER_NAME

    docker run -d --name cyberindex --network="host" -v $HOME/.cyberindex:/root/.cyberindex cyberindex:latest

else
    echo "Done."
fi