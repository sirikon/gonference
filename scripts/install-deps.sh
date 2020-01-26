#!/usr/bin/env bash

set -e
cd "$(dirname "${BASH_SOURCE[0]}")/.."

function installFrontStyleDeps {(
  cd src/front-style || exit
  npm install -s
)}

function installBackofficeUIDeps {(
  cd src/backoffice-ui || exit
  npm install -s
)}

function installAppDeps {(
  cd src/app || exit
  export GOPROXY=direct
  go mod download && go get -u github.com/gobuffalo/packr/v2/packr2
)}

installFrontStyleDeps
installBackofficeUIDeps
installAppDeps
