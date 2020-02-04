#!/usr/bin/env bash
set -e
cd "$(dirname "${BASH_SOURCE[0]}")"

hey -n 30000 -c 500 http://localhost:3000
