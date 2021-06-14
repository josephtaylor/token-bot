package telegram

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"token-bot/etherscan"
	"token-bot/uniswap"
)

func SendBalanceMessage(update *Update) (*Message, error) {
	token, err := uniswap.GetDefaultToken()
	var message string
	if err != nil {
		message = fmt.Sprintf("error: failed to load token: %s", err)
		return SendMessage(update, message)
	}
	address := parseAddress(update)
	if "" == address {
		message = fmt.Sprintf("error: no wallet address provided")
		return SendMessage(update, message)
	}
	balance, err := etherscan.GetBalance(address, token.ID)
	if err != nil {
		message = fmt.Sprintf("error: failed to load balance: %s", err)
		return SendMessage(update, message)
	}
	decimals, _ := strconv.Atoi(token.Decimals)
	message = fmt.Sprintf(`💰💰 [Wallet](%s) Balance 💰💰

*%s:* %f
*eth:* %0.6f
*usd:* $%.2f`,
		fmt.Sprintf("https://etherscan.io/address/%s", address),
		token.Symbol,
		balance/math.Pow(10, float64(decimals)),
		token.EthPrice*balance/math.Pow(10, float64(decimals)),
		balance*token.PriceFloat/math.Pow(10, float64(decimals)),
	)
	return SendMessage(update, message)
}

func parseAddress(update *Update) string {
	text := update.Message.Text
	text = strings.Replace(text, "/balance", "", -1)
	text = strings.Replace(text, "@tokenBot", "", -1)
	text = strings.TrimSpace(text)
	return text
}
