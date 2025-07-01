package functions

import (
	"strings"
	"unicode"
)

func ApplayQuotes(slice []string) []string {
	if len(slice) == 0 {
		return []string{}
	}
	str := strings.Join(slice, " ")
	sliceStr := []rune(str)

	checkOpen := false
	checkClose := false
	indexOpen := 0
	indexClose := 0
	newStr := ""

	if len(str) == 1 {
		return []string{string(str[0])}
	}

	for i := 0; i < len(sliceStr); i++ {
		if i+1 < len(sliceStr) && i-1 >= 0 && IsQuotes(sliceStr[i]) && unicode.IsLetter(sliceStr[i+1]) && unicode.IsLetter(sliceStr[i-1]) {
			newStr += string(sliceStr[i])
			continue
		}
		if IsQuotes(sliceStr[i]) && !checkOpen {
			indexOpen = i
			checkOpen = true
			for j := i + 1; j < len(sliceStr); j++ {
				if j+1 < len(sliceStr) && j-1 >= 0 && IsQuotes(sliceStr[j]) && unicode.IsLetter(sliceStr[j+1]) && unicode.IsLetter(sliceStr[j-1]) {
					continue
				}
				if IsQuotes(sliceStr[j]) && checkOpen {
					indexClose = j
					checkClose = true
					break
				}
			}
		}
		if checkClose {
			if indexClose+1 < len(sliceStr) && (IsPonctuation(sliceStr[indexClose+1]) || sliceStr[indexClose+1] == ' ') {
				newStr += " '" + strings.TrimSpace(string(sliceStr[indexOpen+1:indexClose])) + "'"
			} else {
				newStr += " '" + strings.TrimSpace(string(sliceStr[indexOpen+1:indexClose])) + "' "
			}
			checkClose = false
			checkOpen = false
			i = indexClose
		} else {
			newStr += string(sliceStr[i])
		}
	}

	newSlice := []string{}
	newSlice = []string{newStr}
	newSlice = CLeanSlice(newSlice)
	return newSlice
}

func IsQuotes(r rune) bool {
	if r == '\'' {
		return true
	}
	return false
}
