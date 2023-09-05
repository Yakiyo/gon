package utils

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
)

// wether stdout is terminal or not
func IsAtty() bool {
	fd := os.Stdout.Fd()
	return (isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd))
}

// wether to allow color or not based on - valid levels: always, auto, never
func ParseColorLevel(lvl string) bool {
	lvl = strings.ToLower(lvl)
	var colorAllowed bool
	switch lvl {
	case "never":
		colorAllowed = false
	case "always":
		colorAllowed = true
	default:
		if lvl != "auto" {
			log.Warn("Invalid color level received. Switching to auto", "received", lvl)
		}
		_, noColor := os.LookupEnv("NO_COLOR")
		colorAllowed = !noColor && IsAtty()
	}

	return colorAllowed
}

// Set color
func SetColor(lvl string) {
	color.NoColor = !ParseColorLevel(lvl)
}
