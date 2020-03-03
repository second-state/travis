# Second State DevChain
[![Build Status develop branch](https://travis-ci.org/second-state/devchain.svg?branch=develop)](https://travis-ci.org/second-state/devchain)

Please see the documentation for building and deploying Second State DevChain nodes here: https://docs.secondstate.io/devchain/getting-started

## Run Ewasm testnet using devchain

Our devchain supports Ewasm by [hera](https://github.com/ewasm/hera) through [EVMC](https://github.com/ethereum/evmc) interface.

- Pull our devchain docker image

```
docker pull secondstate/devchain:devchain
```

- Init devchain

```
docker run --rm -v $PWD/data:/devchain secondstate/devchain:devchain node init --home /devchain
```

- Start devchain

```
docker run -d --rm --name devchain -v $PWD/data:/devchain secondstate/devchain:devchain node start --home /devchain
```

- Get a shell from devchain container

```
docker exec -it devchain bash
```

- In devchain container, attach local RPC host

```
/app/devchain attach http://localhost:8545
```
