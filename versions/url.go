package versions

import "runtime"

// taken from https://github.com/golang/dl/blob/55c644201171612f902249a444b1bc505b0ef6a9/internal/version/version.go#L421

// generate url to the archive for a specific version
func VersionArchiveUrl(version string, arch string) string {
	os := runtime.GOOS
	if arch == "" {
		arch = runtime.GOARCH
	}
	ext := ".tar.gz"
	if os == "windows" {
		ext = ".zip"
	}

	if os == "linux" && arch == "arm" {
		arch = "armv61"
	}
	return "https://go.dev/dl/go" + version + "." + os + "-" + arch + ext
}
