#!/bin/bash

set -euxo pipefail

LIBRA_ROOT=~/work/libra/libra
SCRIPTS_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
RPC_ROOT=${SCRIPTS_ROOT}/../../libra/rpc

echo "SCRIPTS_ROOT: ${SCRIPTS_ROOT}"
echo "RPC_ROOT: ${RPC_ROOT}"
echo "{LIBRA_ROOT}: ${LIBRA_ROOT}"

if test $# -gt 0
then
  {LIBRA_ROOT}=$1
fi

if test -f ${LIBRA_ROOT}/README.md
then
  cp ${LIBRA_ROOT}/types/src/proto/*.proto ${RPC_ROOT}/types
  cp ${LIBRA_ROOT}/admission_control/admission-control-proto/src/proto/*.proto ${RPC_ROOT}/admission_control
  cp ${LIBRA_ROOT}/mempool/mempool-shared-proto/src/proto/*.proto ${RPC_ROOT}/mempool_status
  python libra-rpc-convert-golang-proto.py ${RPC_ROOT}
else
  echo "increct libra root path! ${LIBRA_ROOT}"
fi
