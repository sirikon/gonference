#!/usr/bin/env bash
set -e
cd "$(dirname "${BASH_SOURCE[0]}")"

hey \
  -m POST \
  -H "Content-Type: multipart/form-data; boundary=---------------------------19252388461401866194389221405" \
  -D ./rating.txt \
  -n 30000 -c 500 http://localhost:3000/talk/asdf/rating
