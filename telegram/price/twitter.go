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
	twitter := strings.ReplaceAll(config.App.Twitter, "_", "\\_")
	return fmt.Sprintf("🐦 *Twitter:* [%s](%s)",
		twitter,
		fmt.Sprintf("https://twitter.com/%s", strings.Replace(twitter, "@", "", 1)))
}
