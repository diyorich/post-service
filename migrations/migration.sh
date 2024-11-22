#!/bin/bash

export MIGRATION_DIR=./migrations
export DB_HOST="post-db"
export DB_PORT="5432"
export DB_NAME="post"
export DB_USER="post"
export DB_PASSWORD="post"
export DB_SSL=disable

export PG_DSN="host=${DB_HOST} port=${DB_PORT} dbname=${DB_NAME} user=${DB_USER} password=${DB_PASSWORD} sslmode=${DB_SSL}"

sleep 2 && goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" up -v
