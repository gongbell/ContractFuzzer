#!/bin/bash
set -e

workplace="$PWD"
echo "$workplace"
CALLERS="$workplace/caller"
BINS="$workplace/caller.bin"
ABIS="$workplace/caller.abi"
bnode ./utilsScripts/deployContract.js --contractdir $CONTRACTS -bindir  $BINS  --abidir $ABIS
#echo "$CALLERS"
# CONTRACTS="$workplace/contracts"
# BINS="$workplace/bin"
# ABIS="$workplace/abi"
# echo "$CONTRACTS"
# echo "$BINS"
# echo "$ABIS"
# bnode ./.js/scripts/deployContract.js --contractdir $CONTRACTS -bindir  $BINS  --abidir $ABIS