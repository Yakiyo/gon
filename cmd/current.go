package cmd

import (
	"fmt"
	"os/exec"
	"strings"

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
		_, err := exec.LookPath("go")
		if err != nil {
			fmt.Println("No version of Go is currently active")
			return nil
		}
		out, err := exec.Command("go", "version").Output()
		if err != nil {
			return anyhow("Unable to run Go from command line", err)
		}
		// usually output of `go version` is as follows
		// "go version go{{version}} {{platform}}/{{architecture}}"
		// so it's a 4 word string, so we slice and dice em to get the version
		output := strings.Split(string(out), " ")
		if len(output) < 4 {
			// unexpected output format, so just return it
			fmt.Println(strings.Join(output, " "))
			return nil
		}
		version := "v" + strings.TrimPrefix(output[2], "go")
		fmt.Println(version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
