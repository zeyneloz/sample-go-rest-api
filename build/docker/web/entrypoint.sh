#!/bin/bash
set -e

while !</dev/tcp/$POSTGRES_HOST/$POSTGRES_PORT; do
    sleep 1;
done;
./main