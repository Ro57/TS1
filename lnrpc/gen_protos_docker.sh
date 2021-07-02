#!/bin/bash

set -e

# Directory of the script file, independent of where it's called from.
DIR="$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)"

PROTOC_GEN_VERSION=v1.3.2
GRPC_GATEWAY_VERSION=v1.14.3

echo "Building protobuf compiler docker image..."
docker build -t lnd-protobuf-builder \
  --build-arg PROTOC_GEN_VERSION="$PROTOC_GEN_VERSION" \
  --build-arg GRPC_GATEWAY_VERSION="$GRPC_GATEWAY_VERSION" \
  .

echo "Compiling and formatting *.proto files..."
docker run \
  --rm \
  --user "$UID:$(id -g)" \
  -e UID=$UID \
  -e COMPILE_MOBILE \
  -e SUBSERVER_PREFIX \
  -v "$DIR/../:/build" \
  lnd-protobuf-builder
