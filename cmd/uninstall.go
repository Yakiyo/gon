package cmd

import (
	"fmt"
	"os"

	"github.com/Yakiyo/gon/aliases"
	"github.com/Yakiyo/gon/versions"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a locally installed version",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	SilenceUsage:  true,
	SilenceErrors: true,
	Args:          cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		version := args[0]
		version = versions.SafeVStr(version)
		vdir, exists := versions.VersionDir(version)
		if !exists {
			return fmt.Errorf("Version %v is not installed locally", version)
		}
		err := os.RemoveAll(vdir)
		if err != nil {
			return anyhow("Unable to remove version directory", err)
		}
		al, err := aliases.AliasMap()
		if err != nil {
			log.Error("Failed to resolve aliases. Skipping alias removals. Please do it manually", "err", err)
			return nil
		}
		val, exists := flipMap(al)[version]
		// we got no aliases, so return early
		if !exists {
			return nil
		}
		for _, v := range val {
			delete(al, v)
		}
		err = aliases.SaveAliases(al)
		if err != nil {
			return anyhow("Error when saving new alias file", err)
		}
		fmt.Printf("Successfully uninstalled version %v\n", version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uninstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uninstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
