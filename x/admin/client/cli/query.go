package cli

import (
	"fmt"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/admin/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Aliases:                    []string{"register"},
		Short:                      "Querying commands for the rx-sharing admin module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	nameserviceQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryPatient(storeKey, cdc),
	)...)

	return nameserviceQueryCmd
}

// GetCmdQueryPatient queries information about a patient
func GetCmdQueryPatient(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "patient [pubkey]",
		Short: "query patient by public key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			pubkey := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/patient/%s", queryRoute, pubkey), nil)
			if err != nil {
				fmt.Printf("Could not fetch patient - %s \n", pubkey)
				return nil
			}

			var out types.Patient
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
	return cmd
}
