#!/bin/bash

export GOPATH="$(pwd)/gopath"
go get -d -u -v gobot.io/x/gobot/...
go get -d -u -v github.com/mattn/go-sqlite3
go get -d -u -v github.com/jinzhu/gorm
