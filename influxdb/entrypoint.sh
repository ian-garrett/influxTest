#!/usr/bin/env sh

influx -host=localhost -port=8086 \
-execute="CREATE USER ian " \
         "WITH PASSWORD garrett " \
         "WITH ALL PRIVILEGES"

