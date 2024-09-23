package api

import (
	"fmt"

	"cosmossdk.io/math"
)

func converterGetCoin(resp resultGetCoin) (*CoinInfo, error) {
	var res *CoinInfo = &CoinInfo{}
	var ok bool

	res.Symbol = resp.Result.Symbol
	res.Title = resp.Result.Title
	res.CRR = resp.Result.Crr
	res.Volume, ok = math.NewIntFromString(resp.Result.Volume)
	if !ok {
		return nil, fmt.Errorf("cannot convert volume '%s' to math.Int", resp.Result.Volume)
	}
	res.Reserve, ok = math.NewIntFromString(resp.Result.Reserve)
	if !ok {
		return nil, fmt.Errorf("cannot convert reserve '%s' to math.Int", resp.Result.Reserve)
	}
	res.LimitVolume, ok = math.NewIntFromString(resp.Result.LimitVolume)
	if !ok {
		return nil, fmt.Errorf("cannot convert limit volume '%s' to math.Int", resp.Result.LimitVolume)
	}
	res.Creator = resp.Result.Creator

	return res, nil
}

func converterGetCoins(resp resultGetCoins) ([]CoinInfo, error) {
	var ok bool
	res := make([]CoinInfo, len(resp.Result.Coins))
	for i, coin := range resp.Result.Coins {
		res[i].Symbol = coin.Symbol
		res[i].Title = coin.Title
		res[i].CRR = coin.Crr
		res[i].Volume, ok = math.NewIntFromString(coin.Volume)
		if !ok {
			return nil, fmt.Errorf("cannot convert volume '%s' to math.Int", coin.Volume)
		}
		res[i].Reserve, ok = math.NewIntFromString(coin.Reserve)
		if !ok {
			return nil, fmt.Errorf("cannot convert reserve '%s' to math.Int", coin.Reserve)
		}
		res[i].LimitVolume, ok = math.NewIntFromString(coin.LimitVolume)
		if !ok {
			return nil, fmt.Errorf("cannot convert limit volume '%s' to math.Int", coin.LimitVolume)
		}
		res[i].Creator = coin.Creator
	}
	return res, nil
}

func converterGetAddress(resp resultGetAddress) (*AddressInfo, error) {
	var res AddressInfo
	res.Address = resp.Result.Address
	res.Type = resp.Result.Type
	res.Nonce = resp.Result.Nonce
	return &res, nil
}

func converterGetBlockByHeight(resp resultGetBlock) (*BlockInfo, error) {
	var res BlockInfo
	res.Height = resp.Result.Height
	res.Reward = resp.Result.Reward
	res.TxsCount = resp.Result.TxsCount
	res.EvmTxsCount = int64(resp.Result.EvmBlock.TransactionsCount)
	res.ValidatorsCount = resp.Result.ValidatorsCount
	return &res, nil
}

func converterGetNFTCollection(resp resultGetNFTCollection) (interface{}, error) {
	return resp.Result, nil
}

func converterGetTxByHash(resp resultGetTx) (*TxInfo, error) {
	var res TxInfo
	res.Block = resp.Result.BlockId
	//res.From = resp.Result.From
	//res.To = resp.Result.To
	res.Hash = resp.Result.Hash
	res.Status = resp.Result.Status
	res.Code = resp.Result.Code
	return &res, nil
}

func converterGetTxs(resp resultGetTxs) ([]TxInfo, error) {
	var res = make([]TxInfo, len(resp.Result.Txs))
	for i, tx := range resp.Result.Txs {
		res[i].Block = tx.BlockId
		//res[i].From = tx.From
		//res[i].To = tx.To
		res[i].Hash = tx.Hash
		res[i].Status = tx.Status
		res[i].Code = tx.Code
	}
	return res, nil
}

/*
func converterGetAddressTxs(resp resultGetAddressTxs) ([]TxInfo, error) {
	var res = make([]TxInfo, len(resp.Result.Txs))
	for i, tx := range resp.Result.Txs {
		res[i].Block = tx.BlockId
		res[i].From = tx.From
		res[i].To = tx.To
		res[i].Hash = tx.Hash
		res[i].Status = tx.Status
		res[i].Code = tx.Code
	}
	return res, nil
}
*/

