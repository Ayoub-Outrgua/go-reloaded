package functions

import "unicode"

func ApplayVowels(slice []string) []string {
	if len(slice) == 0 {
		return []string{}
	}
	for i := range slice {
		if i+1 < len(slice) && IsVowel(rune(slice[i+1][0])) {
			str := slice[i]
			if str == "a" || str == "A" {
				slice[i] += "n"
			}
			if len(str) >= 2 {
				if CheckBefor(rune(str[0])) && (str[1] == 'a' || str[1] == 'A') {
					slice[i] += "n"
				}
			}
		}
	}
	return slice
}

func IsVowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' || r == 'H' {
		return true
	}
	return false
}

func CheckBefor(r rune) bool {
	if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
		return true
	}
	return false
}
