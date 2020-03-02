package cli

import (
	"bufio"
	"fmt"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/patient/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
)

// GetTxCmd return admin sub-command for tx
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	faucetTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "admin transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	faucetTxCmd.AddCommand(flags.PostCommands(
		GetTestCommand(cdc),
		//GetCmdMintFor(cdc),
	)...)

	return faucetTxCmd
}

//GetCmdWithdraw is the CLI command for mining coin
func GetTestCommand(cdc *codec.Codec) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "test",
		Short: "test command",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())

			test := viper.GetString("test")

			fmt.Println(test)
			//cliCtx := context.NewCLIContext().WithCodec(cdc)
			//
			//txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			passwd, _ := input.GetPassword("input password", inBuf)

			fmt.Println("==============>", passwd)

			return nil
		},
	}

	cmd.Flags().String(cli.HomeFlag, "", "node's home directory")
	cmd.Flags().String("test", "", "test")

	return cmd
}
