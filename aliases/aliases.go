package aliases

import (
	"github.com/Yakiyo/gon/utils"
	"github.com/Yakiyo/gon/utils/where"
)

func ReadAliasFile() string {
	aliasPath := where.Aliases()
	utils.EnsureFile(aliasPath, "{}")
	return ""
}