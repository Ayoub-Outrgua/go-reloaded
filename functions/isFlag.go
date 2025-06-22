package functions

import (
	"strings"
	"unicode"
)

func IsFlag(word string) bool {
	// (up, 45)
	// (up)
	slice := strings.Fields(word)
	word = strings.Join(slice, "")
	if word == "(up)" || word == "(low)" || word == "(cap)" {
		return true
	}
	if word[0] == '(' && word[len(word)-1] == ')' {
		if strings.Contains(word, ",") {
			index := strings.Index(word, ",")
			if word[1:index] != "low" && word[1:index] != "up" && word[1:index] != "cap" {
				return false
			}
			checkWord := word[index+1 : len(word)-1]
			for _, v := range checkWord {
				if !unicode.IsDigit(v) {
					return false
				}
			}
		} else {
			return false
		}
	} else {
		return false
	}
	return true
}
