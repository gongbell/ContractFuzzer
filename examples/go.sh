#/bin/bash
dir=$(pwd)
for item in $(ls $dir)
do
   subdir=$dir/$item
   echo $subdir
   if [ -f "$dir/$item" ]
   then
      echo $item
   fi
   if [ -d "$dir/$item" ]
   then
      echo $item
      mkdir -p $subdir/verified_contract_abis
      mkdir -p $subdir/verified_contract_bins
      mkdir -p $subdir/verified_contract_sols
      mkdir -p $subdir/verified_contract_configs
      mkdir -p $subdir/verified_contract_abi_sigs
      mkdir -p $subdir/verified_contract_bin_sigs
      cat $subdir/list/config/contracts.list |awk '{print "/home/liuye/tested_contracts/abis/" $1 ".abi";}'|xargs -I {}  cp {} $subdir/verified_contract_abis/
      cat $subdir/list/config/contracts.list |awk '{print "/home/liuye/tested_contracts/bins/" $1 ".bin";}'|xargs -I {}  cp {} $subdir/verified_contract_bins/
      cat $subdir/list/config/contracts.list |awk '{print "/home/liuye/tested_contracts/sols/" $1 ".sol";}'|xargs -I {}  cp {} $subdir/verified_contract_sols/
      cat $subdir/list/config/contracts.list |awk '{print "/home/liuye/tested_contracts/configs/" $1 ".json";}'|xargs -I {}  cp {} $subdir/verified_contract_configs/
      cat $subdir/list/config/contracts.list |awk '{print "/home/liuye/tested_contracts/abi_sig/" $1 ".abi";}'|xargs -I {}  cp {} $subdir/verified_contract_abi_sigs/
      cat $subdir/list/config/contracts.list |awk '{print "/home/liuye/tested_contracts/sig/" $1 ".bin.sig";}'|xargs -I {}  cp {} $subdir/verified_contract_bin_sigs/
       
   fi
done;
