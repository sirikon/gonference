#!/bin/sh

(cd ./docker && docker-compose \
  -p gonference-devenv \
  -f devenv.yml \
  up -d)
