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
	fmt.Println("str : ", str)
	sliceStr := []rune(str)
	fmt.Println("sliceStr : ", sliceStr)
	newStr := ""
	for i := 0; i < len(sliceStr); i++ {
		if i+1 < len(sliceStr) && sliceStr[i] == 'a' && IsVowel(sliceStr[i+1]) {
			newStr += "an"
		} else {
			newStr += string(sliceStr[i])
		}
	}
	newSlice := []string{}
	newSlice = []string{newStr}
	return newSlice
}

func IsVowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h' {
		return true
	}
	return false
}
