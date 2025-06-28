package functions

import (
	"fmt"
	"strings"
)

func ApplayVowels(slice []string) []string {
	if len(slice) == 0 {
		return []string{}
	}
	str := strings.Join(slice, " ")
	newStr := ""
	for i, v := range str {
		fmt.Println("v : ", v)
		if (i+2 < len(str)) && (v == 'a' || v == 'A') && (IsVowel(rune(str[i+2]))) && rune(str[i+1]) == ' ' {
			newStr += "an"
		} else {
			newStr += string(v)
		}
	}
	newSlice := []string{}
	newSlice = []string{newStr}
	return newSlice
}

func IsVowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' || r == 'H' {
		return true
	}
	return false
}
