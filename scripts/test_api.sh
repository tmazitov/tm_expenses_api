#!/bin/bash

BASE_URL="http://localhost:8080"

echo "--- 1. Create expense (valid) ---"
curl -s -w "\nStatus: %{http_code}\n" -X POST "$BASE_URL/expense" \
    -H "Content-Type: application/json" \
    -d '{"name": "Coffee"}' | cat

echo ""
echo "--- 2. Create expense (empty name) ---"
curl -s -w "\nStatus: %{http_code}\n" -X POST "$BASE_URL/expense" \
    -H "Content-Type: application/json" \
    -d '{"name": ""}' | cat

echo ""
echo "--- 3. Create expense (missing body) ---"
curl -s -w "\nStatus: %{http_code}\n" -X POST "$BASE_URL/expense" \
    -H "Content-Type: application/json" | cat

echo ""
echo "--- 4. Create expense (name is too long) ---"
curl -s -w "\nStatus: %{http_code}\n" -X POST "$BASE_URL/expense" \
    -H "Content-Type: application/json" \
    -d '{"name": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}' | cat


echo ""
echo "--- 5. Create expense with long name (valid) ---"
curl -s -w "\nStatus: %{http_code}\n" -X POST "$BASE_URL/expense" \
    -H "Content-Type: application/json" \
    -d '{"name": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}' | cat
