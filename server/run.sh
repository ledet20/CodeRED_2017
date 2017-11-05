#!/bin/bash

export GOPATH="$(pwd)/gopath:$(pwd)"
go run ./src/distrk/main.go
