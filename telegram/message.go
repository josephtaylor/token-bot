package telegram

import (
	"strconv"
	"strings"
)

func SendMessage(update *Update, message string) (*Message, error) {
	client := NewTelegramClient()
	id := update.Message.Chat.ID
	response := &Message{}
	message = strings.Replace(message, ".", "\\.", -1)
	err := client.Execute(
		client.GetRequest("sendMessage").
			Param("chat_id", strconv.FormatInt(id, 10)).
			Param("text", message).
			Param("parse_mode", "MarkdownV2").
			Param("disable_web_page_preview", "true"),
		response)
	return response, err
}
