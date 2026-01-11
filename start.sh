#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/internal/db/migration -database "$DATABASE_URL" -verbose up

echo "start the app"
exec "$@"