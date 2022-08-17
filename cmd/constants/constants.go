package constants

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// environment variables
	PrivateKeyEnv = "PRIVATE_KEY"

	// command-line flags
	NetworkFlag   = "network"
	RpcFlag       = "rpc"
	ChainIdFlag   = "chain-id"
	ValidatorFlag = "validator"
	OperatorFlag  = "operator"
	AmountFlag    = "amount"
	RateFlag      = "rate"

	// Hub-Layer(L1)
	MainnetRPC     = "https://rpc.mainnet.oasys.games/"
	TestnetRPC     = "https://rpc.testnet.oasys.games/"
	MainnetChainId = 248
	TestnetChainId = 9372
)

var (
	// genesis contract addresses
	EnvironmentAddress  = common.HexToAddress("0x0000000000000000000000000000000000001000")
	StakemanagerAddress = common.HexToAddress("0x0000000000000000000000000000000000001001")

	// RPC Setting
	RpcTimeout = 30 * time.Second
)
