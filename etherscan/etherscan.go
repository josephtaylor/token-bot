package etherscan

import (
	"strconv"
	"token-bot/config"
	"token-bot/httpclient"
)

type BalanceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func GetBalance(address string, contractAddress string) (float64, error) {
	client := httpclient.NewClient(config.App.Etherscan.BaseUri)
	response := &BalanceResponse{}
	err := client.Execute(
		client.GetRequest().
			Param("module", "account").
			Param("action", "tokenbalance").
			Param("contractAddress", contractAddress).
			Param("address", address).
			Param("tag", "latest").
			Param("apiKey", config.App.Etherscan.ApiKey),
		response)
	if err != nil {
		return 0, err
	}
	balance, _ := strconv.ParseFloat(response.Result, 64)
	return balance, nil
}
