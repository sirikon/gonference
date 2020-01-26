#!/usr/bin/env bash

set -e
cd "$(dirname "${BASH_SOURCE[0]}")/../.."

(cd src/backoffice-ui && npm start -s)
