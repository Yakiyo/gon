package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Yakiyo/gon/utils/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var validArgs = []string{"bin", "current", "root"}

// pathCmd represents the path command
var pathCmd = &cobra.Command{
	Use:   "path bin|current|root",
	Short: "Show paths used by gon",
	Long: `Show paths used by gon.
	
This can be used to easily add the required path(s) to a shell's env/sys PATH
bin     - Shows the path to the directory where active Go executable is situated
current - Shows the path to the directory of the active Go version
root    - Shows the root directory used by gon`,
	SilenceErrors: true,
	SilenceUsage:  true,
	Args:          cobra.ExactArgs(1),
	ValidArgs:     validArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		arg := strings.ToLower(args[0])
		if !lo.Contains(validArgs, arg) {
			return fmt.Errorf("Missing required argument. Must be one of %v", strings.Join(validArgs, ", "))
		}
		if arg == "bin" {
			fmt.Println(filepath.Join(where.Bin(), "bin"))
		} else if arg == "current" {
			fmt.Println(where.Bin())
		} else if arg == "root" {
			fmt.Println(where.RootDir())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pathCmd)
}
