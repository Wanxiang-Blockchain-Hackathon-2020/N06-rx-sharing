package cli

import (
	"bufio"
	"fmt"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/crypto"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/admin/exported"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/patient/internal/types"
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
	faucetTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "patient transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	faucetTxCmd.AddCommand(flags.PostCommands(
		GetCmdAuthorize(cdc),
	)...)

	return faucetTxCmd
}

//GetCmdAuthorize is the CLI command for mining coin
func GetCmdAuthorize(cdc *codec.Codec) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "authorize",
		Short: " ",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			recipient := viper.GetString("recipient")
			id := viper.GetString("rx-id")

			cpt := viper.GetString("crypto")

			kb, err1 := keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), cmd.InOrStdin())
			if err1 != nil {
				return err1
			}

			priv, err2 := kb.ExportPrivateKeyObject(cliCtx.GetFromName(), "1234567890")
			if err2 != nil {
				return err2
			}

			patient, err3 := crypto.GenerateKey(cpt, priv.Bytes())
			if err3 != nil {
				return err3
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/rx/%s/%s", "patient", patient, id), nil)
			if err != nil {
				fmt.Printf("Does NOT found Rx %s in %s case history\n", id, patient)
				return nil
			}

			var rx exported.Rx
			cdc.MustUnmarshalJSON(res, &rx)

			res2, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/permits/%s", "patient", id), nil)
			if err != nil {
				fmt.Printf("Does NOT found permits for %s in %s case history\n", id, patient)
				return nil
			}

			var permits exported.RxPermission
			envelope := ""
			cdc.MustUnmarshalJSON(res2, &permits)
			for _, t := range permits {
				if t.Visitor == patient {
					envelope = t.Envelope
					break
				}
			}

			newone, _ := crypto.RenewEnvelope(cpt, priv.Bytes(), rx.Doctor, recipient, envelope)

			msg := types.NewMsgAuthorizeRx(cliCtx.GetFromAddress(), patient, recipient, id, newone)
			errs := msg.ValidateBasic()
			if errs != nil {
				return errs
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String("recipient", "", "name of drug store")
	cmd.Flags().String("rx-id", "", "id/number of Rx")
	cmd.Flags().String("crypto", "x25519", "algorithm used for encrypted the Rx data ")

	return cmd
}
