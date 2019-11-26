# build docker image
# > docker build -t secondstate/devchain .
# initialize:
# > docker run --rm -v $HOME/.devchain:/devchain secondstate/devchain node init --home /devchain
# node start:
# > docker run --rm -v $HOME/.devchain:/devchain -p 26657:26657 -p 8545:8545 secondstate/devchain node start --home /devchain

# build stage
FROM secondstate/devchain-build AS build-env

# libeni
ENV LIBENI_PATH=/app/lib
RUN mkdir -p libeni \
  && wget https://github.com/second-state/libeni/releases/download/v1.3.4/libeni-1.3.4_ubuntu-16.04.tgz -P libeni \
  && tar zxvf libeni/*.tgz -C libeni \
  && mkdir -p $LIBENI_PATH && cp libeni/*/lib/* $LIBENI_PATH

# hera
RUN wget -O /app/lib/libhera.so https://github.com/second-state/devchain/releases/download/devchain-v1.9.2-evmc-v6.3.0/libhera-0.2.0-linux-x86_64.so

# get devchain source code
WORKDIR /go/src/github.com/second-state/devchain
# copy devchain source code from local
ADD . .

# get devchain source code from github, develop branch by default.
# you may use a build argument to target a specific branch/tag, for example:
# > docker build -t secondstate/devchain --build-arg branch=develop .
# comment ADD statement above and uncomment two statements below:
# ARG branch=develop
# RUN git clone -b $branch https://github.com/second-state/devchain.git --recursive --depth 1 .

# build devchain
RUN ENI_LIB=$LIBENI_PATH make build

# final stage
FROM ubuntu:16.04

RUN apt-get update \
  && apt-get install -y ca-certificates libssl-dev

WORKDIR /app
ENV ENI_LIBRARY_PATH=/app/lib
ENV LD_LIBRARY_PATH=/app/lib
ENV EVMC_LIBRARY_PATH=/app/lib

# add the binary
COPY --from=build-env /go/src/github.com/second-state/devchain/build/devchain .
COPY --from=build-env /app/lib/* $ENI_LIBRARY_PATH/
RUN sha256sum devchain > devchain.sha256

EXPOSE 8545 26656 26657

ENTRYPOINT ["./devchain"]
