// Package for finding directories and files related to the app
package where

import (
	"path/filepath"

	"github.com/Yakiyo/go-template/utils/meta"
	"github.com/charmbracelet/log"

	homedir "github.com/mitchellh/go-homedir"
)

// app root dir, this is private and is to be used interally within the package only
// default: $HOME/{{appName}} - i.e. if app is kubernetes, it would be ~/kubernetes
var root string

func init() {
	root = filepath.Join(Home(), meta.AppName)
}

// find user home dir. In our case, we panic when we can't get it
func Home() string {
	path, err := homedir.Dir()
	if err != nil {
		log.Fatal("Unable to locate user home dir, consider manually setting the value of $HOME/$USERPROFILE env var")
	}
	return path
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

// default config file path - ~/{{appname}}/{{appname}}.toml
// toml is more human friendly so its a good config file syntax
func Config() string {
	return filepath.Join(root, meta.AppName+".toml")
}
