package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "token",
	Short: "Token Bot",
	Run:   showHelp,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func showHelp(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		logrus.Fatal(err)
	}
}
