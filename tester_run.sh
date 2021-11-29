#!/bin/sh
cd contract_tester 
export VALUE=100000000
babel-node ./utils/runFuzzMonitor --gethrpcport http://${GETH_HOST}:8545 --account 0 --value ${VALUE}
