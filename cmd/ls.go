package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yakiyo/gon/aliases"
	"github.com/Yakiyo/gon/utils"
	"github.com/Yakiyo/gon/utils/where"
	"github.com/fatih/color"
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
		vers, err := listLocals()
		if err != nil {
			return err
		}
		if len(vers) < 1 {
			fmt.Println("No versions installed")
			return nil
		}
		am, err := aliases.AliasMap()
		// if theres any error when getting alias map, just set it to an empty
		// map and continue
		if err != nil {
			am = map[string]string{}
		}
		current, _ := getCurrent()
		fm := flipMap(am)
		var s string
		lo.ForEach[string](lo.Reverse(vers), func(i string, _ int) {
			als, ok := fm[i]
			if ok {
				s = "• " + i + " " + conceal(strings.Join(als, " "))
			} else {
				s = "• " + i
			}
			if i == current {
				s = color.CyanString(s)
			}
			fmt.Println(s)
		})
		return nil
	},
}

var conceal = color.New(color.Faint).Sprint

func init() {
	rootCmd.AddCommand(lsCmd)
}

func flipMap(m map[string]string) map[string][]string {
	res := map[string][]string{}
	for k, v := range m {
		old, ok := res[v]
		if ok {
			old = append(old, k)
		} else {
			old = []string{k}
		}
		res[v] = old
	}
	return res
}

// list locally installed versions
func listLocals() ([]string, error) {
	intls := where.Installations()
	if !utils.PathExists(intls) {
		return []string{}, nil
	}
	dirs, err := os.ReadDir(intls)
	if err != nil {
		return []string{}, fmt.Errorf("Unable to read dir entries for installations due to error %v", err)
	}
	vers := lo.FilterMap[fs.DirEntry, string](dirs, func(f fs.DirEntry, _ int) (string, bool) {
		if !f.IsDir() || strings.HasPrefix(f.Name(), ".") {
			return "", false
		}
		return filepath.Base(f.Name()), true
	})
	return vers, nil
}
