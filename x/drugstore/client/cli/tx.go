package cli

import (
	"bufio"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/crypto"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/drugstore/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetTxCmd return admin sub-command for tx
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "drugstore transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(flags.PostCommands(
		GetCmdSaleDrug(cdc),
	)...)

	return txCmd
}

//GetCmdSaleDrug is the CLI command for mining coin
func GetCmdSaleDrug(cdc *codec.Codec) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "sale",
		Short: "sale drugs to patient according the rx",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			patient := viper.GetString("patient")
			rx := viper.GetString("rx-id")
			cpt := viper.GetString("crypto")

			kb, err1 := keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), cmd.InOrStdin())
			if err1 != nil {
				return err1
			}

			priv, err2 := kb.ExportPrivateKeyObject(cliCtx.GetFromName(), "1234567890")
			if err2 != nil {
				return err2
			}

			drugstore, err3 := crypto.GenerateKey(cpt, priv.Bytes())
			if err3 != nil {
				return err3
			}

			msg := types.NewMsgSaleDrugs(cliCtx.GetFromAddress(), patient, drugstore, rx)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String("patient", "", "name of drug store")
	cmd.Flags().String("rx-id", "", "id/number of Rx")
	cmd.Flags().String("crypto", "x25519", "algorithm used for encrypted the Rx data ")

	cmd.MarkFlagRequired("rx-id")
	cmd.MarkFlagRequired("patient")

	return cmd
}
