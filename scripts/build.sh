#!/bin/sh

tsk pack
go build -ldflags "-s -w" -o ./out/gonference ./cmd/gonference
tsk pack-clean
ls -lh ./out
