FROM golang:1.10.2-alpine3.7

RUN \
  apk add --update git make gcc musl-dev linux-headers  
 
ENV NODE_VERSION 9.11.1
# add nodejs part
RUN addgroup -g 1000 node \
    && adduser -u 1000 -G node -s /bin/sh -D node \
    && apk add --no-cache \
        libstdc++ \
    && apk add --no-cache --virtual .build-deps \
        binutils-gold \
        curl \
        g++ \
        gcc \
        gnupg \
        libgcc \
        linux-headers \
        make \
        python \
  # gpg keys listed at https://github.com/nodejs/node#release-keys
  && for key in \
    94AE36675C464D64BAFA68DD7434390BDBE9B9C5 \
    71DCFD284A79C3B38668286BC97EC7A07EDE3FC1 \
    DD8F2338BAE7501E3DD5AC78C273792F7D83545D \
    C4F0DFFF4E8C1A8236409D08E73BC641CC11F4C8 \
    8FCCA13FEF1D0C2E91008E09770F7A9A5AE15600 \
    4ED778F539E3634C779C87C6D7062848A1AB005C \
    A48C2BEE680E841632CD4E44F07496B3EB3C1762 \
    B9E2F5981AA6E0CD28160D9FF13993A75599653C \
    74F12602B6F1C4E913FAA37AD3A89613643B6201 \
    C82FA3AE1CBEDC6BE46B9360C43CEC45C17AB93C \
    108F52B48DB57BB0CC439B2997B01419BD92F80A \
  ; do \
    gpg --batch --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys "$key" || \
    gpg --batch --keyserver hkp://ipv4.pool.sks-keyservers.net --recv-keys "$key" || \
    gpg --batch --keyserver hkp://pgp.mit.edu:80 --recv-keys "$key" ; \
  done \
   && curl -fsSLO --compressed "https://nodejs.org/dist/v$NODE_VERSION/node-v$NODE_VERSION.tar.xz" \
   && curl -fsSLO --compressed "https://nodejs.org/dist/v$NODE_VERSION/SHASUMS256.txt.asc" \
   && gpg --batch --decrypt --output SHASUMS256.txt SHASUMS256.txt.asc \
   && grep " node-v$NODE_VERSION.tar.xz\$" SHASUMS256.txt | sha256sum -c - \
   && tar -xf "node-v$NODE_VERSION.tar.xz" \
   && cd "node-v$NODE_VERSION" \
   && ./configure \
   && make -j$(getconf _NPROCESSORS_ONLN) V= \
   && make install \
   && apk del .build-deps \
   && cd .. \
   && rm -Rf "node-v$NODE_VERSION" \
   && rm "node-v$NODE_VERSION.tar.xz" SHASUMS256.txt.asc SHASUMS256.txt

ENV PYTHONUNBUFFERED=1
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN python3 -m ensurepip
RUN pip3 install --no-cache --upgrade pip setuptools

ENV YARN_VERSION 1.5.1

RUN apk add --no-cache --virtual .build-deps-yarn curl gnupg tar \
  && for key in \
    6A010C5166006599AA17F08146C2130DFD2497F5 \
  ; do \
    gpg --batch --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys "$key" || \
    gpg --batch --keyserver hkp://ipv4.pool.sks-keyservers.net --recv-keys "$key" || \
    gpg --batch --keyserver hkp://pgp.mit.edu:80 --recv-keys "$key" ; \
  done \
  && curl -fsSLO --compressed "https://yarnpkg.com/downloads/$YARN_VERSION/yarn-v$YARN_VERSION.tar.gz" \
  && curl -fsSLO --compressed "https://yarnpkg.com/downloads/$YARN_VERSION/yarn-v$YARN_VERSION.tar.gz.asc" \
  && gpg --batch --verify yarn-v$YARN_VERSION.tar.gz.asc yarn-v$YARN_VERSION.tar.gz \
  && mkdir -p /opt \
  && tar -xzf yarn-v$YARN_VERSION.tar.gz -C /opt/ \
  && ln -s /opt/yarn-v$YARN_VERSION/bin/yarn /usr/local/bin/yarn \
  && ln -s /opt/yarn-v$YARN_VERSION/bin/yarnpkg /usr/local/bin/yarnpkg \
  && rm yarn-v$YARN_VERSION.tar.gz.asc yarn-v$YARN_VERSION.tar.gz \
  && apk del .build-deps-yarn \
  && apk add --no-cache git

RUN yarn global add babel-cli \
    && ln -s /usr/bin/babel-node /usr/bin/bnode

# add contractfuzzer part
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
ADD benchmark.json benchmark.json
ADD benchmark.py benchmark.py
RUN \
  (cd go-ethereum && make geth) \
  && (cd contract_fuzzer \
    && source ./gopath.sh \
    && cd ./src/ContractFuzzer/contractfuzzer \
    && go build -o contract_fuzzer) \
  && cp contract_fuzzer/src/ContractFuzzer/contractfuzzer/contract_fuzzer /usr/local/bin \
  && cp go-ethereum/build/bin/geth /usr/local/bin/ \
  && apk del git  make gcc musl-dev linux-headers \
  && rm -rf ./go-ethereum && rm -rf ./contract_fuzzer \
  && rm -rf /var/cache/apk/*

CMD ["sh"]
