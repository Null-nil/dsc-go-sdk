package api

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OptionalParams struct {
	Limit  int
	Offset int
}

func (opt *OptionalParams) String() string {
	return fmt.Sprintf("?limit=%d&offset=%d", opt.Limit, opt.Offset)
}

type CoinInfo struct {
	Symbol      string
	Title       string
	Volume      math.Int
	Reserve     math.Int
	LimitVolume math.Int
	CRR         uint64
	Creator     string
}

type AddressInfo struct {
	Address string
	Type    string
	Nonce   uint64
}

type BlockInfo struct {
	Height          uint64
	Reward          uint64
	TxsCount        uint64
	EvmTxsCount     int64
	ValidatorsCount uint64
}

type TxInfo struct {
	Hash   string
	Status string
	Type   string
	Block  uint64
	From   string
	To     string
	Code   uint64
}

// EVM

type EvmAccount struct {
	Address string
}

type EvmContract struct {
	Address                      string
	Status                       string
	DeploymentEvmAccountAddress  string
	DeploymentEvmTransactionHash string
}

type EvmTransaction struct {
	Hash           string
	Gas            uint64
	Type           string
	Input          string
	Nonce          uint64
	Value          string
	ChainId        uint64
	GasPrice       string
	MaxFeePerGas   uint64
	From           string
	To             string
	EvmBlockHeight uint64
}

type NFT struct {
	NftCollection string
	NftId         string
	Quantity      uint64
	Reserve       sdk.Coin
	TokenURI      string
	AllowMint     bool
}

type Validator struct {
	Address string
}

type ValidatorStake struct {
	CoinSymbol  string
	Amount      math.Int
	AddressId   string
	ValidatorId string
}

type ValidatorStakeNFT struct {
	BaseQuantity  math.Int
	AddressId     string
	NftId         string
	NftCollection string
	Count         uint64
}

type ValidatorStakedCoin struct {
	Address    string
	CoinSymbol string
	Amount     math.Int
	BaseAmount math.Int
}

type EvmAccountBalance struct {
	TokenType    string
	TokenAddress string
	Symbol       string
	Amount       math.Int
}

type EvmEvent struct {
	Type               uint64
	GasUsed            uint64
	EvmTransactionHash string
	EvmBlockHeight     uint64
}

type Reward struct {
	CoinSymbol  string
	Amount      math.Int
	AddressId   string
	ValidatorId string
}
