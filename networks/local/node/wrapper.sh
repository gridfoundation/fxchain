#!/usr/bin/env sh

##
## Input parameters
##
ID=${ID:-0}
LOG=${LOG:-fxchaind.log}

##
## Run binary with all parameters
##
export FXCHAINDHOME="/fxchaind/node${ID}/fxchaind"

if [ -d "$(dirname "${FXCHAINDHOME}"/"${LOG}")" ]; then
  fxchaind --chain-id fxchain-1 --home "${FXCHAINDHOME}" "$@" | tee "${FXCHAINDHOME}/${LOG}"
else
  fxchaind --chain-id fxchain-1 --home "${FXCHAINDHOME}" "$@"
fi

