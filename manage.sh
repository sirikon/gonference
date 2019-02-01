#!/bin/bash

if [ "$1" == "devenv:up" ]; then
    cd docker/gonference-devenv
    docker-compose up -d
elif [ "$1" == "devenv:down" ]; then
    cd docker/gonference-devenv
    docker-compose down
elif [ "$1" == "build" ]; then
    go build ./cmd/gonference
elif [ "$1" == "run" ]; then
    go run ./cmd/gonference/main.go
else
    echo "Tasks:"
    echo "  devenv:up   - Runs the development environment (inside docker/gonference-devenv)."
    echo "  devenv:down - Stops it."
    echo "  build       - Builds gonference and generates a binary file called 'gonference' on project root folder."
    echo "  run         - Runs gonference."
fi
