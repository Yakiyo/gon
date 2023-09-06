package versions

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/Yakiyo/gom/utils/where"
	json "github.com/json-iterator/go"
	"github.com/samber/lo"
)

// a simple regex to do some hacky weird-ass scraping
var spanRegex = regexp.MustCompile(`\<span\>go(?P<version>([0-9]+|\.)+)\</span\>`)

// list all versions
//
// this uses the local versions cache file, if its stale or does not exist,
// it generates the file and then uses it
func List() ([]string, error) {
	if needToUpdate() {
		err := createFile()
		if err != nil {
			return []string{}, err
		}
	}
	b, err := os.ReadFile(where.VersionCache())
	if err != nil {
		return []string{}, err
	}
	versions := []string{}
	err = json.Unmarshal(b, &versions)
	return versions, err
}

func fetchVersions() (versions []string, err error) {
	versions = []string{}
	r, err := http.Get("https://go.dev/dl/")
	if err != nil {
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	str := string(body)
	matches := spanRegex.FindAllStringSubmatch(str, -1)
	if matches == nil || len(matches) < 1 {
		err = fmt.Errorf("Regex did not match any version from https://go.dev/dl/. Please file a bug report")
		return
	}
	versions = lo.FilterMap[[]string, string](matches, func(i []string, _ int) (string, bool) {
		if len(i) < 2 {
			return "", false
		}
		return i[1], true
	})
	versions = lo.Uniq(versions)
	return
}
