#!/bin/bash

set -euo pipefail

# wait for postgres to be up
RETRIES=10
until PGPASSWORD=$PGPASS psql --host=$PGHOST --user=$PGUSER -d $PGDB --command "\q" >/dev/null 2>&1 || [ $RETRIES -eq 0 ]; do
    echo >&2 "Waiting for the PostgreSQL server to start. $((RETRIES)) remaining attempts..."
    RETRIES=$((RETRIES -= 1))
    sleep 2
done

# start go server
./app
