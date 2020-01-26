#!/usr/bin/env bash

(
  export GOPROXY=direct
  cd ./src/app/workdir || exit
  go run ..
)
