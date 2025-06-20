package functions

import (
	"strings"
	"unicode"
)

func Capitalized(s string) string {
	str := ""
	check := false
	for _, v := range s {
		if unicode.IsLetter(v) && !check {
			str += strings.ToUpper(string(v))
			check = true
		} else if check {
			str += strings.ToLower(string(v))
		} else if !(unicode.IsLetter(v)) {
			str += string(v)
		}
	}
	return str
}
