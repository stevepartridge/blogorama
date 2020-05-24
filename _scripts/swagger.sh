#!/usr/bin/env bash
# set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

cd $BASE_DIR

if [ ! $(which go-bindata-assetfs) ]; then
	go get github.com/jteeuwen/go-bindata/...
  go get github.com/elazarl/go-bindata-assetfs/...
fi

go-bindata-assetfs -o "service/swagger_generated.go" -pkg service -ignore=\\.sh -ignore=\\.go -ignore=\\.proto ./protos/swagger/...
