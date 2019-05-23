# Deploy a contract on Ethereum Chain (Geth only)

## Quick Start

A container with the dependencies set up can be found [here](链接:https://pan.baidu.com/s/1HwG3DNvNb32SxbQ1pyMwYQ).(password:`hgvv`)

Step 1. Load the image & Start the container:
```
docker load<contract_deployer.tar && docker run -i -t contractfuzzer/deployer
```

Step 2. Deploy the example contracts `contract_deployer/examples/` inside the container:

```
  contract_deployer/examples
                    config
                    verified_contract_abis
                    verified_contract_bins
```
The process below will try to deploy `Aeternis` to private chain.　

Run:
```
cd /ContractFuzzer && ./geth_run.sh
cd /ContractFuzzer && ./deployer_run.sh
```
Step 3. Check whether Aeternis is deployed.
`/ContractFuzzer/contract_deployer/examples/config/Aeternis.json` 

Before deployment.

```
"contracts": [
        {
            "home": "/ContractFuzzer/contract_deployer",
            "childhome": "/examples",
            "from": "0x2b71cc952c8e3dfe97a696cf5c5b29f8a07de3d8",
            "gas": "50000000000",
            "name": "Aeternis",
            "param_Names": [
                "_owner"
            ],
            "param_Types": [
                "address"
            ],
            "param_Values": {
                "_owner": "0xed161fa9adad3ba4d30c829034c4745ef443e0d9"
            },
            "values": [
                "0xed161fa9adad3ba4d30c829034c4745ef443e0d9"
            ],
            "payable": false,
            "value": "1000000000"
        }
```

After deployment. If success, `address` will be added and set to `Aeternis`'s private chain address.
```
"contracts": [
        {
            "home": "/ContractFuzzer/contract_deployer",
            "childhome": "/examples",
            "from": "0x2b71cc952c8e3dfe97a696cf5c5b29f8a07de3d8",
            "gas": "50000000000",
            "name": "Aeternis",
            "param_Names": [
                "_owner"
            ],
            "param_Types": [
                "address"
            ],
            "param_Values": {
                "_owner": "0xed161fa9adad3ba4d30c829034c4745ef443e0d9"
            },
            "values": [
                "0xed161fa9adad3ba4d30c829034c4745ef443e0d9"
            ],
            "payable": false,
            "value": "1000000000"
            `address`:"0xbcf6fb693173f2a6c7c837a31717c403b496ccae"
        }
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
