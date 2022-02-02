#!/usr/bin/env bash

echo "RUNNING APPLICATION ..."

echo "WAITING FOR DATABASE ..."
while ! nc -z ${DB_HOST} ${DB_PORT}; do sleep 2; done  # TODO: this pars should probably be removed or 
                                                       #       monitoring_worker should have a different
echo "CONNECTED DATABASE ..."                          #       entry point script
go run ./main.go $FUNC