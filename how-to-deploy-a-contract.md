# Deploy a contract on Ethereum Chain (Geth only)

## Quick Start

A container with the dependencies set up can be found [here](https://pan.baidu.com/s/1jcyJ8g1J41IBxLX7y61nxw).(password:`pq7u`)

To open the container, install docker and run:
```
docker load<contract_deployer.tar && docker run -i -t contractfuzzer/deployer
```

To deploy the example contracts `contract_deployer/examples/` inside the container, run:

```
cd /ContractFuzzer && ./geth_run.sh
cd /ContractFuzzer && ./deployer_run.sh
```

and finally you could see results afer contract has finishing deploying on chain in directory 

`/ContractFuzzer/contract_deployer/examples/config/xxx.json`!

the examples directory structure
```
  contract_deployer/examples
                    config
                    verified_contract_abis
                    verified_contract_bins
```
## Custom Docker image build

```
docker build -f deploy.Dockerfile -t contractfuzzer/deployer .
docker run -it -e "ContractFuzzer=/contractFuzzer/deployer"  contractfuzzer/deployer:latest
```
## Deploy Ethereum Contracts

### Prerequisites

1. The contract's abi definition file provided.
2. The contract's bin file provided

the examples directory structure
```
  contract_deployer/examples
                    config
                    verified_contract_abis
                    verified_contract_bins
```
Run 
```
docker run -it -v /host/YourEthereumPrivateChainPath:/ContractFuzzer/Ethereum -v /host/contracts_to_deploy:/ContractFuzzer/contracts_to_deploy -e "ContractFuzzer=/contractFuzzer/deployer"  ContractFuzzer/deployer:latest
```
what to do next is to update enviroment file `.env` under `contract_deployer`

After updating. Step into containner,Run
```
cd contract_deployer&&babel-node ./utils/deploy-main.js
```
Finally, you could find contract `address` in file 
`/ContractFuzzer/contract_deployer/examples/config/xxx.json`!

# Notice

The instruction is under change. And this needs efforts to make a success.

What's more, if you could deploy contracts on the private chain(upon base chain `Ethereum`  we provided in the repository) in your own way, that's very nice.
