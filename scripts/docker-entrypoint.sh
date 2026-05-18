#!/bin/sh
# set -e

# DB_HOST="${DB_HOST:-postgres}"
# DB_PORT="${DB_PORT:-5432}"
# TIMEOUT=300

# echo "Waiting for PostgreSQL at ${DB_HOST}:${DB_PORT}..."
# i=0
# until nc -z "$DB_HOST" "$DB_PORT" 2>/dev/null; do
#   i=$((i + 1))
#   if [ "$i" -ge "$TIMEOUT" ]; then
#     echo "PostgreSQL did not become ready in ${TIMEOUT}s, exiting."
#     exit 1
#   fi
#   sleep 1
# done
# echo "PostgreSQL is ready."

exec /bin/server
