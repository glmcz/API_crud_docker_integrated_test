#!/bin/bash
set -e

while ! nc -z postgres; do
    echo "Waiting for the database to become available..."
    sleep 1
done

./app