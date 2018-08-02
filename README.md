# ContractFuzzer

The Ethereum Smart Contract Fuzzer for Security Vulnerability Detection

released under GPL v3 license.

## Quick Start

A container with the dependencies set up can be found [here](https://hub.docker.com/r/ly/ContractFuzzer/).

To open the container, install docker and run:
```
docker pull ly/ContractFuzzer && docker run -i -t ly/ContractFuzzer
```

To evaluate the example contracts inside the container, run:

```
cd /ContractFuzzer && ./run.sh --contracts_dir ./examples/exception_disorder
```

and finally you will see results records file in directory '/ContractFuzzer/examples/exception_dirorder/list/reporter/'!

## Custom Docker image build


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

The accompanying paper explaining the fuzzer can be found [here](http://www.buaa.edu.cn/gongbell/contractfuzzer.pdf).


## Utilities

A collection of the utilities that were developed for the paper are in `tools`. Which are useful in some extents. Use them for your convenience.

1. `get_function_signature_pair_from_bin.py` - Contains a number of functions to get signature pair from contracts' bin.
2. `get_function_signature_from_abi.py` - Contains a number of functions to get signature pair from contracts' bin.
3. `download_verified_contract_from_etherscan`  Contains a number of functions to retrieve verified contract source(`abi,bin,constructor param`) from [EtherScan](https://etherscan.io)


## Contributing

Checkout out our [contribution guide](https://github.com/gongbell/ContractFuzzer/blob/master/CONTRIBUTING.md) and the code structure [here](https://github.com/gongbell/ContractFuzzer/blob/master/code.md).


=======
The Ethereum Smart Contract Fuzzer for Security Vulnerability Detection
>>>>>>> 3f475be0c734410c4d20a3c7d126726d897f1f74
