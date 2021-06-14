package price

import (
	"fmt"
	"token-bot/uniswap"
)

type BuyItem struct {
}

func (b *BuyItem) GetMessage(token *uniswap.Token) string {
	link := fmt.Sprintf("https://app.uniswap.org/#/swap?outputCurrency=%s", token.ID)
	return fmt.Sprintf("🔀 *Buy:* [Uniswap](%s)", link)
}
