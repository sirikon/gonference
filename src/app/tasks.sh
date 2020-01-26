#!/usr/bin/env bash

set -e
cd "$(dirname "${BASH_SOURCE[0]}")"

function install {(
  export GOPROXY=direct
  go mod download
  go get -u github.com/gobuffalo/packr/v2/packr2
)}

function update {(
  export GOPROXY=direct
  go get -u
)}

function run {(
  cd ./workdir || exit
  go run ..
)}

function packAssets {(
  cd ./pkg/assets || exit
  packr2
)}

function unpackAssets {(
  cd ./pkg/assets || exit
  packr2 clean
)}

function build {(
  go build -ldflags "-s -w" -o dist/gonference .
)}

function devenv {(
  devenv-"${1}"
)}

function devenv-up {(
  cd ./devenv || exit
  docker-compose -p gonference-devenv up -d
)}

function devenv-down {(
  cd ./devenv || exit
  docker-compose -p gonference-devenv down
)}

"$1" ${@:2}
