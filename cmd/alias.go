package cmd

import (
	"fmt"
	"strings"

	"github.com/Yakiyo/gon/aliases"
	"github.com/Yakiyo/gon/versions"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// aliasCmd represents the alias command
var aliasCmd = &cobra.Command{
	Use:           "alias version alias",
	Short:         "Create an alias to a locally installed version",
	Long:          `Create an alias to an installed version. This alias can be later used in other places to refer to that command`,
	SilenceUsage:  true,
	SilenceErrors: true,
	Args:          cobra.ExactArgs(2),
	Example: "gon alias v1.21.0 default\n" +
		"gon alias v1.17.9 legacy",
	RunE: func(cmd *cobra.Command, args []string) error {
		version, alias := args[0], args[1]
		if strings.ToLower(alias) == "latest" {
			return fmt.Errorf("`latest` is not acceptable as an alias. It is used to refer latest Go version in the app. Please use a different name.")
		}
		version = versions.SafeVStr(version)
		_, exists := versions.VersionDir(version)
		if !exists {
			return fmt.Errorf("Version %v is not installed locally, cannot make alias for it", version)
		}
		aliasmap, err := aliases.AliasMap()
		if err != nil {
			return anyhow("Error when reading alias file", err)
		}
		if v, ok := aliasmap[alias]; ok {
			log.Warn("Existing alias with same name exists, overriding it", "version", v)
		}
		aliasmap[alias] = version
		err = aliases.SaveAliases(aliasmap)
		if err != nil {
			return anyhow("Failed to write json to alias file", err)
		}
		log.Info("New json", "aliases", aliasmap)
		fmt.Printf("Successfully added alias %v for version %v\n", alias, version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(aliasCmd)
}
