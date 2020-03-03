package cli

import (
	"fmt"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/crypto"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/admin/exported"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/spf13/viper"

	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/patient/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the rx-sharing admin module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	nameserviceQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryRxs(storeKey, cdc),
		GetCmdQueryRx(storeKey, cdc),
		GetCmdQueryPermits(storeKey, cdc),
	)...)

	return nameserviceQueryCmd
}

// GetCmdQueryRxs queries information about rxs of a patient
func GetCmdQueryRxs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rxs",
		Short: "query rxs by public key",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			cpt := viper.GetString("crypto")
			keyname := viper.GetString("keyname")

			//kb, err1 := keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), cmd.InOrStdin())
			kb, err1 := keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), cmd.InOrStdin())
			if err1 != nil {
				return err1
			}

			priv, err2 := kb.ExportPrivateKeyObject(keyname, "1234567890")
			if err2 != nil {
				return err2
			}

			patient, err3 := crypto.GenerateKey(cpt, priv.Bytes())
			if err3 != nil {
				return err3
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/rxs/%s", queryRoute, patient), nil)
			if err != nil {
				fmt.Printf("Do NOT found any Rxs for patient: %s \n", patient)
				return nil
			}

			var out exported.CaseHistory
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}

	cmd.Flags().String("crypto", "x25519", "algorithm used for encrypted the rx data ")
	cmd.Flags().String("keyname", "", "specify user")

	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	viper.BindPFlag(flags.FlagKeyringBackend, cmd.Flags().Lookup(flags.FlagKeyringBackend))
	cmd.PersistentFlags().StringP(flags.FlagHome, "", "", "directory for config and data")

	cmd.MarkFlagRequired("keyname")

	return cmd
}

// GetCmdQueryRx queries information about rxs of a patient
func GetCmdQueryRx(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rx",
		Short: "query rx by id",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			fmt.Println("query:", queryRoute)

			cpt := viper.GetString("crypto")
			id := viper.GetString("rx-id")
			keyname := viper.GetString("keyname")
			decrypt := viper.GetBool("decrypt")

			kb, err1 := keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), cmd.InOrStdin())
			if err1 != nil {
				return err1
			}

			priv, err2 := kb.ExportPrivateKeyObject(keyname, "1234567890")
			if err2 != nil {
				return err2
			}

			patient, err3 := crypto.GenerateKey(cpt, priv.Bytes())
			if err3 != nil {
				return err3
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/rx/%s/%s", queryRoute, patient, id), nil)
			if err != nil {
				fmt.Printf("Does NOT found Rx %s in %s case history\n", id, patient)
				return nil
			}

			var rx exported.Rx
			cdc.MustUnmarshalJSON(res, &rx)

			if decrypt {

				res2, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/permits/%s", queryRoute, id), nil)
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

				plaintext, _ := crypto.Dencrypt(cpt, rx.Doctor, priv.Bytes(), envelope, rx.Encrypted)
				rx.Encrypted = plaintext
			}

			return cliCtx.PrintOutput(rx)
		},
	}

	cmd.Flags().String("crypto", "x25519", "algorithm used for encrypted the rx data ")
	cmd.Flags().String("keyname", "", "specify user")
	cmd.Flags().String("rx-id", "", "specifiy id of Rx")
	cmd.Flags().Bool("decrypt", false, "Set true to decrypt the encrypted content ")

	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	viper.BindPFlag(flags.FlagKeyringBackend, cmd.Flags().Lookup(flags.FlagKeyringBackend))
	cmd.PersistentFlags().StringP(flags.FlagHome, "", "", "directory for config and data")

	cmd.MarkFlagRequired("keyname")
	cmd.MarkFlagRequired("rx-id")

	return cmd
}

// GetCmdQueryPermits queries information about rxs of a patient
func GetCmdQueryPermits(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "permits",
		Short: "query permits of Rx by id",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			fmt.Println("query:", queryRoute)

			cpt := viper.GetString("crypto")
			id := viper.GetString("rx-id")
			keyname := viper.GetString("keyname")

			kb, err1 := keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), cmd.InOrStdin())
			if err1 != nil {
				return err1
			}

			priv, err2 := kb.ExportPrivateKeyObject(keyname, "1234567890")
			if err2 != nil {
				return err2
			}

			patient, err3 := crypto.GenerateKey(cpt, priv.Bytes())
			if err3 != nil {
				return err3
			}

			res2, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/permits/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("Does NOT found permits for %s in %s case history\n", id, patient)
				return nil
			}

			var permits exported.RxPermission
			cdc.MustUnmarshalJSON(res2, &permits)
			return cliCtx.PrintOutput(permits)
		},
	}

	cmd.Flags().String("crypto", "x25519", "algorithm used for encrypted the rx data ")
	cmd.Flags().String("keyname", "", "specify user")
	cmd.Flags().String("rx-id", "", "specifiy id of Rx")

	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	viper.BindPFlag(flags.FlagKeyringBackend, cmd.Flags().Lookup(flags.FlagKeyringBackend))
	cmd.PersistentFlags().StringP(flags.FlagHome, "", "", "directory for config and data")

	cmd.MarkFlagRequired("keyname")
	cmd.MarkFlagRequired("rx-id")

	return cmd
}
