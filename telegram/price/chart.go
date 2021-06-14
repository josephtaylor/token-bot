package price

import (
	"fmt"
	"token-bot/config"
	"token-bot/uniswap"
)

type ChartItem struct {
}

func (c *ChartItem) GetMessage(_ *uniswap.Token) string {
	link := fmt.Sprintf("https://www.dextools.io/app/uniswap/pair-explorer/%s", config.App.PairAddress)
	return fmt.Sprintf("📈 *Chart:* [DexTools](%s)", link)
}
