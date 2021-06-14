package telegram

func SendHelpMessage(update *Update) (*Message, error) {
	message := "🤖 tokenBot 🤖\n\n" +
		"*Available Commands:*\n\n" +
		"`/price` \\- print price information\n" +
		"`/help` \\- print this help message\n" +
		"`/balance \\[address\\]` \\- look up current balance for an address"
	return SendMessage(update, message)
}
