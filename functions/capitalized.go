package functions

import "strings"

func Capitalized(s string) string {
	str := ""
	for i, v := range s {
		if i == 0 {
			str += strings.ToUpper(string(v))
		} else {
			str += strings.ToLower(string(v))
		}
	}
	return str
}
