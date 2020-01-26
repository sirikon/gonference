#!/usr/bin/env bash

set -e
cd "$(dirname "${BASH_SOURCE[0]}")/.."

(
  export GOPROXY=direct
  cd ./src/app/workdir || exit
  go run ..
)
