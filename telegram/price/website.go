package price

import (
	"fmt"
	"token-bot/config"
	"token-bot/uniswap"
)

type WebsiteItem struct {
}

func (w *WebsiteItem) GetMessage(_ *uniswap.Token) string {
	return fmt.Sprintf("🌐 *Website:* [%s](%s)",
		config.App.Website,
		fmt.Sprintf("https://%s", config.App.Website))
}
