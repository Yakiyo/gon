package config

import (
	"github.com/Yakiyo/go-template/utils/meta"
	"github.com/Yakiyo/go-template/utils/where"
	"github.com/charmbracelet/log"
	v "github.com/spf13/viper"
)

func init() {
	// config file named after the app
	v.SetConfigName(meta.AppName)
	v.SetConfigType("toml")

	// search in default directory
	v.AddConfigPath(where.RootDir())

	// add default values here
	v.SetDefault("log_level", "warn")
	v.SetDefault("color", "auto")
	v.SetDefault("root_dir", where.RootDir())
}

// read config file
func Read() {
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(v.ConfigFileNotFoundError); ok {
			log.Info("Missing log file in default location. Using defaults")
		} else {
			log.Fatal("Error reading config", "err", err)
		}
	}
}