func converterGetEvmContracts(resp resultGetEvmContracts) ([]EvmContract, error) {
	var res = make([]EvmContract, len(resp.Result.EvmContracts))
	for i, cntr := range resp.Result.EvmContracts {
		res[i].Address = cntr.Address
		res[i].Status = cntr.Status
		res[i].DeploymentEvmAccountAddress = cntr.DeploymentEvmAccountAddress
		res[i].DeploymentEvmTransactionHash = cntr.DeploymentEvmTransactionHash
	}
	return res, nil
}

func converterGetEvmTransactions(resp resultGetEvmTransactions) ([]EvmTransaction, error) {
	var res = make([]EvmTransaction, len(resp.Result.EvmTransactions))
	for i, tx := range resp.Result.EvmTransactions {
		res[i].Hash = tx.Hash
		res[i].Gas = tx.Gas
		res[i].Type = tx.Type
		res[i].Input = tx.Input
		res[i].Nonce = tx.Nonce
		res[i].Value = tx.Value
		res[i].ChainId = tx.ChainId
		res[i].GasPrice = tx.GasPrice
		res[i].MaxFeePerGas = tx.MaxFeePerGas
		res[i].From = tx.From
		res[i].To = tx.To
		res[i].EvmBlockHeight = tx.EvmBlockHeight
	}
	return res, nil
}

func converterGetEvmAccounts(resp resultGetEvmAccounts) ([]EvmAccount, error) {
	var res = make([]EvmAccount, len(resp.Result.EvmAccounts))
	for i, acc := range resp.Result.EvmAccounts {
		res[i].Address = acc.Address
	}
	return res, nil
}

/*
	func converterGetAllNFT(resp resultGetAllNFT) ([]NFT, error) {
		var err error
		var ok bool
		var res = make([]NFT, len(resp.Result.Nfts))
		for i, nft := range resp.Result.Nfts {
			res[i].NftCollection = nft.NftCollection
			res[i].NftId = nft.NftId
			res[i].Quantity, err = strconv.ParseUint(nft.Quantity, 10, 64)
			if err != nil {
				return nil, err
			}
			res[i].Reserve, ok = math.NewIntFromString(nft.Reserve)
			if !ok {
				return nil, fmt.Errorf("cannot convert reserve '%s' to math.Int", nft.Reserve)
			}
			res[i].Sender = nft.Sender
			res[i].Recipient = nft.Recipient
		}
		return res, nil
	}
*/
func converterGetValidatorsByKind(resp resultGetValidatorsByKind) ([]Validator, error) {
	var res = make([]Validator, len(resp.Result.Validators))
	for i, val := range resp.Result.Validators {
		res[i].Address = val.Address
	}
	return res, nil
}

func converterGetValidator(resp resultGetValidator) (*Validator, error) {
	return &Validator{Address: resp.Result.Address}, nil
}

func converterGetValidatorStakes(resp resultGetValidatorStakes) ([]ValidatorStake, error) {
	var ok bool
	var res = make([]ValidatorStake, len(resp.Result.Stakes))
	for i, stake := range resp.Result.Stakes {
		res[i].ValidatorId = stake.ValidatorId
		res[i].AddressId = stake.AddressId
		res[i].Amount, ok = math.NewIntFromString(stake.Amount)
		if !ok {
			return nil, fmt.Errorf("cannot convert amount '%s' to math.Int", stake.Amount)
		}
		res[i].CoinSymbol = stake.CoinSymbol
	}
	return res, nil
}

func converterGetValidatorStakesNFT(resp resultGetValidatorStakesNFT) ([]ValidatorStakeNFT, error) {
	var ok bool
	var res = make([]ValidatorStakeNFT, len(resp.Result.Stakes))
	for i, stake := range resp.Result.Stakes {
		res[i].AddressId = stake.AddressId
		res[i].BaseQuantity, ok = math.NewIntFromString(stake.BaseQuantity)
		if !ok {
			return nil, fmt.Errorf("cannot convert BaseQuantity '%s' to math.Int", stake.BaseQuantity)
		}
		res[i].NftId = stake.Nft.NftId
		res[i].NftCollection = stake.Nft.NftCollection
		res[i].Count = stake.Count
	}
	return res, nil
}

