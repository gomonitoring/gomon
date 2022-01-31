#!/usr/bin/env bash

echo "RUNNING APPLICATION ..."

echo "WAITING FOR DATABASE ..."
while ! nc -z ${DB_HOST} ${DB_PORT}; do sleep 2; done
echo "CONNECTED DATABASE ..."
./main
