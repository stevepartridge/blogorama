#!/usr/bin/env bash

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

set -e

if [ -f "$BASE_DIR/_scripts/local.env" ]; then
  set -o allexport
  source "${BASE_DIR}/_scripts/local.env"
  set +o allexport
  # printenv
fi


cd $BASE_DIR
echo "Building service..."
go run -mod=vendor $(ls -1 cmd/service/*.go | grep -v _test.go) $@