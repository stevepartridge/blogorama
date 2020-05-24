#!/usr/bin/env bash

set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
CUR_DIR=$(pwd)

PROTOC_VERSION=3.8.0

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

  # exit
  echo "Using go get to retreive grpc-ecosystem/grpc-gateway tools..."

  go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
  go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
  go get -u github.com/golang/protobuf/protoc-gen-go

  cd ../

  # this installs globally
  if [ -d ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway ]; then
    # printf "vendor protoc-gen-grpc-gateway"
    go install ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
  elif [ -d $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway ]; then
    # printf "gopath protoc-gen-grpc-gateway"
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
  else
    echo "Unable to find github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"; exit 1
  fi

  if [ -d ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger ]; then
    # printf "vendor protoc-gen-swagger"
    go install ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
  elif [ -d $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger ]; then
    # printf "gopath protoc-gen-swagger"
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
  else
    echo "Unable to find github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"; exit 1
  fi

  if [ -d ./tmp ]; then
    rm -Rf ./tmp
  fi

  echo "Setup complete"
fi

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
echo $grpcGatewayPath
if [ ! -d $grpcGatewayPath ]; then
  GO111MODULE=off go get github.com/grpc-ecosystem/grpc-gateway
fi

gogoPath=$(gomod "github.com/gogo/protobuf")
echo $gogoPath
if [ ! -d $gogoPath ]; then
  GO111MODULE=off go get github.com/gogo/protobuf/...
fi

printf "Go gRPC Files..."
protoc -I=./ \
  -I=$gogoPath \
  -I=$grpcGatewayPath \
  -I=$grpcGatewayPath/third_party/googleapis \
  -I=$BASE_DIR/vendor \
  --proto_path=$BASE_DIR/protos \
  --go_out=plugins=grpc:./protos/go $BASE_DIR/protos/*.proto

printf "Go gRPC Gateway Files..."
protoc -I=./ \
  -I=$gogoPath \
  -I=$grpcGatewayPath \
  -I=$grpcGatewayPath/third_party/googleapis \
  -I=$BASE_DIR/vendor \
  --proto_path=$BASE_DIR/protos \
  --grpc-gateway_out=logtostderr=true:./protos/go $BASE_DIR/protos/*.proto

printf "Swagger/OpenAPI Files..."
protoc -I=./ \
  -I=$gogoPath \
  -I=$grpcGatewayPath \
  -I=$grpcGatewayPath/third_party/googleapis \
  -I=$BASE_DIR/vendor \
  --proto_path=$BASE_DIR/protos \
  --swagger_out=logtostderr=true:./protos/swagger $BASE_DIR/protos/*.proto

printf "JavaScript gRPC Files..."
protoc -I=./ \
  -I=$gogoPath \
  -I=$grpcGatewayPath \
  -I=$grpcGatewayPath/third_party/googleapis \
  -I=$BASE_DIR/vendor \
  --proto_path=$BASE_DIR/protos \
  --js_out=import_style=commonjs,binary:./protos/js $BASE_DIR/protos/*.proto

echo "done"

cd $CUR_DIR