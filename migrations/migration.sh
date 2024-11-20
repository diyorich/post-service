#!/bin/bash

export MIGRATION_DIR=./migrations
export DB_HOST="service-db"
export DB_PORT="5432"
export DB_NAME="service"
export DB_USER="service"
export DB_PASSWORD="a8B3kLm2P"
export DB_SSL=disable

export PG_DSN="host=${DB_HOST} port=${DB_PORT} dbname=${DB_NAME} user=${DB_USER} password=${DB_PASSWORD} sslmode=${DB_SSL}"

sleep 2 && goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" up -v