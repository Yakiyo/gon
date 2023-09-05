package cmd

import (
	"os"

	"github.com/Yakiyo/go-template/utils"
	"github.com/Yakiyo/go-template/utils/config"
	logger "github.com/Yakiyo/go-template/utils/log"
	"github.com/Yakiyo/go-template/utils/meta"
	"github.com/charmbracelet/log"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     meta.AppName,
	Short:   "Shorter description",
	Long:    `Longer description`,
	Version: meta.Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config.BindFlags(cmd)

		if cfg := utils.Must(cmd.Flags().GetString("config")); cfg != "" {
			// when passed the `--config/-c` flag, set viper
			// to explicitly read from that file, this will
			// error if the config file does not exist
			viper.SetConfigFile(cfg)
		}

		// read config
		config.Read()

		logger.SetLevel(viper.GetString("log_level"))
		utils.SetColor(viper.GetString("color"))

		log.Debug(viper.AllSettings())
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
	// allow users to set custom config files
	f.StringP("config", "c", "", "Path to config file")
	// dont mention debug level, usually users dont need, only on the dev side
	f.String("log-level", "", "Set log level [info, warn, error, fatal]")
	f.String("color", "", "Set color output [always, auto, never]")
}
