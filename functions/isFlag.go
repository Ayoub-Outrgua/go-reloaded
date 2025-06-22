package functions

import "strings"

func IsFlag(word string) bool {
	// (up, 45)
	// (up)
	slice := strings.Fields(word)
	word = strings.Join(slice,"")
	if word[len(word)-1] == ')' && word[0] == '(' {

	} else {
		return false
	}
	return true
}
