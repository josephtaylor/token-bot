package telegram

func SendAboutMessage(update *Update) (*Message, error) {
	message := "🤖 Token Bot 🤖\n\n" +
		"*Created by:* @blue\\_hunnids\n\n" +
		"*Source code:* [GitHub](https://github.com/josepthaylor/token-bot)\n" +
		"*Issue tracker:* [GitHub Issues](https://github.com/josephtaylor/token-bot/issues)\n\n" +
		"*Donations:*\nsend eth or tokens to the following address:\n" +
		"`0x5a61F59F41bE917129d12051F19d29B595452535`"
	return SendMessage(update, message)
}
