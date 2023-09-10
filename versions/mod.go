package versions

import (
	"fmt"
	"os"
	"strings"

	"github.com/Yakiyo/gon/utils"
	"github.com/samber/lo"
)

// read go.mod file in current directory and parse it to get go version
func FromGoMod() (string, error) {
	var version string
	if !utils.PathExists("./go.mod") {
		return "", fmt.Errorf("No version provided and current directory does not contain a go.mod file. Provide a version explicitly.")
	}
	b, err := os.ReadFile("./go.mod")
	if err != nil {
		return "", fmt.Errorf("Unable to read local go.mod file, received error %s", err)
	}
	str := string(b)
	// find line that contains `go {{ version }}`
	vline, ok := lo.Find[string](strings.Split(str, "\n"), func(item string) bool {
		return strings.HasPrefix(item, "go")
	})
	if !ok {
		return "", fmt.Errorf("Invalid go.mod file, file does not specify go version")
	}
	version = SafeVStr(strings.TrimSpace(strings.ReplaceAll(vline, "go", "")))
	if version == "" {
		return "", fmt.Errorf("Could not parse version from go.mod file. Line contains %s", vline)
	}
	return version, nil
}
