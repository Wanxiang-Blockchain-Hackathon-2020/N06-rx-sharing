package cli

import (
	"bufio"
	"fmt"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/admin/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
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

//GetCmdWithdraw is the CLI command for mining coin
func GetCmdMint(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "prescribe",
		Short: "mint coin to sender address",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			input.GetPassword("input password", inBuf)

			encrypted := ""
			envelope := ""
			memo := ""

			msg := types.NewMsgPrescribe(cliCtx.GetFromAddress(), cliCtx.GetFromAddress(), encrypted, envelope, memo)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
