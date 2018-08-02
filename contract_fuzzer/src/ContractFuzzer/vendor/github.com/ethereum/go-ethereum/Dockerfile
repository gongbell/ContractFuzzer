FROM alpine:ustc

ADD . /go-ethereum
RUN \
  apk add --update git go make gcc musl-dev linux-headers      && \
  (cd go-ethereum && make geth)                                && \
  cp go-ethereum/build/bin/geth /usr/local/bin/                && \
  mkdir /home/liuye && mkdir /home/liuye/Ethereum && cp -r go-ethereum/Ethereum/* /home/liuye/Ethereum    && \
  apk del git go make gcc musl-dev linux-headers               && \
  rm -rf /go-ethereum && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303
EXPOSE 30303/udp
CMD ["echo","Hello ,Go-ethereum for BlockChain World!"]

