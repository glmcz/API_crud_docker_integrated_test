#!/bin/bash
set -e

while ! nc -z postgres 5432; do
    echo "Waiting for the database to become available..."
    sleep 1
done

./app