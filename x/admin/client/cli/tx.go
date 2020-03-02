package cli

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/crypto"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/admin/internal/types"
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
	"time"
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
		GetRegisterDoctor(cdc),
		GetRegisterDrugstore(cdc),
		GetRegisterPatient(cdc),
	)...)

	return faucetTxCmd
}

//GetCmdWithdraw is the CLI command for mining coin
func GetTestCommand(cdc *codec.Codec) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "keygen",
		Short: "keygen command",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			//inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			cpt := viper.GetString("crypto")

			kb, err := keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), cmd.InOrStdin())
			if err != nil {
				return err
			}

			priv, err2 := kb.ExportPrivateKeyObject(cliCtx.GetFromName(), "1234567890")
			if err2 != nil {
				return err2
			}

			pubkey, err3 := crypto.GenerateKey(cpt, priv.Bytes())
			if err3 != nil {
				return err3
			}

			fmt.Println(pubkey)

			return nil
		},
	}

	cmd.Flags().String("crypto", "", "specify the algorithm to crypto, available value[sm2, ed25519]")

	return cmd
}

//GetRegisterDoctor is the CLI command for register doctor
func GetRegisterDoctor(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "doctor --name 张三 --pubkey bjSOrqv/lVbwYwnUB+HTZybvxCc2CBl2w012W6uhZj0= ",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			name := viper.GetString("name")
			gender := viper.GetString("gender")
			hospital := viper.GetString("hospital")
			department := viper.GetString("department")
			title := viper.GetString("title")
			introduction := viper.GetString("introduction")

			pubkey := viper.GetString("pubkey")

			msg := types.NewMsgRegisterDocter(cliCtx.GetFromAddress(), pubkey, name, gender, hospital, department, title, introduction)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String("name", "", "name of doctor")
	cmd.Flags().String("gender", "", "gender of doctor")
	cmd.Flags().String("hospital", "", "hospital of doctor")
	cmd.Flags().String("department", "", "department of doctor")
	cmd.Flags().String("title", "", "title of doctor")
	cmd.Flags().String("introduction", "", "introduction of doctor")
	cmd.Flags().String("pubkey", "", "address of doctor")

	return cmd
}

//GetRegisterPatient is the CLI command for register doctor
func GetRegisterPatient(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "patient",
		Short: "patient --name 张三 --pubkey bjSOrqv/lVbwYwnUB+HTZybvxCc2CBl2w012W6uhZj0= ",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			name := viper.GetString("name")
			gender := viper.GetString("gender")
			birthday, errb := time.Parse(time.RFC3339, viper.GetString("birthday")+"T00:00:00Z")
			if errb != nil {
				return types.ErrInputInvalid
			}

			pubkey := viper.GetString("pubkey")

			//加密级别：一般
			json := viper.GetString("other")
			encrypted := base64.StdEncoding.EncodeToString([]byte(json))
			envelope := ""

			msg := types.NewMsgRegisterPatient(cliCtx.GetFromAddress(), pubkey, name, gender, birthday, encrypted, envelope)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String("name", "", "name of patient")
	cmd.Flags().String("gender", "", "gender of patient")
	cmd.Flags().String("birthday", "", "hospital of patient")
	cmd.Flags().String("other", "", "department of patient")
	cmd.Flags().String("pubkey", "", "pubkey of patient")

	return cmd
}

//GetRegisterDrugstore is the CLI command for register doctor
func GetRegisterDrugstore(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "drugstore",
		Short: "drugstore --name 张三 --pubkey bjSOrqv/lVbwYwnUB+HTZybvxCc2CBl2w012W6uhZj0= ",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			name := viper.GetString("name")
			phone := viper.GetString("phone")
			group := viper.GetString("group")
			biztime := viper.GetString("biztime")
			location := viper.GetString("location")

			pubkey := viper.GetString("pubkey")

			msg := types.NewMsgRegisterDrugStore(cliCtx.GetFromAddress(), pubkey, name, phone, group, biztime, location)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String("name", "", "name of drug store")
	cmd.Flags().String("phone", "", "specify a business phone number")
	cmd.Flags().String("group", "", "group of drug store")
	cmd.Flags().String("biztime", "", "business time of drug store")
	cmd.Flags().String("location", "", "address of drug store")
	cmd.Flags().String("pubkey", "", "pubkey of drug store")

	return cmd
}
