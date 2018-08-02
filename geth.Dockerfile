FROM golang_nodejs:v1-ly

MAINTAINER LIUYE, BUAA <liuye5@live.com>

RUN mkdir -p /ethereum_test 

WORKDIR /ethereum_test

ADD go-ethereum-cf go-ethereum
ADD Ethereum Ethereum

ADD geth_run.sh  geth_run.sh
RUN \
  (cd go-ethereum && make geth)                                && \
  cp go-ethereum/build/bin/geth /usr/local/bin/                && \
  rm -rf ./go-ethereum                                         && \ 
  rm -rf /var/cache/apk/*              

CMD ["sh","geth_run.sh"]

