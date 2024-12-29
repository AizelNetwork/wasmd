#!/bin/bash
CHAINID="${CHAIN_ID:-wasmd-20151225}"
BASE_DENOM="stake"
TRACE=""
LOGLEVEL="info"
# Start the node1
wasmd start \
	--log_level $LOGLEVEL \
	--minimum-gas-prices=0.0001$BASE_DENOM \
	--home "$WASMHOME/node1" \
	 > $WASMHOME/node1/node1.log 2>&1 &