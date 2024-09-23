package api

import (
	"fmt"
	"strings"
)

// custom methods are in waiting for swagger

type resultGetNFT struct {
	Ok     bool `json:"ok"`
	Result struct {
		NftId         string `json:"nftId"`
		NftCollection string `json:"nftCollection"`
		Quantity      uint64 `json:"quantity"`
		CoinDenom     string `json:"coinDenom"`
		TokenURI      string `json:"tokenUri"`
		AllowMint     bool   `json:"allowMint"`
	} `json:"result"`
}

// /nfts/{nftId}
// Get nft by Id
func (api *API) GetNFT(nftId string) (*NFT, error) {

	var r = strings.NewReplacer(
		"{nftId}", fmt.Sprintf("%s", nftId),
	)
	var link = r.Replace("/nfts/{nftId}")

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetNFT{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetNFT(respValue)
}
