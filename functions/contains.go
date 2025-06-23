package functions

import (
	"strings"
)

func Contains(word string) bool {
	// index := 0
	if strings.Contains(word, ".") {
		// index = strings.Index(word, ".")
		return true
	} else if strings.Contains(word, ",") {
		// index = strings.Index(word, ",")
		return true
	} else if strings.Contains(word, "!") {
		// index = strings.Index(word, "!")
		return true
	} else if strings.Contains(word, "?") {
		// index = strings.Index(word, "?")
		return true
	} else if strings.Contains(word, ":") {
		// index = strings.Index(word, ":")
		return true
	} else if strings.Contains(word, ";") {
		// index = strings.Index(word, ";")
		return true
	}
	return false
}
