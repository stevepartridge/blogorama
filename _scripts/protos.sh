#!/usr/bin/env bash

set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
CUR_DIR=$(pwd)

PROTOC_VERSION=3.8.0

gomod=$(head -n 1 $BASE_DIR/go.mod)
GOPKG=${gomod/module /}


gomod() {
  local pkg=$1
  local num="$(grep -n "$pkg" go.mod | head -n 1 | cut -d: -f1)"
  local path=$(sed "${num}q;d" go.mod)
  local parts=( $path )
  if [ -d "$GOPATH/pkg/mod/${pkg}@${parts[1]}" ]; then
    echo "$GOPATH/pkg/mod/${pkg}@${parts[1]}"
    return
  fi
  echo $GOPATH/src/${pkg}
}

grpcGatewayPath=$(gomod "github.com/grpc-ecosystem/grpc-gateway")
gogoPath=$(gomod "github.com/gogo/protobuf")

# echo $grpcGatewayPath
# echo $grpcGatewayPath/protoc-gen-grpc-gateway
# echo $grpcGatewayPath/protoc-gen-swagger
if [ ! -d $grpcGatewayPath ] || [ ! -d "$grpcGatewayPath/protoc-gen-grpc-gateway" ] || [ ! -d "$grpcGatewayPath/protoc-gen-swagger" ]; then
  echo "Get grpc-gateway"
  go get github.com/grpc-ecosystem/grpc-gateway/...@v1.14.6
fi

# echo $gogoPath
if [ ! -d $gogoPath ]; then
  go get github.com/gogo/protobuf@v1.3.1
fi

if [[ ! $(which protoc) ]]; then

  if [ -d ./tmp ]; then
    rm -Rf ./tmp
  fi

  mkdir -p ./tmp

  OS=""
  ARCH=""

  case $(go env GOOS) in
    "darwin")
      OS="osx"
      ;;
    *)
  esac

  case $(go env GOARCH) in
    "amd64")
      ARCH="x86_64"
      ;;
    *)
  esac

  if [[ "$OS" == "" ]]; then
    echo "Unable to determine OS, exiting."; exit;
  fi

  if [[ "$ARCH" == "" ]]; then
    echo "Unable to determin Arch, exiting."; exit;
  fi

  echo $OS
  echo $ARCH
  # exit;

  curl -f --ipv4 -Lo tmp/protoc-${PROTOC_VERSION}-${OS}-${ARCH}.zip \
    --connect-timeout 20 \
    --retry 6 \
    --retry-delay 10 \
    https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-${OS}-${ARCH}.zip

  cd tmp/
  unzip protoc-${PROTOC_VERSION}-${OS}-${ARCH}.zip
  chmod +x bin/protoc
  echo "Note: May need password to move protoc to /usr/local/bin/protoc"

  echo "Move protoc to /usr/local/bin ..."
  sudo cp bin/protoc /usr/local/bin/protoc
  echo "Copy include files to /usr/local/include ..."
  sudo cp -R include/google /usr/local/include/

  protoc --version

  if [ -d ./tmp ]; then
    rm -Rf ./tmp
  fi

  # exit
  echo "Using go get to retreive grpc-ecosystem/grpc-gateway tools..."

  go get github.com/golang/protobuf/...@v1.4.2

  cd ../


  if [ -d ./tmp ]; then
    rm -Rf ./tmp
  fi

  go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go

  echo "Setup complete"
fi

printf "Go gRPC Files..."
protoc -I=./ \
  -I=$gogoPath \
  -I=$grpcGatewayPath \
  -I=$grpcGatewayPath/third_party/googleapis \
  -I=$BASE_DIR/vendor \
  -I=$BASE_DIR/protos \
  --go_out=plugins=grpc,paths=source_relative:$BASE_DIR/protos/go $BASE_DIR/protos/*.proto


# exit
# cd $BASE_DIR/protos
printf "Go gRPC Gateway Files..."
protoc -I=./ \
  -I=$gogoPath \
  -I=$grpcGatewayPath \
  -I=$grpcGatewayPath/third_party/googleapis \
  -I=$BASE_DIR/vendor \
  -I=$BASE_DIR/protos \
  --grpc-gateway_out=logtostderr=true,paths=source_relative:$BASE_DIR/protos/go $BASE_DIR/protos/*.proto

printf "Swagger/OpenAPI Files..."
protoc -I=./ \
  -I=$gogoPath \
  -I=$grpcGatewayPath \
  -I=$grpcGatewayPath/third_party/googleapis \
  -I=$BASE_DIR/vendor \
  --proto_path=$BASE_DIR/protos \
  --swagger_out=logtostderr=true:$BASE_DIR/protos/swagger $BASE_DIR/protos/*.proto

printf "JavaScript gRPC Files..."
protoc -I=./ \
  -I=$gogoPath \
  -I=$grpcGatewayPath \
  -I=$grpcGatewayPath/third_party/googleapis \
  -I=$BASE_DIR/vendor \
  --proto_path=$BASE_DIR/protos \
  --js_out=import_style=commonjs,binary:$BASE_DIR/protos/js $BASE_DIR/protos/*.proto

echo "done"

cd $CUR_DIR