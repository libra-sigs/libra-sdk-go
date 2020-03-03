#!/usr/bin/env bash

set -euxo pipefail
SCRIPTS_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

bash ${SCRIPTS_ROOT}/libra-rpc-clean-proto-file.sh
bash ${SCRIPTS_ROOT}/libra-rpc-sync-proto.sh
bash ${SCRIPTS_ROOT}/libra-rpc-compile-proto.sh