func converterGetEvmAccount(resp resultGetEvmAccount) (*EvmAccount, error) {
	return &EvmAccount{
		Address: resp.Result.Address,
	}, nil
}

func converterGetEvmContract(resp resultGetEvmContract) (*EvmContract, error) {
	return &EvmContract{
		Address:                      resp.Result.Address,
		Status:                       resp.Result.Status,
		DeploymentEvmAccountAddress:  resp.Result.DeploymentEvmAccountAddress,
		DeploymentEvmTransactionHash: resp.Result.DeploymentEvmTransactionHash,
	}, nil
}

func converterGetEvmTransaction(resp resultGetEvmTransaction) (*EvmTransaction, error) {
	return &EvmTransaction{
		Hash:           resp.Result.Hash,
		Gas:            resp.Result.Gas,
		Type:           resp.Result.Type,
		Input:          resp.Result.Input,
		Nonce:          resp.Result.Nonce,
		Value:          resp.Result.Value,
		ChainId:        resp.Result.ChainId,
		GasPrice:       resp.Result.GasPrice,
		MaxFeePerGas:   resp.Result.MaxFeePerGas,
		From:           resp.Result.From,
		To:             resp.Result.To,
		EvmBlockHeight: resp.Result.EvmBlockHeight,
	}, nil
}

func converterGetBlockTransactions(resp resultGetBlockTransactions) ([]TxInfo, error) {
	var res = make([]TxInfo, len(resp.Result.Txs))
	for i, tx := range resp.Result.Txs {
		res[i].Block = tx.BlockId
		//res[i].From = tx.From
		res[i].Hash = tx.Hash
		res[i].Status = tx.Status
		//res[i].To = tx.To
		res[i].Type = tx.Type
	}
	return res, nil
}

func converterGetEvmContractTransactions(resp resultGetEvmContractTransactions) ([]EvmTransaction, error) {
	var res = make([]EvmTransaction, len(resp.Result.EvmContractTransactions))
	for i, tx := range resp.Result.EvmContractTransactions {
		res[i].Hash = tx.Hash
		res[i].Gas = tx.Gas
		res[i].Type = tx.Type
		res[i].Input = tx.Input
		res[i].Nonce = tx.Nonce
		res[i].Value = tx.Value
		res[i].ChainId = tx.ChainId
		res[i].GasPrice = tx.GasPrice
		res[i].MaxFeePerGas = tx.MaxFeePerGas
		res[i].From = tx.From
		res[i].To = tx.To
		res[i].EvmBlockHeight = tx.EvmBlockHeight
	}
	return res, nil
}

/*
	func converterGetValidatorsCoins(resp resultGetValidatorsCoins) ([]ValidatorStakedCoin, error) {
		var res = make([]ValidatorStakedCoin, 0)
		for _, result := range resp.Result {
			for _, coin := range result.Stakes.Coins {
				amount, ok := math.NewIntFromString(coin.Amount)
				if !ok {
					return nil, fmt.Errorf("cannot convert amount '%s' to math.Int", coin.Amount)
				}
				baseAmount, ok := math.NewIntFromString(coin.BaseAmount)
				if !ok {
					return nil, fmt.Errorf("cannot convert base amount '%s' to math.Int", coin.BaseAmount)
				}
				res = append(res, ValidatorStakedCoin{
					Address:    result.Stakes.Address,
					CoinSymbol: coin.CoinSymbol,
					Amount:     amount,
					BaseAmount: baseAmount,
				})
			}
		}
		return res, nil
	}
*/
func converterGetEvmAccountBalances(resp resultGetEvmAccountBalances) ([]EvmAccountBalance, error) {
	var res = make([]EvmAccountBalance, 0)
	for _, balance := range resp.Result.EvmTokenAccountBalance {
		for _, erc20 := range balance.EvmAccountERC20TokenBalances {
			amount, ok := math.NewIntFromString(erc20.Amount)
			if !ok {
				return nil, fmt.Errorf("cannot convert amount '%s' to math.Int", erc20.Amount)
			}
			res = append(res, EvmAccountBalance{
				TokenType:    "ERC20",
				TokenAddress: erc20.EvmTokenAddress,
				Symbol:       erc20.EvmToken.Symbol,
				Amount:       amount,
			})
		}
		for _, erc721 := range balance.EvmAccountERC721TokenBalance {
			amount, ok := math.NewIntFromString(erc721.Amount)
			if !ok {
				return nil, fmt.Errorf("cannot convert amount '%s' to math.Int", erc721.Amount)
			}
			res = append(res, EvmAccountBalance{
				TokenType:    "ERC721",
				TokenAddress: erc721.EvmTokenAddress,
				Symbol:       erc721.EvmToken.Symbol,
				Amount:       amount,
			})
		}
		for _, erc1155 := range balance.EvmAccountERC1155TokenBalance {
			amount, ok := math.NewIntFromString(erc1155.Amount)
			if !ok {
				return nil, fmt.Errorf("cannot convert amount '%s' to math.Int", erc1155.Amount)
			}
			res = append(res, EvmAccountBalance{
				TokenType:    "ERC1155",
				TokenAddress: erc1155.EvmTokenAddress,
				Symbol:       erc1155.EvmToken.Symbol,
				Amount:       amount,
			})
		}
	}
	return res, nil
}

