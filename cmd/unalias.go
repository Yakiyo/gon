package cmd

import (
	"fmt"

	"github.com/Yakiyo/gon/aliases"
	"github.com/spf13/cobra"
)

// unaliasCmd represents the unalias command
var unaliasCmd = &cobra.Command{
	Use:           "unalias",
	Short:         "Remove an alias",
	SilenceUsage:  true,
	SilenceErrors: true,
	Args:          cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		am, err := aliases.AliasMap()
		if err != nil {
			return err
		}
		key := args[0]
		if _, ok := am[key]; !ok {
			return fmt.Errorf("No alias with name %v exists", key)
		}
		delete(am, key)
		err = aliases.SaveAliases(am)
		if err != nil {
			return anyhow("Error when saving new alias map", err)
		}
		fmt.Println("Successfully removed alias")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(unaliasCmd)
}
