package uniswap

import (
	"context"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/machinebox/graphql"
	"math"
	"strconv"
	"token-bot/config"
)

func GetToken(pairId string, totalTokens float64) (*Token, error) {
	pair, err := getPair(pairId)
	if err != nil {
		return nil, err
	}
	ethPrice, err := getEthPrice()
	if err != nil {
		return nil, err
	}
	pair.Token.EthPrice = pair.Price
	pair.Token.Price = printPrice(pair.Price * ethPrice)
	pair.Token.PriceFloat = pair.Price * ethPrice
	pair.Token.MarketCap = printMarketCap(pair.Price * ethPrice * totalTokens)
	return &pair.Token, nil
}

func GetDefaultToken() (*Token, error) {
	return GetToken(config.App.PairAddress, config.App.TotalTokens)
}

func printPrice(value float64) string {
	return fmt.Sprintf("%.14f", math.Round(value*100000000000000)/100000000000000)
}

func printMarketCap(value float64) string {
	return fmt.Sprintf("$%s", humanize.Comma(int64(value)))
}

func getPair(pairId string) (*Pair, error) {
	client := graphql.NewClient(config.App.Uniswap.BaseUri)

	req := graphql.NewRequest(buildPriceQuery(pairId))

	ctx := context.Background()
	response := &Response{}

	if err := client.Run(ctx, req, response); err != nil {
		return nil, fmt.Errorf("failed to load price %s", err)
	}
	priceString := response.Pair.Token1Price
	price, _ := strconv.ParseFloat(priceString, 64)
	totalSupplyString := response.Pair.TotalSupply
	totalSupply, _ := strconv.ParseFloat(totalSupplyString, 64)
	return &Pair{
		Price:       price,
		TotalSupply: totalSupply,
		Token:       response.Pair.Token,
	}, nil
}

func getEthPrice() (float64, error) {
	client := graphql.NewClient(config.App.Uniswap.BaseUri)

	req := graphql.NewRequest("{ bundle(id:\"1\") { ethPrice } }")

	ctx := context.Background()
	var response map[string]interface{}

	if err := client.Run(ctx, req, &response); err != nil {
		return 0, fmt.Errorf("failed to load price %s", err)
	}
	priceString := response["bundle"].(map[string]interface{})["ethPrice"].(string)
	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse price %s", err)
	}
	return price, nil
}

func buildPriceQuery(pairId string) string {
	return fmt.Sprintf(`{ 
		pair(id:"%s") {
			token1Price 
			totalSupply
			token0 {
				id
				symbol
				name
				decimals
			}
		} 
	}`, pairId)
}
