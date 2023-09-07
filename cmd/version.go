package cmd

import (
	"os"
	"text/template"

	"github.com/Yakiyo/gon/utils/meta"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:           "version",
	Short:         "Version for gon",
	Long:          "View current version of gon. This is the same as the `--version/-v` flag",
	Args:          cobra.NoArgs,
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		tmpl := lo.Must(template.New("version").Parse(cmd.VersionTemplate()))
		lo.Must0(tmpl.Execute(os.Stdout, map[string]string{
			"Name":    meta.AppName,
			"Version": cmd.Root().Version,
		}))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
