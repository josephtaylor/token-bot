package uniswap

type Pair struct {
	Price       float64
	TotalSupply float64
	Token       Token
}

type Token struct {
	ID         string  `json:"id"`
	Symbol     string  `json:"symbol"`
	Name       string  `json:"name"`
	Price      string  `json:"price"`
	EthPrice   float64 `json:"ethPrice"`
	PriceFloat float64 `json:"priceFloat"`
	MarketCap  string  `json:"marketCap"`
	Decimals   string  `json:"decimals"`
}

type Response struct {
	Pair struct {
		Token1Price string `json:"token1Price"`
		TotalSupply string `json:"totalSupply"`
		Token       Token  `json:"token0"`
	} `json:"pair"`
}
