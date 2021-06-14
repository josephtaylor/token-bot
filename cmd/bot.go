package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"token-bot/logging"
	"token-bot/telegram"
)

func newBotCmd() *cobra.Command {
	var debug *bool
	cmd := &cobra.Command{
		Use:   "bot",
		Short: "Start the token telegram bot",
		Run: func(cmd *cobra.Command, args []string) {
			if *debug {
				logging.ConfigureLogging(true)
				logrus.Debug("debug logging enabled")
			}
			logrus.Info("initializing token telegram bot")
			err := telegram.Process()
			if err != nil {
				logrus.Fatal(err)
			}
		},
	}
	flags := cmd.Flags()
	debug = flags.BoolP("debug", "d", false, "enable debug logging")
	return cmd
}

func init() {
	rootCmd.AddCommand(newBotCmd())
}
