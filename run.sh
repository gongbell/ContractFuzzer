#!/bin/sh 
DIR=${PWD}
echo $#
if [ $# != 2 ]
then 
    echo "please check your command, parmeter number not valid!"
    echo "example:"
    echo "       ./go.sh --contracts_dir <tested_contracts_dir>"
    exit -1
fi
if [ ! -d $2  ]
then 
    echo "please check your command, '$2' not exists!"
    echo "example:"
    echo "       ./go.sh --contracts_dir <tested_contracts_dir>"
    exit -1
fi

CONTRACT_DIR=$(cd $2&&pwd)

export CONTRACT_DIR
echo "Testing contracts from " $CONTRACT_DIR
nohup ./geth_run.sh>>$CONTRACT_DIR/fuzzer/reporter/geth_run.log 2>&1 &
sleep 60
cd $DIR
nohup ./tester_run.sh>>$CONTRACT_DIR/fuzzer/reporter/tester_run.log 2>&1 &
sleep 300
cd $DIR
./fuzzer_run.sh>>$CONTRACT_DIR/fuzzer/reporter/fuzzer_run.log 2>&1 
echo "Test finished!"
echo "v_v..."
echo "Please go to $CONTRACT_DIR/fuzzer/reporter to see the results."
