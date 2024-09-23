package api

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

var (
	bigE15 = new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil)
	bigE18 = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	sdkE15 = sdk.NewIntFromBigInt(bigE15)
	sdkE18 = sdk.NewIntFromBigInt(bigE18)
)

func BipToPip(bip sdk.Int) sdk.Int {
	return bip.Mul(sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)))
}

// EtherToWei convert number 1 to 1 * 10^18
func EtherToWei(ether sdk.Int) sdk.Int {
	return ether.Mul(sdkE18)
}

// FinneyToWei convert number 1 to 1 * 10^15
func FinneyToWei(finney sdk.Int) sdk.Int {
	return finney.Mul(sdkE15)
}

// WeiToFinney convert 1 * 10^15 to 1
func WeiToFinney(wei sdk.Int) sdk.Int {
	return wei.Quo(sdkE15)
}

// WeiToEther convert 1 * 10^18 to 1
func WeiToEther(wei sdk.Int) sdk.Int {
	return wei.Quo(sdkE18)
}

// GetDecimalAddressFromBech32 returns the sdk.Account address of given address, while
// also changing bech32 human readable prefix (HRP) to the value set on the global sdk.Config (eg: `dx`).
// The function fails if the provided bech32 address is invalid.
func GetDecimalAddressFromBech32(address string) (sdk.AccAddress, error) {

	addressBz, err := sdk.GetFromBech32(address, "d0")
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress
	}

	// safety check: shouldn't happen
	if err := sdk.VerifyAddressFormat(addressBz); err != nil {
		return nil, err
	}

	return sdk.AccAddress(addressBz), nil
}

// GetDecimalAddressFromHex returns the sdk.Account address of given address.
// The function fails if the provided hex address is invalid or does not start with 0x.
func GetDecimalAddressFromHex(address string) (sdk.AccAddress, error) {
	addressBz, err := hexutil.Decode(address)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress
	}

	// safety check: shouldn't happen
	if err := sdk.VerifyAddressFormat(addressBz); err != nil {
		return nil, err
	}

	return sdk.AccAddress(addressBz), nil
}
