package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/line/lbm-sdk/client"
	"github.com/line/lbm-sdk/client/flags"
	"github.com/line/lbm-sdk/client/tx"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

const (
	FlagUri     = "uri"
	FlagUriHash = "uri-hash"
)

// NewTxCmd returns the transaction commands for the module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        composable.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", composable.ModuleName),
		Long:                       "manipulate composable nfts",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		// NewTxCmdUpdateParams(),
		NewTxCmdSend(),
		NewTxCmdAttach(),
		NewTxCmdDetach(),
		NewTxCmdNewClass(),
		NewTxCmdUpdateClass(),
		NewTxCmdMintNFT(),
		NewTxCmdBurnNFT(),
		NewTxCmdUpdateNFT(),
	)

	return txCmd
}

// func NewTxCmdUpdateParams() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "update-params [authority] [params-json]",
// 		Args:  cobra.ExactArgs(2),
// 		Short: "Update the module parameters",
// 		Example: `
// Example of the content of params-json:

// {
//   "max_descendants": 15
// }
// `,
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			if err := validateGenerateOnly(cmd); err != nil {
// 				return err
// 			}

// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			params, err := parseParams(clientCtx.Codec, args[1])
// 			if err != nil {
// 				return err
// 			}

// 			msg := composable.MsgUpdateParams{
// 				Authority: args[0],
// 				Params:    *params,
// 			}
// 			if err := msg.ValidateBasic(); err != nil {
// 				return err
// 			}

// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)

// 	return cmd
// }

func NewTxCmdSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [sender] [recipient] [id]",
		Args:  cobra.ExactArgs(3),
		Short: "Send an nft from one account to another account",
		Example: `
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			sender := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, sender); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			nft, err := ParseNFT(args[2])
			if err != nil {
				return err
			}

			msg := composable.MsgSend{
				Sender:    sender,
				Recipient: args[1],
				Nft:       *nft,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdAttach() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "attach [owner] [subject-id] [target-id]",
		Args:    cobra.ExactArgs(3),
		Short:   "Attach a root nft to another nft",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subject, err := ParseNFT(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "subject")
			}

			target, err := ParseNFT(args[2])
			if err != nil {
				return sdkerrors.Wrap(err, "target")
			}

			msg := composable.MsgAttach{
				Owner:   owner,
				Subject: *subject,
				Target:  *target,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdDetach() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "detach [owner] [id]",
		Args:    cobra.ExactArgs(2),
		Short:   "Detach an nft from another nft",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			nft, err := ParseNFT(args[1])
			if err != nil {
				return err
			}

			msg := composable.MsgDetach{
				Owner: owner,
				Nft:   *nft,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdNewClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "new-class [owner] [--uri] [--uri-hash]",
		Args:    cobra.ExactArgs(1),
		Short:   "create a class",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// traits

			msg := composable.MsgNewClass{
				Owner: owner,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdUpdateClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-class [class-id] [--uri] [--uri-hash]",
		Args:    cobra.ExactArgs(1),
		Short:   "update a class",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			classID := args[0]
			if err := composable.ValidateClassID(classID); err != nil {
				return err
			}

			owner := composable.ClassOwner(classID).String()
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := composable.MsgUpdateClass{
				ClassId: classID,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdMintNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "mint-nft [class-id] [recipient] [--uri] [--uri-hash]",
		Args:    cobra.ExactArgs(2),
		Short:   "mint an nft",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			classID := args[0]
			if err := composable.ValidateClassID(classID); err != nil {
				return err
			}

			owner := composable.ClassOwner(classID).String()
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := composable.MsgMintNFT{
				ClassId:   classID,
				Recipient: args[1],
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdBurnNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-nft [owner] [id]",
		Args:  cobra.ExactArgs(2),
		Short: "burn an nft",
		Example: `
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			nft, err := ParseNFT(args[1])
			if err != nil {
				return err
			}

			msg := composable.MsgBurnNFT{
				Owner: owner,
				Nft:   *nft,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdUpdateNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-nft [id] [--uri] [--uri-hash]",
		Args:    cobra.ExactArgs(1),
		Short:   "update an nft",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			nft, err := ParseNFT(args[0])
			if err != nil {
				return err
			}

			owner := composable.ClassOwner(nft.ClassId).String()
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := composable.MsgUpdateNFT{
				Nft: *nft,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
