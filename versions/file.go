package versions

import (
	"fmt"
	"os"
	"time"

	"github.com/Yakiyo/gon/utils"
	"github.com/Yakiyo/gon/utils/where"
	json "github.com/json-iterator/go"
	"github.com/samber/lo"
)

var maxDur = lo.Must(time.ParseDuration("96h")) // 4 days

func createFile() error {
	versions, err := fetchVersions()
	if err != nil {
		return err
	}
	versionsJson, err := json.Marshal(versions)
	if err != nil {
		return fmt.Errorf("Unable to convert versions array to json. Received %s", err)
	}
	lo.Must0(utils.EnsureDir(where.RootDir()))
	return os.WriteFile(where.VersionCache(), versionsJson, os.ModePerm)
}

// determine wether the versions file needs to be updated or not.
// if file doesnt exist, or it has been more than 4 days since it was updated, we update the file
func needToUpdate() bool {
	vfile := where.VersionCache()
	if !utils.PathExists(vfile) {
		return true
	}
	stats, err := os.Stat(vfile)
	if err != nil {
		return true
	}
	return time.Since(stats.ModTime()) > maxDur
}
