package helpers

import "strings"

// DataGet Get an item from an array or object using "dot" notation./**
func DataGet(target any, key string, def any) any {
	for _, segment := range strings.Split(key, ".") {
		switch typed := target.(type) {
		case map[string]any:
			item, isTrue := typed[segment]
			if !isTrue {
				return def
			}

			target = item
		default:
			return def
		}
	}

	return target
}
