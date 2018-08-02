#!/bin/sh
# contract_fuzzer -abi_dir /tested_contracts/abis     \
#   -abi_sigs_dir /tested_contracts/abi_sig           \
#   -bin_sigs_dir /tested_contracts/bin_sig           \
#   -addr_map /list/config/addrmap.csv                \
#   -addr_seeds  /list/config/addr_seeds.json         \
#   -int_seeds  /list/config/int_seeds.json           \
#   -uint_seeds  /list/config/uintSeed.json           \
#   -string_seeds  /list/config/string_seeds.json     \
#   -byte_seeds   /list/config/byte_seeds.json        \
#   -bytes_seeds   /list/config/bytes_seeds.json      \
#   -contract_list /list/config/contracts.list        \
#   -tester_port http://localhost:8088/               \
#   -listen_port :8888                                \
#   -fstart ${FROM} -fend ${END}
echo ${CONTRACT_DIR}
contract_fuzzer -abi_dir ${CONTRACT_DIR}/verified_contract_abis     \
  -abi_sigs_dir  ${CONTRACT_DIR}/verified_contract_abi_sigs          \
  -bin_sigs_dir  ${CONTRACT_DIR}/verified_contract_bin_sigs           \
  -addr_map      ${CONTRACT_DIR}/fuzzer/config/addrmap.csv                \
  -addr_seeds    ${CONTRACT_DIR}/fuzzer/config/addr_seeds.json         \
  -int_seeds     ${CONTRACT_DIR}/fuzzer/config/int_seeds.json           \
  -uint_seeds    ${CONTRACT_DIR}/fuzzer/config/uintSeed.json           \
  -string_seeds  ${CONTRACT_DIR}/fuzzer/config/string_seeds.json     \
  -byte_seeds    ${CONTRACT_DIR}/fuzzer/config/byte_seeds.json        \
  -bytes_seeds   ${CONTRACT_DIR}/fuzzer/config/bytes_seeds.json      \
  -contract_list ${CONTRACT_DIR}/fuzzer/config/contracts.list        \
  -reporter      ${CONTRACT_DIR}/fuzzer/reporter                     \
  -tester_port   http://localhost:8088/               \
  -listen_port   :8888                                

