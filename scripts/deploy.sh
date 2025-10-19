#!/bin/bash

set -e

cd projects
docker compose pull sakila
docker compose up -d --no-deps sakila
docker image prune
