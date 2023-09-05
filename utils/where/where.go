// Package for finding directories and files related to the app
package where

import (
	"path/filepath"

	"github.com/Yakiyo/gom/utils/meta"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"

	homedir "github.com/mitchellh/go-homedir"
)

var join = filepath.Join

var root string

func init() {
	root = join(Home(), meta.AppName)
}

// find user home dir. In our case, we panic when we can't get it
func Home() string {
	path, err := homedir.Dir()
	if err != nil {
		log.Fatal("Unable to locate user home dir, consider manually setting the value of $HOME/$USERPROFILE env var")
	}
	return path
}

// get the `bin` directory, this is where the current active version is stored
func Bin() string {
	if conf := viper.GetString("bin"); conf != "" {
		return conf
	}
	return join(root, "go")
}

// get installations directory
func Installations() string {
	return join(root, "installations")
}

// get the `alias.json` file where info about aliases is stored
func Aliases() string {
	return join(root, "alias.json")
}

// set root dir
func SetRoot(path string) {
	path, _ = homedir.Expand(path)
	root = path
}

// app root dir
func RootDir() string {
	return root
}
