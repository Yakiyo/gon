package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Yakiyo/gon/utils"
	"github.com/Yakiyo/gon/utils/where"
	"github.com/spf13/cobra"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:           "current",
	Short:         "View currently active Go version",
	Args:          cobra.NoArgs,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		version, err := getCurrent()
		if err != nil {
			return err
		}
		if version == "" {
			fmt.Println("No version currently active")
			return nil
		}
		fmt.Println("v" + version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}

func getCurrent() (string, error) {
	cur := where.Current()
	if !utils.PathExists(cur) {
		return "", nil
	}
	link, err := os.Readlink(cur)
	if err != nil {
		return "", anyhow("Error when reading link", err)
	}
	return filepath.Base(link), nil
}
