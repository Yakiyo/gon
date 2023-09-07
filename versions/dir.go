package versions

import (
	"path/filepath"

	"github.com/Yakiyo/gon/utils"
	"github.com/Yakiyo/gon/utils/where"
)

// get path to a version dir. the bool represents wether the
// directory exists or not
func VersionDir(version string) (string, bool) {
	vdir := filepath.Join(where.Installations(), version)
	return vdir, utils.PathExists(vdir)
}
