package functions

import "strings"

func ApplayPunctuations(slice []string) []string {
	
	str := strings.Join(slice, " ")
	strFinal := ""
	sliceStr := []rune(str)

	for i := 0; i < len(sliceStr); i++ {
		if i+1 <= len(sliceStr)-1 && sliceStr[i] == ' ' && IsPonctuation(sliceStr[i+1]) {
			continue
		}
		if i+1 <= len(sliceStr)-1 && IsPonctuation(sliceStr[i]) && sliceStr[i+1] != ' ' &&
			!IsPonctuation(sliceStr[i+1]) {
			strFinal += string(sliceStr[i]) + " "
		} else {
			strFinal += string(sliceStr[i])
		}
	}

	slice = strings.Fields(strFinal)
	return slice
}

func IsPonctuation(r rune) bool {
	if r == ',' || r == '.' || r == '!' || r == '?' || r == ';' || r == ':' {
		return true
	}
	return false
}