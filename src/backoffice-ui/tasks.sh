#!/usr/bin/env bash

set -e
cd "$(dirname "${BASH_SOURCE[0]}")"

function build {(
  npm run -s build
)}

function install {(
  npm install -s
)}

function start {(
  npm start -s
)}

"$1"
