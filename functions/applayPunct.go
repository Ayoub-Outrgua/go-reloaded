package functions

import "strings"

func ApplayPunct(slice []string) []string {
	str := ""
	s := strings.Join(slice, " ")

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && s[i] == ' ' && IsPonctuation(rune(s[i+1])) {
			continue
		}
		if i < len(s)-1 && IsPonctuation(rune(s[i])) && s[i+1] != ' ' && !IsPonctuation(rune(s[i+1])) {
			str += string(s[i]) + " "
		} else {
			str += string(s[i])
		}
	}
	slice = strings.Fields(str)
	return slice
}

func IsPonctuation(c rune) bool {
	if c == ',' || c == '.' || c == '!' || c == '?' || c == ';' || c == ':' {
		return true
	}
	return false
}

// func ApplayPunct(slice []string, index int, word string, i int) []string {
// 	if index == 0 && len(word) > 1 {
// 		if i-1 >= 0 {
// 			slice[i-1] = slice[i-1] + string(word[0])
// 			slice[i] = slice[i][1:]
// 			i--
// 		}
// 	} else if index == 0 && len(word) == 1 {
// 		if i-1 > 0 {
// 			slice[i-1] = slice[i-1] + string(word[0])
// 			slice[i] = ""
// 			slice = CLeanSlice(slice)
// 			i--
// 		}
// 	} else if index > 0 && index < len(word)-1 {
// 		slice[i] = word[:index+1]
// 		endSlice := slice[i+1:]
// 		addToSlice := []string{word[index+1:]}
// 		// fmt.Println("slice [i] after +++ : ", slice[i])
// 		// fmt.Println("endSlice : ", endSlice)
// 		// fmt.Println("slice[:i+1] : ", slice[:i+1])
// 		// fmt.Println("addToSlice : ", addToSlice)
// 		slice = append(slice[:i+1], append(addToSlice, endSlice...)...)
// 		i--
// 	}
// 	return slice
// }

// if index == 0 && len(word) > 1 {
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
