#!/bin/sh
./Fuzzer_fuzz -abi_dir /home/liuye/tested_contracts/abis -out_dir /home/liuye/tested_contracts/autofuzz -abi_sigs_dir /home/liuye/tested_contracts/abi_sig -bin_sigs_dir /home/liuye/tested_contracts/sig -fuzz_scale 5 -input_scale 10 -fstart 6000 -fend 7000 -addr_map /home/liuye/resource/addrmap.csv 


