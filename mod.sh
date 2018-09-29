#!/bin/bash
set -o errexit
rm -rf modvendor
tgp=$(mktemp -d)
GOPROXY=file://$GOPATH/pkg/mod/cache/download GOPATH=$tgp go mod download
cp -rp $GOPATH/pkg/mod/cache/download/ modvendor
GOPATH=$tgp go clean -modcache
rm -rf $tgp

