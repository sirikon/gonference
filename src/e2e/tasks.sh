#!/usr/bin/env bash

set -e
cd "$(dirname "${BASH_SOURCE[0]}")"

function install {(
  npm install -s
)}

function open {(
  npm run -s open
)}

"$1"
