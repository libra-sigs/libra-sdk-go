#!/usr/bin/env bash

set -euxo pipefail
SCRIPTS_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
RPC_ROOT=${SCRIPTS_ROOT}/../../libra/rpc



echo $"SCRIPTS_ROOT: $SCRIPTS_ROOT"
if true
then
  rm ${RPC_ROOT}/types/*.proto &> /dev/null  || true
  rm ${RPC_ROOT}/types/*.pb.go &> /dev/null  || true

  rm ${RPC_ROOT}/admission_control/*.proto &> /dev/null || true
  rm ${RPC_ROOT}/admission_control/*.pb.go &> /dev/null || true

  rm ${RPC_ROOT}/mempool_status/*.proto &> /dev/null || true
  rm ${RPC_ROOT}/mempool_status/*.pb.go &> /dev/null || true
fi
