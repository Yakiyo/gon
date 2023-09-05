package config

import (
	"github.com/spf13/cobra"
	v "github.com/spf13/viper"
)

// bind flags to viper
func BindFlags(cmd *cobra.Command) {

	look := cmd.Flags().Lookup
	bind := func(key string, flag string) { v.BindPFlag(key, look(flag)) }

	// bind command line flags to viper keys
	bind("log_level", "log-level")
	bind("color", "color")
}
