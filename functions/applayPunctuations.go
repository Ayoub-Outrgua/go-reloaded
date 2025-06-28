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

	// for i,v := range str {
	// 	if i+1 <= len(str)-1 && v == ' ' && IsPonctuation(v) {
	// 		continue
	// 	}
	// 	if i+1 <= len(str)-1 && IsPonctuation(v) && str[i+1] != ' ' &&
	// 		!IsPonctuation(rune(str[i+1])) {
	// 		strFinal += string(v) + " "
	// 	} else {
	// 		strFinal += string(v)
	// 	}
	// }


	slice = strings.Fields(strFinal)
	return slice
}

// func Punctuations(slice []string) []string {
// 	for i := 0; i < len(slice); i++ {
// 		// fmt.Println("slice [i] befor --- : ", slice[i])
// 		word := slice[i]
// 		check, index := Contains(word)
// 		// if i+1 < len(slice)-1 {
// 		// 	check2 := ContainsOnlyPunct(slice[i+1])
// 		// 	if check2 {

// 		// 	}
// 		// }
// 		// indexDif := 0
// 		// indexEqual := 0
// 		// checkEqual := false
// 		// for i, v := range word {
// 		// 	if v != '.' && v != ',' && v != '!' && v != '?' && v != ':' && v != ';' && checkEqual {
// 		// 		indexDif = i
// 		// 	} else {
// 		// 		checkEqual = true
// 		// 		indexEqual = i
// 		// 	}
// 		// }

// 		if check {
// 			if index == 0 && len(word) > 1 {
// 				if i-1 >= 0 {
// 					slice[i-1] = slice[i-1] + string(word[0])
// 					slice[i] = slice[i][1:]
// 					i--
// 				}
// 			} else if index == 0 && len(word) == 1 {
// 				if i-1 > 0 {
// 					slice[i-1] = slice[i-1] + string(word[0])
// 					slice[i] = ""
// 					slice = CLeanSlice(slice)
// 					i--
// 				}
// 			} else if index > 0 && index < len(word)-1 {
// 				slice[i] = word[:index+1]
// 				endSlice := slice[i+1:]
// 				addToSlice := []string{word[index+1:]}
// 				// fmt.Println("slice [i] after +++ : ", slice[i])
// 				// fmt.Println("endSlice : ", endSlice)
// 				// fmt.Println("slice[:i+1] : ", slice[:i+1])
// 				// fmt.Println("addToSlice : ", addToSlice)
// 				slice = append(slice[:i+1], append(addToSlice, endSlice...)...)
// 				i--
// 			}
// 		}
// 	}
// 	return slice
// }
