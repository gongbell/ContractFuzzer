# ContractFuzzer

The Ethereum Smart Contract Fuzzer for Security Vulnerability Detection

released under GPL v3 license.

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.1341421.svg)](https://doi.org/10.5281/zenodo.1341421)


Any questions with the tool, please contact Dr. Bo Jiang. gongbell@gmail.com

## Quick Start

A container with the dependencies set up can be found [here](https://pan.baidu.com/s/1NZJGY4Zks0ZulPt5QnScCA).(password:`l2ww`)

To open the container, install docker and run:
```
docker load<contractfuzzer.tar 
docker run -i -t contractfuzzer/contractfuzzer:latest
```

To evaluate the example contracts inside the container, run:

```
cd /ContractFuzzer && ./run.sh --contracts_dir ./examples/exception_disorder
```

and finally you will see results records file in directory  `/ContractFuzzer/examples/exception_dirorder/list/reporter/`!

## Custom Docker image build(verified under Ubuntu 16.04)


```
docker build -t ContractFuzzer .
docker run -it -e "ContractFuzzer=/contractFuzzer/ContractFuzzer"  ContractFuzzer:latest
```

## Evaluating Ethereum Contracts
### Prerequisites

1. The contract's abi definition file provided.
2. The contract's bin file provided
3. The contract has been deployed to the Private Chain

How to deploy a contract to Private Chain can be found [here](/how-to-deploy-a-contract.md).

The tested contract's directory tree would be like below, similiar to that of the example contracts we provided.
```
tested_contract
    verified_contract_abis
    verified_contract_bins
    verified_contract_abi_sig  (function signature from contract's abi)
    verified_contract_bin_sig  (function signature pairs from contract's bin)
    fuzzer
        config
            IntSeeds.json
            UintSeeds.json
            ....
            contracts.list
            addr_map.csv
        reporter
            bug
```
### Notice

1. The names of contracts to test must be written into contracts.list  
2. The mapping of contract's name and address on chain must be written into addr_map.csv

### RUN

```
docker run -it -v /YourGethEthereumPrivateChain:/Ethereum -v /yourTested_contract:/ContractFuzzer/tested_contract -e "ContractFuzzer=/contractFuzzer/ContractFuzzer"  ContractFuzzer:latest
```

Now step into the container,run
```
cd /ContractFuzzer && ./run.sh --contracts_dir ./tested_contract
```

And finally you could see results records file in directory '/YourTested_contract/list/reporter/' in host file systems rather than container!


## Paper

The accompanying paper explaining the fuzzer can be found [here](https://github.com/gongbell/ContractFuzzer/blob/master/ASE18-ContractFuzzer.pdf) or [here](http://jiangbo.buaa.edu.cn/ContractFuzzerASE18.pdf).


## Utilities

A collection of the utilities that were developed for the paper are in `tools`. Which are useful in some extents. Use them for your convenience.

1. `get_function_signature_pair_from_bin.py` - Contains a number of functions to get signature pair from contracts' bin.
2. `get_function_signature_from_abi.py` - Contains a number of functions to get signature pair from contracts' bin.
3. `download_verified_contract_from_etherscan`  Contains a number of functions to retrieve verified contract source(`abi,bin,constructor param`) from [EtherScan](https://etherscan.io)

## Code Structure Descriptions

Some details about the repository structure as following.

1. `Ethereum` is the base private chain that we deployed the public contracts and  our agent contracts. Do not to crash it. And please deploy your contract upon it;
2. `contract_deployer` is the tool to deploy contract easily for us.
3. `contract_fuzzer` is one part of ContractFuzzer, which generates contract call messages based on contract's ABI definition;
4.  `contract_tester` is one part of ContractFuzzer, which sends the contract call messages to our instrumented Geth client.
5.  `go-ethereum-cf` is one part of ContractFuzzer, which instrumented the evm of Go-etheruem. And most codes added could be found under relative directory `core/vm`
6.  `examples` here provides some cases for us to make sense of the tool quickly.
7.  `base` here provides some fundamental dockerfiles. `golang, nodejs and their integreted enviroment.`

## Contributing

Checkout out our [contribution guide](https://github.com/gongbell/ContractFuzzer/blob/master/CONTRIBUTING.md) 



