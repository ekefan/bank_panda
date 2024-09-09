#!/bin/sh

# set -e to ensure the script retuns immediately if a command returns a non-zero status
set -e

echo "run db migration"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the up"

# exec "$@" takes all parameters passed to the script and run it

exec "$@"
