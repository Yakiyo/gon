package cmd

import (
	"os"
	"strings"

	"github.com/Yakiyo/gom/utils"
	"github.com/Yakiyo/gom/utils/config"
	logger "github.com/Yakiyo/gom/utils/log"
	"github.com/Yakiyo/gom/utils/meta"
	"github.com/Yakiyo/gom/utils/where"
	"github.com/charmbracelet/log"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   meta.AppName,
	Short: "Go Version Manager",
	Long: `Gom is an easy to use version manager for Go, designed with simplicity and speed in mind.
Gom itself is written in Go.
Gom is open-source. For queries, issues or bug reports, visit:
https://github.com/Yakiyo/gom`,
	Version: meta.Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		viper.SetEnvPrefix("GOM")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
		viper.AutomaticEnv()
		config.BindFlags(cmd)
		if err := viper.ReadInConfig(); err != nil {
			// It's okay if there isn't a config file
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return err
			}
		}

		logger.SetLevel(viper.GetString("log_level"))
		utils.SetColor(viper.GetString("color"))
		log.Debug(viper.AllSettings())
		where.SetRoot(viper.GetString("root_dir"))
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiCyan + cc.Bold + cc.Underline,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Bold,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		FlagsDataType:   cc.Italic + cc.HiBlue,
		NoExtraNewlines: true,
	})

	f := rootCmd.PersistentFlags()
	// dont mention debug level, usually users dont need, only used on the dev side
	f.String("log-level", "", "Set log level [info, warn, error, fatal]")
	f.String("color", "", "Set color output [always, auto, never]")
	f.String("bin", "", "Set bin directory. Defaults to ~/gom/go")
	f.String("root-dir", "", "Root directory for gom to use. Defaults to ~/gom")
	f.String("arch", "", "Override architecture to use")
}
