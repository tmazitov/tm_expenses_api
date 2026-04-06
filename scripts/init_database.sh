#!/bin/bash
set -e

echo "=== PostgreSQL Database Setup ==="
echo ""

# Username
read -p "Enter username: " DB_USER
if [[ -z "$DB_USER" ]]; then
    echo "Error: username cannot be empty"
    exit 1
fi

# Password
read -s -p "Enter password for '$DB_USER': " DB_PASS
echo ""
if [[ -z "$DB_PASS" ]]; then
    echo "Error: password cannot be empty"
    exit 1
fi

read -s -p "Confirm password: " DB_PASS_CONFIRM
echo ""
if [[ "$DB_PASS" != "$DB_PASS_CONFIRM" ]]; then
    echo "Error: passwords do not match"
    exit 1
fi

# Database name
read -p "Enter database name: " DB_NAME
if [[ -z "$DB_NAME" ]]; then
    echo "Error: database name cannot be empty"
    exit 1
fi

# Schema name
read -p "Enter schema name [core]: " DB_SCHEMA
DB_SCHEMA="${DB_SCHEMA:-core}"

# Admin user (who runs the script)
read -p "Enter PostgreSQL admin username [postgres]: " POSTGRES_USER
POSTGRES_USER="${POSTGRES_USER:-postgres}"

# Confirmation
echo ""
echo "--- Summary ---"
echo "  Admin user : $POSTGRES_USER"
echo "  New user   : $DB_USER"
echo "  Database   : $DB_NAME"
echo "  Schema     : $DB_SCHEMA"
echo "---------------"
read -p "Proceed? [y/N]: " CONFIRM
if [[ "$CONFIRM" != "y" && "$CONFIRM" != "Y" ]]; then
    echo "Aborted."
    exit 0
fi

echo ""
echo "Running init script..."

psql -h localhost -U "$POSTGRES_USER" \
    -v db_name="$DB_NAME" \
    -v db_user="$DB_USER" \
    -v db_pass="$DB_PASS" \
    -v db_schema="$DB_SCHEMA" \
    -f ./init_database.sql

echo "Database '$DB_NAME' with user '$DB_USER' initialized successfully."
