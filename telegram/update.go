package telegram

import (
	"fmt"
	"strconv"
	"strings"
	"token-bot/config"
	"token-bot/httpclient"
)

func NewTelegramClient() *httpclient.Client {
	return httpclient.NewClient(config.App.Telegram.BaseUri)
}

func GetUpdates(latestId int64) ([]Update, error) {
	client := NewTelegramClient()
	result := &UpdateResult{}
	err := client.Execute(
		client.GetRequest("getUpdates").
			Param("offset", strconv.FormatInt(latestId, 10)).
			Param("timeout", "1").
			Param("allowed_updates", "[\"message\"]"),
		result)
	if err != nil {
		return nil, fmt.Errorf("failed to get updates: %s", err)
	}
	if !result.OK {
		return nil, fmt.Errorf("error loading updates")
	}
	return result.Result, nil
}

func ProcessUpdate(update *Update) error {
	if nil == update || nil == update.Message || !containsBotCommand(update) {
		return nil
	}
	if strings.Contains(update.Message.Text, "/price") {
		_, err := SendPriceMessage(update)
		return err
	}
	if strings.Contains(update.Message.Text, "/help") {
		_, err := SendHelpMessage(update)
		return err
	}
	if strings.Contains(update.Message.Text, "/balance") {
		_, err := SendBalanceMessage(update)
		return err
	}
	// send unknown command message.
	return nil
}

func containsBotCommand(update *Update) bool {
	if len(update.Message.Entities) == 0 {
		return false
	}
	for _, entity := range update.Message.Entities {
		if entity.Type == "bot_command" {
			return true
		}
	}
	return false
}
