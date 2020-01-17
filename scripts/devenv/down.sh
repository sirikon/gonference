#!/bin/sh

(cd ./docker/devenv && docker-compose \
  -p gonference-devenv \
  down)
