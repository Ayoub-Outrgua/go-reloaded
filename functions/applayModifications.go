package functions

import (
	"strings"
)

func ApplayModifications(str string) string {
	lines := strings.Split(str, "\n")
	var FinalSlice [][]string
	for _, line := range lines {
		words := strings.Fields(line)
		words = ApplayFlag(words)
		words = ApplayPunctuations(words)
		words = ApplayQuotes(words)
		words = ApplayVowels(words)
		FinalSlice = append(FinalSlice, words)
	}

	strFinal := ""
	for i, v := range FinalSlice {
		v = CLeanSlice(v)
		strFinal += strings.Join(v, " ")
		if i < len(FinalSlice)-1 {
			strFinal += "\n"
		}
	}
	return strFinal
}
