#!/bin/bash
set -e

MIGRATIONS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/../db/migrations"

read -p "Enter migration name: " MIGRATION_NAME
if [[ -z "$MIGRATION_NAME" ]]; then
    echo "Error: migration name cannot be empty"
    exit 1
fi

mkdir -p "$MIGRATIONS_DIR"
goose -dir "$MIGRATIONS_DIR" create "$MIGRATION_NAME" sql

echo "Migration created in $MIGRATIONS_DIR"