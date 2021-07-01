#! /bin/bash

# Stops and rm all cyberindex docker containers

docker stop cyberindex
docker stop cyberindex_postgres
docker stop cyberindex_hasura
docker rm cyberindex
docker rm cyberindex_postgres
docker rm cyberindex_hasura
docker network rm cyberindex_cyberindex-net
docker rmi cyberindex:latest
# rm -rf $HOME/.cyberindex # remove home directory with db 