package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

func init() {
	log.SetLevel(log.WarnLevel)
	log.SetReportTimestamp(false)
	log.SetOutput(os.Stderr)
}

func SetLevel(lvl string) {
	log.SetLevel(log.ParseLevel(lvl))
	log.Info("Setting log level", "level", lvl)
}
