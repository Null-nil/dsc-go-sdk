package api

import (
	"fmt"

	"bitbucket.org/decimalteam/go-smart-node/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/encoding"
)

type TxDecoded struct {
	Msg  sdk.Msg  // transaction message
	Memo string   // transaction memo
	Fee  sdk.Coin // payed fee
	// error info
	Code      int    // error code, 0 - no errors
	Codespace string // codespace if code > 0
}

func decodeTransaction(txbytes []byte) (TxDecoded, error) {
	encodingConfig := encoding.MakeConfig(app.ModuleBasics)
	decoder := encodingConfig.TxConfig.TxDecoder()
	tx, err := decoder(txbytes)
	if err != nil {
		return TxDecoded{}, err
	}
	if len(tx.GetMsgs()) == 0 {
		return TxDecoded{}, fmt.Errorf("empty transaction")
	}
	var result TxDecoded
	result.Msg = tx.GetMsgs()[0]
	txmemo, ok := tx.(sdk.TxWithMemo)
	if ok {
		result.Memo = txmemo.GetMemo()
	}
	return result, nil
}
