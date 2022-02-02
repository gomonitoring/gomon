#!/usr/bin/env bash

echo "RUNNING APPLICATION ..."
if [ $FUNC != "monitoring_worker" ]
then
    echo "WAITING FOR DATABASE ..."
    while ! nc -z ${DB_HOST} ${DB_PORT}; do sleep 2; done
    echo "CONNECTED DATABASE ..."   
fi
go run ./main.go $FUNC