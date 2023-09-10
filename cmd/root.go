package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/Yakiyo/gon/utils"
	"github.com/Yakiyo/gon/utils/config"
	logger "github.com/Yakiyo/gon/utils/log"
	"github.com/Yakiyo/gon/utils/meta"
	"github.com/Yakiyo/gon/utils/where"
	"github.com/charmbracelet/log"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   meta.AppName,
	Short: "Go Version Manager",
	Long: `Gon is an easy to use version manager for Go, written in Go, designed with simplicity and speed in mind.

For queries, issues or bug reports, visit: https://github.com/Yakiyo/gon`,
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
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(
		fmt.Sprintf("{{ .Name }} version {{ .Version }} %v/%v (built at %v)\n",
			runtime.GOOS,
			runtime.GOARCH,
			meta.BuiltAt,
		),
	)

	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiCyan + cc.Bold + cc.Underline,
		Commands:        cc.HiYellow + cc.Bold,
		CmdShortDescr:   cc.Bold,
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
	f.String("bin", "", "Set bin directory. Defaults to ~/gon/current")
	f.String("root-dir", "", "Root directory for gon to use. Defaults to ~/gon")
	f.String("arch", "", "Override architecture to use")
}

var anyhow = utils.Anyhow
