package functions

import (
	"strings"
)

func ApplayModifications(str string) string {
	lines := strings.Split(str, "\n")
	var FinalWordsSlice [][]string
	for _, line := range lines {
		words := strings.Fields(line)
		words = ApplayFlag(words)
		words = ApplayPunctuations(words)
		words = ApplayQuotes(words)
		words = ApplayVowels(words)

		// for {
		// 	fmt.Println("inside for ...")
		// 	transformedWords := ApplayFlag(words)
		// 	if reflect.DeepEqual(transformedWords, words) {
		// 		break
		// 	}
		// 	words = transformedWords
		// }

		FinalWordsSlice = append(FinalWordsSlice, words)
	}

	strFinal := ""
	for i, v := range FinalWordsSlice {
		v = CLeanSlice(v)
		strFinal += strings.Join(v, " ")
		if i < len(FinalWordsSlice)-1 {
			strFinal += "\n"
		}
	}
	return strFinal
}
