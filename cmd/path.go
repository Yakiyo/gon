package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/Yakiyo/gom/utils"
	"github.com/Yakiyo/gom/utils/where"
	"github.com/spf13/cobra"
)

// pathCmd represents the path command
var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show root path used by gom",
	Long: `Show root path used by gom.
	
This is handy to automatically add the path to your shell's env`,
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	f := pathCmd.Flags()
	f.Bool("bin", false, "Show the bin directory, within which is the go executable")
	f.Bool("current", false, "Show the current directory, within which the currently used version is stored")
	f.Bool("root", false, "Show gom root dir, the directory used by gom")
	rootCmd.AddCommand(pathCmd)

	// need to define it separately, else it creates a cycle with pathCmd
	pathCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if used("bin") {
			fmt.Println(filepath.Join(where.Bin(), "bin"))
		} else if used("current") {
			fmt.Println(where.Bin())
		} else if used("root") {
			fmt.Println(where.RootDir())
		} else {
			return fmt.Errorf("Must use one of bin, current or root flag for this command")
		}
		return nil
	}
}

func used(flag string) bool {
	return utils.Must(pathCmd.Flags().GetBool(flag))
}
