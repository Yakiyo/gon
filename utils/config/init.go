package config

import (
	"runtime"

	"github.com/Yakiyo/gon/utils/where"
	v "github.com/spf13/viper"
)

func init() {
	// add default values here
	v.SetDefault("log_level", "warn")
	v.SetDefault("color", "auto")
	v.SetDefault("root_dir", where.RootDir())
	v.SetDefault("arch", runtime.GOARCH)
}
