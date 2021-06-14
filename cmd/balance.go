package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math"
	"strconv"
	"token-bot/etherscan"
	"token-bot/uniswap"
)

func newBalanceCmd() *cobra.Command {
	var address *string
	cmd := &cobra.Command{
		Use:   "balance",
		Short: "Get a wallet balance",
		Run: func(cmd *cobra.Command, args []string) {
			token, err := uniswap.GetDefaultToken()
			if err != nil {
				logrus.Fatalf("failed to load token: %s", err)
				return
			}
			balance, err := etherscan.GetBalance(*address, token.ID)
			decimals, _ := strconv.Atoi(token.Decimals)
			ownedAmount := balance * token.PriceFloat / math.Pow(10.0, float64(decimals))
			fmt.Printf("Current %s [ %s ] Balance: %.25f\n", token.Name, token.Symbol, ownedAmount)
		},
	}
	flags := cmd.Flags()
	address = flags.StringP("address", "a", "", "the wallet address")
	_ = cobra.MarkFlagRequired(flags, "address")
	return cmd
}

func init() {
	rootCmd.AddCommand(newBalanceCmd())
}
