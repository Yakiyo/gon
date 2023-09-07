package versions

import "strings"

// remove that `v` in from of a version
func SafeVStr(version string) string {
	return strings.TrimPrefix(strings.TrimSpace(strings.ToLower(version)), "v")
}
