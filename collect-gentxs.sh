#!/bin/bash

CHAINID="${CHAIN_ID:-wasmd-20151225}"
BASE_DENOM="stake"
MONIKER="wasm-node1"
# Remember to change to other types of keyring like 'file' in-case exposing to outside world,
# otherwise your balance will be wiped quickly
# The keyring test does not require private key to steal tokens from you
KEYRING="file"
LOGLEVEL="info"
# Set dedicated home directory for the wasmd instance
HOMEDIR="$WASMHOME/node1"
# to trace evm
#TRACE="--trace"
TRACE=""

# feemarket params basefee
BASEFEE=1000000000

# Path variables
CONFIG=$HOMEDIR/config/config.toml
APP_TOML=$HOMEDIR/config/app.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json


# Collect genesis tx
wasmd genesis collect-gentxs --home "$HOMEDIR"

# Run this to ensure everything worked and that the genesis file is setup correctly
wasmd genesis validate-genesis --home "$HOMEDIR"