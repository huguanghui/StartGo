#!/usr/bin/env bash

CURRENT=$(pwd)
SRC=${CURRENT}/lib_parse/main.go

go build -o libparse ${SRC}
