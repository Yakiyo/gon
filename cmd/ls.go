package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yakiyo/gon/utils"
	"github.com/Yakiyo/gon/utils/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:           "ls",
	Short:         "List locally installed versions",
	Aliases:       []string{"list"},
	SilenceUsage:  true,
	SilenceErrors: true,
	Args:          cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		intls := where.Installations()
		if !utils.PathExists(intls) {
			fmt.Println("No versions installed")
			return nil
		}
		dirs, err := os.ReadDir(intls)
		if err != nil {
			return fmt.Errorf("Unable to read dir entries for installations due to error %v", err)
		}
		vers := lo.FilterMap[fs.DirEntry, string](dirs, func(f fs.DirEntry, _ int) (string, bool) {
			if !f.IsDir() || strings.HasPrefix(f.Name(), ".") {
				return "", false
			}
			return filepath.Base(f.Name()), true
		})
		if len(vers) < 1 {
			fmt.Println("No versions installed")
			return nil
		}
		lo.ForEach[string](lo.Reverse(vers), func(i string, _ int) {
			fmt.Println(i)
		})
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
