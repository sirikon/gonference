#!/usr/bin/env bash

set -e
cd "$(dirname "${BASH_SOURCE[0]}")/../.."

(cd ./docker/devenv && docker-compose \
  -p gonference-devenv \
  down)
