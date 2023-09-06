package govers

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/samber/lo"
)

// a simple regex to do some hacky weird-ass scraping
var spanRegex = regexp.MustCompile(`\<span\>go(?P<version>([0-9]+|\.)+)\</span\>`)

func List() (versions []string, err error) {
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
