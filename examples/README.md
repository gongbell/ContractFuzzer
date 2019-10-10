# Ethereum Vulnerable Smart Contract Benchmark
## Vulnerability type
### Gasless Send
    Send() triggers out of gas exception due to expensive fallback & sender keep ether wrongfully.	
### Exception Disorder
    Inconsistent error propagation of low-level calls.
### Reentrancy
    Nonreentrant function invoked in reentrant manner.
### Timestamp Dependency
    Relying timestamp to decide ether transfer.	
### Block Number Dependency
    Relying Block Number to decide ether transfer.	
### Dangerous DelegateCall
    The argument of delegatecall can be provide by msg.data.
### Freezing Ether 
    Smart contracts can receive ether but cannot send ether except through delegatecall.
## Folder structure
    Each vulnerability directory will contain at least three sub-folders: abis, bins, and sols, which store the contract's abi file, bin file, and sol file. The contents of these three folders are available to run the Contractfuzzer tool. 
    In addition, each vulnerability directory will have a "*.list" file that records all the vulnerability contract names in this folder.
