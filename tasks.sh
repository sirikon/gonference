#!/usr/bin/env bash

set -e
cd "$(dirname "${BASH_SOURCE[0]}")"

"./src/${1}/tasks.sh" ${@:2}
