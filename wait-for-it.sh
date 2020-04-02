#!/bin/sh
# wait-for-postgres.sh

set -e

shift
cmd="$@"
until PGPASSWORD=$DB_PWD psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -c '\q'; do
    >&2 echo "Trying to connect to " $DB_HOST:$DB_PORT with user $DB_USER and pwd $DB_PWD
    >&2 echo "Postgres is unavailable - sleeping"
    sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd
