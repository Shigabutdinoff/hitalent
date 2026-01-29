package config

import (
	"hitalent/helpers"
	"strings"
)

func Config(key string, def any) any {
	if key == "" {
		return def
	}

	parts := strings.SplitN(key, ".", 2)
	switch parts[0] {
	case "database":
		if len(parts) == 1 {
			return Database
		}

		return helpers.DataGet(Database, parts[1], def)
	default:
		return def
	}
}