func converterGetEvmContractEvents(resp resultGetEvmContractEvents) ([]EvmEvent, error) {
	var res = make([]EvmEvent, len(resp.Result.EvmContractEvents))
	for i, event := range resp.Result.EvmContractEvents {
		res[i].EvmBlockHeight = event.EvmBlockHeight
		res[i].EvmTransactionHash = event.EvmTransactionHash
		res[i].GasUsed = event.GasUsed
		res[i].Type = event.Type
	}
	return res, nil
}

/*
func converterGetAddressStakes(resp resultGetAddressStakes) ([]ValidatorStake, error) {
	var res = make([]ValidatorStake, len(resp.Result.Stakes))
	for i, stake := range resp.Result.Stakes {
		res[i].AddressId = stake.AddressId
		amount, ok := math.NewIntFromString(stake.Amount)
		if !ok {
			return nil, fmt.Errorf("cannot convert amount '%s' to math.Int", stake.Amount)
		}
		res[i].Amount = amount
		res[i].CoinSymbol = stake.CoinSymbol
		res[i].ValidatorId = stake.ValidatorId
	}
	return res, nil
}
*/

func converterGetBlocks(resp resultGetBlocks) ([]BlockInfo, error) {
	var res = make([]BlockInfo, len(resp.Result.Blocks))
	for i, block := range resp.Result.Blocks {
		res[i].Height = block.Height
		res[i].Reward = block.Reward
		res[i].ValidatorsCount = block.ValidatorsCount
		res[i].TxsCount = block.TxsCount
	}
	return res, nil
}

/*
func converterGetAddressRewards(resp resultGetAddressRewards) ([]Reward, error) {
	var res = make([]Reward, len(resp.Result.Rewards))
	for i, reward := range resp.Result.Rewards {
		res[i].AddressId = reward.AddressId
		res[i].ValidatorId = reward.ValidatorId
		res[i].CoinSymbol = reward.CoinSymbol
		amount, ok := math.NewIntFromString(reward.Amount)
		if !ok {
			return nil, fmt.Errorf("cannot convert amount '%s' to math.Int", reward.Amount)
		}
		res[i].Amount = amount
	}
	return res, nil
}
*/

func converterGetNFTTransactions(resp resultGetNFTTransactions) ([]TxInfo, error) {
	var res = make([]TxInfo, len(resp.Result.Txs))
	for i, tx := range resp.Result.Txs {
		res[i].Hash = tx.Hash
		res[i].Block = tx.BlockId
		//res[i].From = tx.From
		res[i].Status = tx.Status
		//res[i].To = tx.To
		res[i].Type = tx.Type
	}
	return res, nil
}

func converterGetNFT(resp resultGetNFT) (*NFT, error) {
	var res = &NFT{}

	res.NftCollection = resp.Result.NftCollection
	res.NftId = resp.Result.NftId
	res.Quantity = resp.Result.Quantity
	res.TokenURI = resp.Result.TokenURI
	res.AllowMint = resp.Result.AllowMint

	return res, nil
}
