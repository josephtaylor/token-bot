package price

import (
	"fmt"
	"strings"
	"token-bot/config"
	"token-bot/uniswap"
)

type TwitterItem struct {
}

func (t *TwitterItem) GetMessage(_ *uniswap.Token) string {
	return fmt.Sprintf("🐦 *Twitter:* [%s](%s)",
		config.App.Twitter,
		fmt.Sprintf("https://twitter.com/%s", strings.Replace(config.App.Twitter, "@", "", 1)))
}
