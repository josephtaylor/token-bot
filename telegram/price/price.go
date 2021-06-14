package price

import (
	"fmt"
	"token-bot/uniswap"
)

type PriceItem struct {
}

func (p *PriceItem) GetMessage(token *uniswap.Token) string {
	return fmt.Sprintf("💰 *Price: *$%s", token.Price)
}
