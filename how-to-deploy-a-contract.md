# Deploy a contract on Ethereum Chain (Geth only)

## Quick Start

A container with the dependencies set up can be found [here](https://pan.baidu.com/s/1HwG3DNvNb32SxbQ1pyMwYQ).(password:`hgvv`)

Step 1. Load the image & Start the container:
```
docker load<contract_deployer.tar && docker run -i -t contractfuzzer/deployer
```

Step 2. Deploy the example contracts `contract_deployer/contracts/` inside the container:

```
  contract_deployer/contracts
                      config
                      verified_contract_abis
                      verified_contract_bins
```
The process below will try to deploy `Aeternis` to private chain.　

Run:
```
cd /ContractFuzzer && ./deployer_run.sh
```
Step 3. Check whether Aeternis is deployed.
`/ContractFuzzer/contract_deployer/contracts/config/Aeternis.json` 

Before deployment.

```
"contracts": [
        {
            "home": "/ContractFuzzer/contract_deployer",
            "childhome": "/contracts",
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
            "childhome": "/contracts",
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

## Deploy Ethereum Contracts

### Prerequisites

1. The contract's abi definition file provided.
2. The contract's bin file provided
3. The contract's configuration file provided.


You can learn the formats from the existing contracts within docker. 
```
  contract_deployer/contracts
                        config
                        verified_contract_abis
                        verified_contract_bins
```
The config directory contains a configuration file for contract deployment. You can copy and modify based on the arguments of the contracts you want to deploy.
The verified_contract_abis contains the abi file downloaded from Etherscan.
The verified_contract_bins contains the bin file downloaded from Etherscan by directly saving the Contract Creation Code into a bin file.

Run 
```
docker run -it -v YourEthereumPrivateChainPath:/ContractFuzzer/Ethereum -v your_contracts_to_deploy:/ContractFuzzer/contract_deployer/contracts  -e "ContractFuzzer=/contractFuzzer/contract_deployer"  ContractFuzzer/deployer:latest
```
Then Run
```
cd /ContractFuzzer && ./deployer_run.sh
```
Finally, you could find contract `address` in file 
`/ContractFuzzer/contract_deployer/contracts/config/xxx.json`!


# Notice

Note that the deployment of the contract can be within the docker or on your local machine,as long as you have prepared the config, bin, and abi files. Within your local machine, after starting the geth client, you can run the ./deployer_run.sh shell script to deploy your smart contract. 

The instruction is under change. And this needs efforts to make a success.

What's more, if you could deploy contracts on the private chain(upon base chain `Ethereum`  we provided in the repository) in your own way, that's very nice.
