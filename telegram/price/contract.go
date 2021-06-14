package price

import (
	"fmt"
	"token-bot/uniswap"
)

type ContractItem struct {
}

func (b *ContractItem) GetMessage(token *uniswap.Token) string {
	return fmt.Sprintf("📄 *Contract:* [%s](%s)",
		token.ID,
		fmt.Sprintf("https://etherscan.io/address/%s#code", token.ID))
}
