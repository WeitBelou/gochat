#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

docker-compose rm --force
docker-compose up --build