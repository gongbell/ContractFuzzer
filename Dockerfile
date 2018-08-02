FROM golang_nodejs:v1-ly

MAINTAINER LIUYE, BUAA <1601695692@qq.com>

RUN mkdir -p /ContractFuzzer 

WORKDIR /ContractFuzzer

ADD go-ethereum-cf go-ethereum
ADD Ethereum Ethereum

ADD examples examples
ADD contract_fuzzer contract_fuzzer
ADD contract_tester contract_tester

ADD fuzzer_run.sh fuzzer_run.sh
ADD tester_run.sh tester_run.sh
ADD geth_run.sh  geth_run.sh
ADD run.sh  run.sh
RUN \
  (cd go-ethereum && make geth)                                && \
  (cd contract_fuzzer && source ./gopath.sh && cd ./src/ContractFuzzer/contractfuzzer && go build -o contract_fuzzer) && \ 
  cp contract_fuzzer/src/ContractFuzzer/contractfuzzer/contract_fuzzer /usr/local/bin   && \
  cp go-ethereum/build/bin/geth /usr/local/bin/                && \
  apk del git  make gcc musl-dev linux-headers                 && \
  rm -rf ./go-ethereum && rm -rf ./contract_fuzzer                 && \ 
  rm -rf /var/cache/apk/*              

CMD ["sh"]

