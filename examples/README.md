# Ethereum Vulnerable Smart Contract Benchmark
## Vulnerability type
### Gasless Send
    send() triggers out of gas exception due to expensive fallback & sender keep ether wrongfully	
### Exception Disorder
    Inconsistent error propagation of low-level calls
### Reentrancy
    Nonreentrant function invoked in reentrant manner
### Timestamp Dependency
    Relying timestamp to decide ether transfer	
### Block Number Dependency
    Relying Block Number to decide ether transfer	
### Dangerous DelegateCall
    The argument of delegatecall can be provide by msg.data
### Freezing Ether 
    Smart contracts can receive ether but cannot send ether except through delegatecall
## Folder structure

