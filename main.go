package main

import (
	"token-bot/cmd"
	"token-bot/logging"
)

func main() {
	logging.ConfigureLogging(false)
	cmd.Execute()
}
