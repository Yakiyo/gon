package cmd

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// a simple regex to do some hacky weird-ass scraping
var spanRegex = regexp.MustCompile(`\<span\>go(?P<version>([0-9]+|\.)+)\</span\>`)

// lsRemoteCmd represents the lsRemote command
var lsRemoteCmd = &cobra.Command{
	Use:           "ls-remote",
	Short:         "List all available Go versions",
	Long:          `View all available versions of Go. Available versions are taken from https://go.dev/dl/`,
	Args:          cobra.NoArgs,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := http.Get("https://go.dev/dl/")
		if err != nil {
			return err
		}
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		str := string(body)
		matches := spanRegex.FindAllStringSubmatch(str, -1)
		if matches == nil || len(matches) < 1 {
			return fmt.Errorf("Regex did not match any version from https://go.dev/dl/. Please file a bug report")
		}
		versions := lo.FilterMap[[]string, string](matches, func(i []string, _ int) (string, bool) {
			if len(i) < 2 {
				return "", false
			}
			return i[1], true
		})
		versions = lo.Uniq(versions)
		for _, version := range versions {
			fmt.Println(version)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsRemoteCmd)
}
