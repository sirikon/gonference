#!/bin/sh

tsk pack
go build -o ./out/gonference ./cmd/gonference
tsk pack-clean
