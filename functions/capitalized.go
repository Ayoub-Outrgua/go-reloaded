package functions

import (
	"strings"
	"unicode"
)

func Capitalized(s string) string {
	str := ""
	checkFirstLitter := false
	for _, v := range s {
		if unicode.IsLetter(v) && !checkFirstLitter {
			str += strings.ToUpper(string(v))
			checkFirstLitter = true
		} else if checkFirstLitter {
			str += strings.ToLower(string(v))
		} else if !(unicode.IsLetter(v)) {
			str += string(v)
		}
	}
	return str
}
