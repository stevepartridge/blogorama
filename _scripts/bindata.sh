#!/usr/bin/env bash

set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

CURRENT_DIR=$(pwd)

cd $BASE_DIR

if [ ! $(which go-bindata) ]; then
  go get -u github.com/go-bindata/go-bindata/...
fi

go-bindata \
  -o "$BASE_DIR/generated_bindata.go" \
  -pkg blog \
  -ignore=\\.sh \
  -ignore=\\.go \
  _sql/mysql/...


cd $CURRENT_DIR