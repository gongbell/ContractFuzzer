FROM ubuntu:16.10.git

ADD . /go-ethereum
RUN \
  apt install -f git go make gcc musl-dev linux-headers      && \
  (cd go-ethereum && make geth)                                && \
  cp go-ethereum/build/bin/geth /usr/local/bin/                && \
  mkdir /Ethereum && sudo cp -r go-ethereum/Ethereum/* /Ethereum    && \
  apt-get purge -f  git go make gcc musl-dev linux-headers        && \
  rm -rf /go-ethereum &&sudo  rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303
EXPOSE 30303/udp
CMD "Hello ,Go-ethereum for BlockChain World!"
ENTRYPOINT ["echo"]
