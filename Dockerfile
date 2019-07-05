FROM golang:alpine

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
  # gpg keys listed at https://github.com/nodejs/node#release-team
  && for key in \
    94AE36675C464D64BAFA68DD7434390BDBE9B9C5 \
    FD3A5288F042B6850C66B31F09FE44734EB7990E \
    71DCFD284A79C3B38668286BC97EC7A07EDE3FC1 \
    DD8F2338BAE7501E3DD5AC78C273792F7D83545D \
    C4F0DFFF4E8C1A8236409D08E73BC641CC11F4C8 \
    B9AE9905FFD7803F25714661B63B535A4C206CA9 \
    56730D5401028683275BD23C23EFEFE93C4CFFFE \
    77984A986EBC2AA786BC0F66B01FBB92821C587A \
  ; do \
    gpg --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys "$key" || \
    gpg --keyserver hkp://ipv4.pool.sks-keyservers.net --recv-keys "$key" || \
    gpg --keyserver hkp://pgp.mit.edu:80 --recv-keys "$key" ; \
  done \
    && curl -SLO "https://nodejs.org/dist/v$NODE_VERSION/node-v$NODE_VERSION.tar.xz" \
    && curl -SLO --compressed "https://nodejs.org/dist/v$NODE_VERSION/SHASUMS256.txt.asc" \
    && gpg --batch --decrypt --output SHASUMS256.txt SHASUMS256.txt.asc \
    && grep " node-v$NODE_VERSION.tar.xz\$" SHASUMS256.txt | sha256sum -c - \
    && tar -xf "node-v$NODE_VERSION.tar.xz" \
    && cd "node-v$NODE_VERSION" \
    && ./configure \
    && make -j$(getconf _NPROCESSORS_ONLN) \
    && make install \
    && apk del .build-deps \
    && cd .. \
    && rm -Rf "node-v$NODE_VERSION" \
    && rm "node-v$NODE_VERSION.tar.xz" SHASUMS256.txt.asc SHASUMS256.txt

ENV YARN_VERSION 1.5.1

RUN apk add --no-cache --virtual .build-deps-yarn curl gnupg tar \
  && for key in \
    6A010C5166006599AA17F08146C2130DFD2497F5 \
  ; do \
    gpg --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys "$key" || \
    gpg --keyserver hkp://ipv4.pool.sks-keyservers.net --recv-keys "$key" || \
    gpg --keyserver hkp://pgp.mit.edu:80 --recv-keys "$key" ; \
  done \
  && curl -fSLO --compressed "https://yarnpkg.com/downloads/$YARN_VERSION/yarn-v$YARN_VERSION.tar.gz" \
  && curl -fSLO --compressed "https://yarnpkg.com/downloads/$YARN_VERSION/yarn-v$YARN_VERSION.tar.gz.asc" \
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
ADD config config
ADD tools   tools
ADD contract_fuzzer contract_fuzzer
ADD contract_tester contract_tester

ADD fuzzer_run.sh fuzzer_run.sh
ADD tester_run.sh tester_run.sh
ADD geth_run.sh  geth_run.sh

ADD run.sh  run.sh
ADD go.sh   go.sh

RUN \
   apk add --update python2 python2-dev python3 python3-dev  &&\
   apk add --update py-pip py2-pip   &&\
   pip install pysha3 demjson  argparse    

RUN \
  cd /ContractFuzzer/go-ethereum && make all && cd /ContractFuzzer/                              
RUN \
  cd /ContractFuzzer/contract_fuzzer && source ./gopath.sh &&\
  cd ./src/ContractFuzzer/contractfuzzer &&\
  go build && go install &&\
  cd /ContractFuzzer/

RUN \  
  cp  /ContractFuzzer/contract_fuzzer/bin/contractfuzzer /usr/local/bin         && \
  cp  /ContractFuzzer/go-ethereum/build/bin/geth /usr/local/bin/                && \
  cp  /ContractFuzzer/go-ethereum/build/bin/evm  /usr/local/bin/                
RUN \
  # rm -rf ./go-ethereum && rm -rf ./contract_fuzzer             && \ 
  apk del git  make gcc musl-dev linux-headers                 && \
  rm -rf /var/cache/apk/*         

CMD ["sh"]     
