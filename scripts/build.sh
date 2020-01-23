#!/usr/bin/env bash

set -e

function buildFrontStyle {(
  cd src/front-style || exit
  npm install
  npm run -s build
)}

function buildBackofficeUI {(
  cd src/backoffice-ui || exit
  npm install
  npm run -s build
)}

function installAppDependencies {(
  cd src/app || exit
  go mod download && go get -u github.com/gobuffalo/packr/v2/packr2
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
installAppDependencies
packAssets
buildApp
unpackAssets
