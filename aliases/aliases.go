package aliases

import (
	"os"

	"github.com/Yakiyo/gon/utils"
	"github.com/Yakiyo/gon/utils/where"
	json "github.com/json-iterator/go"
)

// read alias file
func ReadAliasFile() ([]byte, error) {
	aliasPath := where.Aliases()
	err := utils.EnsureFile(aliasPath, "{}")
	if err != nil {
		return []byte{}, utils.Anyhow("Error when creating alias file", err)
	}
	return os.ReadFile(aliasPath)
}

// read aliases to a map
func AliasMap() (map[string]string, error) {
	b, err := ReadAliasFile()
	if err != nil {
		return nil, utils.Anyhow("Failed to read alias file", err)
	}
	var aliases map[string]string
	err = json.Unmarshal(b, &aliases)
	return aliases, err
}

// write map to alias file
func SaveAliases(aliases map[string]string) error {
	j, err := json.MarshalIndent(aliases, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(where.Aliases(), j, os.ModePerm)
}
