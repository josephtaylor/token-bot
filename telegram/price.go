package telegram

import (
	"fmt"
	"strings"
	"token-bot/telegram/price"
	"token-bot/uniswap"
)

func SendPriceMessage(update *Update) (*Message, error) {
	token, err := uniswap.GetDefaultToken()
	var message string
	if err != nil {
		message = fmt.Sprintf("failed to load token: %s", err)
	} else {
		message = fmt.Sprintf("🚀🚀 %s *%s* 🚀🚀\n\n%s",
			token.Name,
			token.Symbol,
			strings.Join(toMessages(price.GetItems(), token), "\n"),
		)
	}
	return SendMessage(update, message)
}

func toMessages(items []price.Item, token *uniswap.Token) []string {
	messages := make([]string, 0)
	for _, item := range items {
		message := item.GetMessage(token)
		if "" != message {
			messages = append(messages, message)
		}
	}
	return messages
}
