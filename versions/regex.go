package versions

import (
	"regexp"
	"strings"
)

var versionRegex = regexp.MustCompile("^" + `v?([0-9]+)(\.[0-9]+)?(\.[0-9]+)?` +
	`(-([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?` +
	`(\+([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?` + "$")

// if string is valid semver or not
func IsValid(version string) bool {
	return versionRegex.MatchString(strings.ToLower(version))
}
