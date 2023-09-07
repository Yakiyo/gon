package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yakiyo/gom/utils"
	"github.com/Yakiyo/gom/utils/where"
	"github.com/Yakiyo/gom/versions"
	"github.com/charmbracelet/log"
	json "github.com/json-iterator/go"
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
	Example: "gom alias v1.21.0 default\n" +
		"gom alias v1.17.9 legacy",
	RunE: func(cmd *cobra.Command, args []string) error {
		version, alias := args[0], args[1]
		if strings.ToLower(alias) == "latest" {
			return fmt.Errorf("`latest` is not acceptable as an alias. It is used to refer latest Go version in the app. Please use a different name.")
		}
		version = versions.SafeVStr(version)
		vdir := filepath.Join(where.Installations(), version)
		if !utils.PathExists(vdir) {
			return fmt.Errorf("Version %v is not installed locally, cannot make alias for it", version)
		}
		aliasPath := where.Aliases()
		err := utils.EnsureFile(aliasPath, "{}")
		if err != nil {
			return anyhow("Unexpected error when creating alias file", err)
		}
		aliasFile, err := os.ReadFile(aliasPath)
		if err != nil {
			return anyhow("Failure when reading alias file", err)
		}
		var aliases map[string]string
		err = json.Unmarshal(aliasFile, &aliases)
		if err != nil {
			return anyhow("Invalid content in alias.json file", err)
		}
		if v, ok := aliases[alias]; ok {
			log.Warn("Existing alias with same name exists, overriding it", "version", v)
		}
		aliases[alias] = version
		content, err := json.MarshalIndent(aliases, "", " ")
		if err != nil {
			return anyhow("Failed to marshal alias json", err)
		}
		err = os.WriteFile(aliasPath, content, os.ModePerm)
		if err != nil {
			return anyhow("Failed to write json to alias file", err)
		}
		log.Info("New json", "aliases", string(content))
		fmt.Printf("Successfully added alias %v for version %v\n", alias, version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(aliasCmd)
}
