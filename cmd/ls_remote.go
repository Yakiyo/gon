package cmd

import (
	"fmt"

	"github.com/Yakiyo/gom/versions"
	"github.com/spf13/cobra"
)

// lsRemoteCmd represents the lsRemote command
var lsRemoteCmd = &cobra.Command{
	Use:           "ls-remote",
	Short:         "List all available Go versions",
	Long:          `View all available versions of Go. Available versions are taken from https://go.dev/dl/`,
	Args:          cobra.NoArgs,
	Aliases:       []string{"list-remote"},
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		versions, err := versions.List()
		if err != nil {
			return err
		}
		for _, version := range versions {
			fmt.Println(version)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsRemoteCmd)
}
