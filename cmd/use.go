package cmd

import (
	"fmt"
	"os"

	"github.com/Yakiyo/gon/aliases"
	"github.com/Yakiyo/gon/utils"
	"github.com/Yakiyo/gon/utils/where"
	"github.com/Yakiyo/gon/versions"
	"github.com/charmbracelet/log"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Set a version",
	Long: `Use an installed version of Go.

This symlinks the versions's directory to the directory shown by ` + "`gon path current`\n" +
		`Use takes maximum one argument, which if present can be a semver compliant version or an alias.
If no argument is provided, it tries to parse the go.mod file from the current directory.

In order to use the ` + "`used` version add `go path bin` to your shell's startup script and add the output to path.",
	SilenceUsage:  true,
	SilenceErrors: true,
	Args:          cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var version string
		if len(args) != 1 {
			var err error
			version, err = versions.FromGoMod()
			if err != nil {
				return err
			}
		} else {
			version = versions.SafeVStr(args[0])
			installed, err := listLocals()
			if err != nil {
				return anyhow("Unable to resolve locally installed versions", err)
			}
			isAVersion := lo.Contains(installed, version)
			// if not a version, try checking if it's an alias or not
			if !isAVersion {
				log.Info("Attempting to match argument to alias")
				al, err := aliases.AliasMap()
				if err != nil {
					return err
				}
				var ok bool
				version, ok = al[args[0]]
				if !ok {
					return fmt.Errorf("Invalid version. No version %v or alias found.", args[0])
				}
				log.Info("Resolved argument %v to version %v", args[0], version)
			}
		}
		if cv, err := getCurrent(); err == nil && cv != "" && version == cv {
			fmt.Printf("Version %v is the currently set one\n", version)
			return nil
		}
		vdir, _ := versions.VersionDir(version)
		current := where.Current()
		if utils.PathExists(current) {
			log.Info("Removing previous link")
			err := os.RemoveAll(current)
			if err != nil {
				return anyhow("Error when removing old symlinked value", err)
			}
		}
		log.Info("Creating symlink", "src", vdir, "dest", current)
		err := os.Symlink(vdir, current)
		if err != nil {
			return anyhow("Error when creating symlink", err)
		}
		fmt.Printf("Successfully set %v as current\n", version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
