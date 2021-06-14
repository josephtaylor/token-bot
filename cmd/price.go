package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"token-bot/uniswap"
)

func newPriceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "price",
		Short: "Get the current price",
		Run: func(cmd *cobra.Command, args []string) {
			token, err := uniswap.GetDefaultToken()
			if err != nil {
				logrus.Fatalf("failed to load token: %s", err)
				return
			}
			fmt.Printf("%s [ %s ] Price: %s\n", token.Name, token.Symbol, token.Price)
		},
	}
	return cmd
}

func init() {
	rootCmd.AddCommand(newPriceCmd())
}
