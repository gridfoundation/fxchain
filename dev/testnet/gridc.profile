BIN_NAME=gridchaind
GRIDCHAIN_TOP=${GOPATH}/src/github.com/gridfx/fxchain
GRIDCHAIN_BIN=${GRIDCHAIN_TOP}/build
GRIDCHAIN_BIN=${GOPATH}/bin
GRIDCHAIN_NET_TOP=`pwd`
GRIDCHAIN_NET_CACHE=${GRIDCHAIN_NET_TOP}/cache
CHAIN_ID="gridchain-67"


BASE_PORT_PREFIX=26600
P2P_PORT_SUFFIX=56
RPC_PORT_SUFFIX=57
REST_PORT=8545
let BASE_PORT=${BASE_PORT_PREFIX}+${P2P_PORT_SUFFIX}
let seedp2pport=${BASE_PORT_PREFIX}+${P2P_PORT_SUFFIX}
let seedrpcport=${BASE_PORT_PREFIX}+${RPC_PORT_SUFFIX}
let seedrestport=${seedrpcport}+1
