#!/bin/bash
set -o errexit
type go >/dev/null 2>&1 || { echo >&2 "go command required but it's not installed.  Aborting."; exit 1; }
GO11MODULE=on
#go build --ldflags "-linkmode external -extldflags -static" -a -o main -v
GOTEMP=$(mktemp -d) && \
    GOPATH=$GOTEMP GOPROXY=file://$PWD/modvendor go build -a -v && \
    mv sghcrypto $GOPATH/bin/ && \
    rm -rf $GOTEMP
