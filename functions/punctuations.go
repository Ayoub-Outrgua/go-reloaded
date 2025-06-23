package functions

import "fmt"

func Punctuations(slice []string) []string {
	if len(slice) == 0 {
		return []string{}
	}
	for i := 0; i < len(slice); i++ {
		// fmt.Println("slice [i] befor --- : ", slice[i])
		word := slice[i]
		check := Contains(word)
		fmt.Println("word : ", word)
		checkFirstIndex := IsFirstIndexPunct(word)
		fmt.Println("checkFirstIndex : ", checkFirstIndex)
		checkValidFlag := false
		word2 := ""
		if checkFirstIndex && len(slice) > 1 && i-1 > 0 && len(slice[i-1]) > 1 && slice[i-1][0] == '(' && slice[i-1][len(slice[i-1])-1] == ')' {
			word2 = slice[i-1]
			fmt.Println("word2 in if-1 : ", word2)
			checkValidFlag = IsFlag(word2)
			fmt.Println("checkValidFlag : ", checkValidFlag)
		} else if checkFirstIndex && (len(slice) > 1 && i-2 > 1 && slice[i-1][len(slice[i-1])-1] == ')' && len(slice[i-1]) > 1) &&
			(len(slice) > 2 && len(slice[i-2]) > 2 && slice[i-2][0] == '(') {
			word2 = slice[i-2] + slice[i-1]
			fmt.Println("word2 in if-2 : ", word2)
			checkValidFlag = IsFlag(word2)
		} else if checkFirstIndex && i-2 > 2 {
			word2 = slice[i-2] + slice[i-1]
			fmt.Println("word2 else : ", word2)
			checkValidFlag = IsFlag(word2)
		}

		// if checkFirstIndex {
		// 	word2 = slice[i-2] + slice[i-1]
		// } else if checkFirstIndex {
		// 	word2 = slice[i-3] + slice[i-2] + slice[i-1]
		// } else if checkFirstIndex {
		// 	word2 = slice[i-2] + slice[i-1]
		// 	checkValidFlag = IsFlag(word2)
		// }

		// checkFirstIndex := IsFirstIndexPunct(word)
		// if condition {

		// }
		// if i+1 < len(slice)-1 {
		// 	check2 := ContainsOnlyPunct(slice[i+1])
		// 	if check2 {

		// 	}
		// }
		// indexDif := 0
		// indexEqual := 0
		// checkEqual := false
		// for i, v := range word {
		// 	if v != '.' && v != ',' && v != '!' && v != '?' && v != ':' && v != ';' && checkEqual {
		// 		indexDif = i
		// 	} else {
		// 		checkEqual = true
		// 		indexEqual = i
		// 	}
		// }
		if check && !checkValidFlag {
			slice = ApplayPunct(slice)
		} else if checkValidFlag {
			// slice = ApplayPunct(slice)
			// slice[i-1] = slice2[0]
		}
	}
	return slice
}
