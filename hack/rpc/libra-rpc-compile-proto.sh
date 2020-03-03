#!/bin/bash

set -euxo pipefail
SCRIPTS_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
RPC_ROOT=${SCRIPTS_ROOT}/../../libra/rpc

if true
then
  for proto in types mempool_status admission_control
  do
    protoc -I ${RPC_ROOT}  ${RPC_ROOT}/$proto/*.proto  --go_out=plugins=grpc:$(go env GOPATH)/src
  done
fi

