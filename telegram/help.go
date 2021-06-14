package telegram

func SendHelpMessage(update *Update) (*Message, error) {
	message := "🤖 Token Bot 🤖\n\n" +
		"*Available Commands:*\n\n" +
		"`/price` \\- print price information\n" +
		"`/help` \\- print this help message\n" +
		"`/balance \\[address\\]` \\- look up current balance for an address\n" +
	    "`/about \\- information about this bot"
	return SendMessage(update, message)
}
