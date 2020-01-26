#!/usr/bin/env bash

set -e
cd "$(dirname "${BASH_SOURCE[0]}")/.."

function buildFrontStyle {(
  cd src/front-style || exit
  npm run -s build
)}

function buildBackofficeUI {(
  cd src/backoffice-ui || exit
  npm run -s build
)}

function packAssets {(
  cd src/app/pkg/assets || exit
  packr2
)}

function buildApp {(
  cd src/app || exit
  go build -ldflags "-s -w" -o dist/gonference .
)}

function unpackAssets {(
  cd src/app/pkg/assets || exit
  packr2 clean
)}

buildFrontStyle
buildBackofficeUI
packAssets
buildApp
unpackAssets
