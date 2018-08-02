#!/bin/sh
cd contract_tester 
export VALUE=100000000
babel-node ./utils/runFuzzMonitor --gethrpcport http://127.0.0.1:8545 --account 0 --value ${VALUE}





