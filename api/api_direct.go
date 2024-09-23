package api

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/go-resty/resty/v2"
)

// API is a struct implementing DSC API iteraction.
type DirectAPI struct {
	rpc  *resty.Client
	rest *resty.Client

	// Parameters
	chainID   string
	baseDenom string
}

// NewDirectAPI creates Decimal API instance for direct interaction with node
func NewDirectAPI(nodeHost string) *DirectAPI {
	return NewDirectAPIWithPorts(nodeHost, 26657, 1317)
}

// NewDirectAPI creates Decimal API instance for direct interaction with node
func NewDirectAPIWithPorts(nodeHost string, tendermintPort int, restPort int) *DirectAPI {
	initConfig()
	return &DirectAPI{
		rpc:  resty.New().SetBaseURL(fmt.Sprintf("http://%s:%d", nodeHost, tendermintPort)).SetTimeout(time.Minute),
		rest: resty.New().SetBaseURL(fmt.Sprintf("http://%s:%d", nodeHost, restPort)).SetTimeout(time.Minute),
	}
}

func (api *DirectAPI) GetParameters() error {
	type respDirectGenesis struct {
		Result struct {
			Genesis struct {
				ChainID  string `json:"chain_id"`
				AppState struct {
					Coin struct {
						Params struct {
							BaseDenom string `json:"base_denom"`
						} `json:"params"`
					} `json:"coin"`
				} `json:"app_state"`
			} `json:"genesis"`
		} `json:"result"`
	}
	// request
	res, err := api.rpc.R().Get("/genesis")
	if err = processConnectionError(res, err); err != nil {
		return err
	}
	// json decode
	respValue := respDirectGenesis{}
	err = universalJSONDecode(res.Body(), &respValue, nil, func() (bool, bool) {
		return respValue.Result.Genesis.ChainID > "", false
	})
	if err != nil {
		return err
	}
	// process results
	api.chainID = respValue.Result.Genesis.ChainID
	api.baseDenom = respValue.Result.Genesis.AppState.Coin.Params.BaseDenom
	return nil
}

// Address requests full information about specified address
func (api *DirectAPI) GetAccountNumberAndSequence(address string) (uint64, uint64, error) {
	type respDirectAddress struct {
		Account struct {
			BaseAccount struct {
				Number   string `json:"account_number"`
				Sequence string `json:"sequence"`
			} `json:"base_account"`
		} `json:"account"`
	}
	// request
	res, err := api.rest.R().Get(fmt.Sprintf("/cosmos/auth/v1beta1/accounts/%s", address))
	if err = processConnectionError(res, err); err != nil {
		return 0, 0, err
	}
	// json decode
	respValue, respErr := respDirectAddress{}, RestError{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Account.BaseAccount.Number > "", respErr.Code != 0
	})
	//empty account (no transactions), it's normal
	if err != nil {
		return 0, 0, joinErrors(err, respErr)
	}
	// process result
	accNumber, _ := strconv.ParseUint(respValue.Account.BaseAccount.Number, 10, 64)
	seq, _ := strconv.ParseUint(respValue.Account.BaseAccount.Sequence, 10, 64)

	return accNumber, seq, nil
}

// Address requests full information about specified address
func (api *DirectAPI) GetAccountBalance(address string) (sdk.Coins, error) {
	type respBalance struct {
		Balances sdk.Coins `json:"balances"`
	}
	// request
	res, err := api.rest.R().Get(fmt.Sprintf("/cosmos/bank/v1beta1/balances/%s", address))
	if err = processConnectionError(res, err); err != nil {
		return sdk.NewCoins(), err
	}
	// json decode
	respValue, respErr := respBalance{}, RestError{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return len(respValue.Balances) > 0, respErr.Code != 0
	})
	if err != nil {
		return sdk.NewCoins(), joinErrors(err, respErr)
	}
	// process result
	return respValue.Balances, nil
}

// BaseCoin() returns base coin symbol from genesis. Need for correct transaction building
func (api *DirectAPI) BaseCoin() string {
	return api.baseDenom
}

// ChainID() returns blockchain network chain id
func (api *DirectAPI) ChainID() string {
	return api.chainID
}

func (api *DirectAPI) BroadcastTxSync(data []byte) (*TxResponse, error) {
	type directSyncResponse struct {
		Result struct {
			Code      int    `json:"code"`
			Log       string `json:"log"`
			Codespace string `json:"codespace"`
			Hash      string `json:"hash"`
		} `json:"result"`
	}
	// request
	res, err := api.rpc.R().Get("/broadcast_tx_sync?tx=0x" + hex.EncodeToString(data))
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue := directSyncResponse{}
	err = universalJSONDecode(res.Body(), &respValue, nil, func() (bool, bool) {
		return respValue.Result.Hash > "", false
	})
	if err != nil {
		return nil, err
	}
	// process result
	return &TxResponse{
		Code:      respValue.Result.Code,
		Hash:      respValue.Result.Hash,
		Log:       respValue.Result.Log,
		Codespace: respValue.Result.Codespace,
	}, nil
}

func (api *DirectAPI) BroadcastTxCommit(data []byte) (*TxResponse, error) {
	type directCommitResponse struct {
		Result struct {
			CheckTx struct {
				Code      int    `json:"code"`
				Log       string `json:"log"`
				Codespace string `json:"codespace"`
			} `json:"check_tx"`
			DeliverTx struct {
				Code      int    `json:"code"`
				Log       string `json:"log"`
				Codespace string `json:"codespace"`
			} `json:"deliver_tx"`
			Hash string `json:"hash"`
		} `json:"result"`
	}
	// request
	res, err := api.rpc.R().Get("/broadcast_tx_commit?tx=0x" + hex.EncodeToString(data))
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue := directCommitResponse{}
	err = universalJSONDecode(res.Body(), &respValue, nil, func() (bool, bool) {
		return respValue.Result.Hash > "", false
	})
	if err != nil {
		return nil, err
	}
	if respValue.Result.CheckTx.Code != 0 {
		return &TxResponse{
			Code:      respValue.Result.CheckTx.Code,
			Hash:      respValue.Result.Hash,
			Log:       respValue.Result.CheckTx.Log,
			Codespace: respValue.Result.CheckTx.Codespace,
		}, nil
	}
	// process result
	return &TxResponse{
		Code:      respValue.Result.DeliverTx.Code,
		Hash:      respValue.Result.Hash,
		Log:       respValue.Result.DeliverTx.Log,
		Codespace: respValue.Result.DeliverTx.Codespace,
	}, nil
}

func (api *DirectAPI) CalculateFee(data []byte, denom string) (sdk.Coin, error) {
	type calculateResponse struct {
		Commission string `json:"commission"`
	}
	// request
	res, err := api.rest.R().Get(fmt.Sprintf("/decimal/fee/v1/calculate_commission?tx_bytes=%s&denom=%s", hex.EncodeToString(data), denom))
	if err = processConnectionError(res, err); err != nil {
		return sdk.Coin{}, err
	}
	// json decode
	respValue, respErr := calculateResponse{}, RestError{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Commission > "", respErr.Code != 0
	})
	if err != nil {
		return sdk.Coin{}, err
	}
	// process result
	amount, ok := sdk.NewIntFromString(respValue.Commission)
	if !ok {
		return sdk.Coin{}, fmt.Errorf("can't convert commission to int")
	}
	return sdk.NewCoin(denom, amount), nil
}
