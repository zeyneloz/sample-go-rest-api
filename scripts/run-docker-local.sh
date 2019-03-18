#!/usr/bin/env bash

DOCKER_COMPOSE_LOCAL_FILE=docker-compose-local.yml

echo
echo "---------------------------"
echo "Stopping containers..."
echo
docker-compose --file $DOCKER_COMPOSE_LOCAL_FILE stop

echo
echo "---------------------------"
echo "removing containers"
echo
docker-compose --file $DOCKER_COMPOSE_LOCAL_FILE rm -f

echo
echo "---------------------------"
echo "Building images"
echo
docker-compose --file $DOCKER_COMPOSE_LOCAL_FILE build

echo
echo "---------------------------"
echo "Starting containers"
echo
docker-compose --file $DOCKER_COMPOSE_LOCAL_FILE up -d