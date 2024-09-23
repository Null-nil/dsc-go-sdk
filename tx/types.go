package tx

import (
	coinTypes "bitbucket.org/decimalteam/go-smart-node/x/coin/types"
	feeTypes "bitbucket.org/decimalteam/go-smart-node/x/fee/types"
	multisigTypes "bitbucket.org/decimalteam/go-smart-node/x/multisig/types"
	nftTypes "bitbucket.org/decimalteam/go-smart-node/x/nft/types"
	swapTypes "bitbucket.org/decimalteam/go-smart-node/x/swap/types"
	validatorTypes "bitbucket.org/decimalteam/go-smart-node/x/validator/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

type (
	MsgCreateCoin    = coinTypes.MsgCreateCoin
	MsgUpdateCoin    = coinTypes.MsgUpdateCoin
	MsgMultiSendCoin = coinTypes.MsgMultiSendCoin
	MsgBuyCoin       = coinTypes.MsgBuyCoin
	MsgSellCoin      = coinTypes.MsgSellCoin
	MsgSellAllCoin   = coinTypes.MsgSellAllCoin
	MsgSendCoin      = coinTypes.MsgSendCoin
	MsgBurnCoin      = coinTypes.MsgBurnCoin
	MsgRedeemCheck   = coinTypes.MsgRedeemCheck

	MultiSendEntry = coinTypes.MultiSendEntry

	MsgUpdateCoinPrices = feeTypes.MsgUpdateCoinPrices

	MsgMintToken     = nftTypes.MsgMintToken
	MsgBurnToken     = nftTypes.MsgBurnToken
	MsgUpdateReserve = nftTypes.MsgUpdateReserve
	MsgSendToken     = nftTypes.MsgSendToken
	MsgUpdateToken   = nftTypes.MsgUpdateToken

	MsgCreateWallet      = multisigTypes.MsgCreateWallet
	MsgCreateTransaction = multisigTypes.MsgCreateTransaction
	MsgSignTransaction   = multisigTypes.MsgSignTransaction

	MsgInitializeSwap = swapTypes.MsgInitializeSwap
	MsgRedeemSwap     = swapTypes.MsgRedeemSwap
	MsgActivateChain  = swapTypes.MsgActivateChain

	MsgCreateValidator       = validatorTypes.MsgCreateValidator
	MsgEditValidator         = validatorTypes.MsgEditValidator
	MsgSetOnline             = validatorTypes.MsgSetOnline
	MsgSetOffline            = validatorTypes.MsgSetOffline
	MsgDelegate              = validatorTypes.MsgDelegate
	MsgDelegateNFT           = validatorTypes.MsgDelegateNFT
	MsgRedelegate            = validatorTypes.MsgRedelegate
	MsgRedelegateNFT         = validatorTypes.MsgRedelegateNFT
	MsgUndelegate            = validatorTypes.MsgUndelegate
	MsgUndelegateNFT         = validatorTypes.MsgUndelegateNFT
	MsgCancelRedelegation    = validatorTypes.MsgCancelRedelegation
	MsgCancelRedelegationNFT = validatorTypes.MsgCancelRedelegationNFT
	MsgCancelUndelegation    = validatorTypes.MsgCancelUndelegation
	MsgCancelUndelegationNFT = validatorTypes.MsgCancelUndelegationNFT

	Description = validatorTypes.Description

	MsgSoftwareUpgrade = upgradetypes.MsgSoftwareUpgrade
	MsgCancelUpgrade   = upgradetypes.MsgCancelUpgrade
	Plan               = upgradetypes.Plan
)

var (
	NewMsgCreateCoin    = coinTypes.NewMsgCreateCoin
	NewMsgUpdateCoin    = coinTypes.NewMsgUpdateCoin
	NewMsgMultiSendCoin = coinTypes.NewMsgMultiSendCoin
	NewMsgBuyCoin       = coinTypes.NewMsgBuyCoin
	NewMsgSellCoin      = coinTypes.NewMsgSellCoin
	NewMsgSellAllCoin   = coinTypes.NewMsgSellAllCoin
	NewMsgSendCoin      = coinTypes.NewMsgSendCoin
	NewMsgBurnCoin      = coinTypes.NewMsgBurnCoin
	NewMsgRedeemCheck   = coinTypes.NewMsgRedeemCheck

	NewMsgMintToken     = nftTypes.NewMsgMintToken
	NewMsgBurnToken     = nftTypes.NewMsgBurnToken
	NewMsgUpdateReserve = nftTypes.NewMsgUpdateReserve
	NewMsgSendToken     = nftTypes.NewMsgSendToken
	NewMsgUpdateToken   = nftTypes.NewMsgUpdateToken

	NewMsgUpdateCoinPrices = feeTypes.NewMsgUpdateCoinPrices

	NewMsgCreateWallet      = multisigTypes.NewMsgCreateWallet
	NewMsgCreateTransaction = multisigTypes.NewMsgCreateTransaction
	NewMsgSignTransaction   = multisigTypes.NewMsgSignTransaction

	NewMsgInitializeSwap = swapTypes.NewMsgInitializeSwap
	NewMsgRedeemSwap     = swapTypes.NewMsgRedeemSwap
	NewMsgActivateChain  = swapTypes.NewMsgActivateChain

	NewMsgCreateValidator       = validatorTypes.NewMsgCreateValidator
	NewMsgEditValidator         = validatorTypes.NewMsgEditValidator
	NewMsgSetOnline             = validatorTypes.NewMsgSetOnline
	NewMsgSetOffline            = validatorTypes.NewMsgSetOffline
	NewMsgDelegate              = validatorTypes.NewMsgDelegate
	NewMsgDelegateNFT           = validatorTypes.NewMsgDelegateNFT
	NewMsgRedelegate            = validatorTypes.NewMsgRedelegate
	NewMsgRedelegateNFT         = validatorTypes.NewMsgRedelegateNFT
	NewMsgUndelegate            = validatorTypes.NewMsgUndelegate
	NewMsgUndelegateNFT         = validatorTypes.NewMsgUndelegateNFT
	NewMsgCancelRedelegation    = validatorTypes.NewMsgCancelRedelegation
	NewMsgCancelRedelegationNFT = validatorTypes.NewMsgCancelRedelegationNFT
	NewMsgCancelUndelegation    = validatorTypes.NewMsgCancelUndelegation
	NewMsgCancelUndelegationNFT = validatorTypes.NewMsgCancelUndelegationNFT

	NewDescription = validatorTypes.NewDescription
)
