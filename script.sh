#!/usr/bin/env bash

curl -sLOk $1/archive/master.zip
unzip master.zip

IFS='/' read -ra ADDR <<< "$1"
NAME=${ADDR[@]: -1}

docker build -t $NAME ./$NAME-master --no-cache

# TODO: Need somewhere to store the docker container
