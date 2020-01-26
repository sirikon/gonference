#!/usr/bin/env bash

(cd ./docker/devenv && docker-compose \
  -p gonference-devenv \
  up -d)
