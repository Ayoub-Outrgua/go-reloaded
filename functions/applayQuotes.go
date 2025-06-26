package functions

import (
	"fmt"
	"strings"
)

func ApplayQuotes(slice []string) []string {
	if len(slice) == 0 {
		return []string{}
	}
	str := strings.Join(slice, " ")
	if len(str) == 0 {
		return []string{}
	}
	checkOpen := false
	checkClose := false
	indexOpen := 0
	indexClose := 0
	newSlice := []string{}
	newStr := ""
	// checkSpaceAfter := false

	if len(str) == 1 {
		newSlice = append(newSlice, string(str[0]))
		return []string{string(str[0])}
	}
	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && i-1 >= 0 && IsQuotes(rune(str[i])) &&
			((str[i+1] >= 'a' && str[i+1] <= 'z') || (str[i+1] >= 'A' && str[i+1] <= 'Z')) &&
			((str[i-1] >= 'a' && str[i-1] <= 'z') || (str[i-1] >= 'A' && str[i-1] <= 'Z')) {
			newStr += string(str[i])
			fmt.Println("for i if 1 :----------")
			fmt.Println("newword = : ", newStr)
			continue
		}
		if IsQuotes(rune(str[i])) && !checkOpen {
			fmt.Println("for i if 2 :----------")
			indexOpen = i
			checkOpen = true
			for j := i + 1; j < len(str); j++ {
				if j+1 < len(str) && j-1 >= 0 && IsQuotes(rune(str[j])) &&
					((str[i+1] >= 'a' && str[i+1] <= 'z') || (str[i+1] >= 'A' && str[i+1] <= 'Z')) &&
					((str[i-1] >= 'a' && str[i-1] <= 'z') || (str[i-1] >= 'A' && str[i-1] <= 'Z')) {
					fmt.Println("for j if 1 :----------")
					continue
				}
				if IsQuotes(rune(str[j])) && checkOpen {
					indexClose = j
					checkClose = true
					fmt.Println("for j if 2 :----------")
					break
				}
			}
		}
		if checkClose {
			fmt.Println("out for if :----------")
			fmt.Println("newword befor = : ", newStr)
			if indexClose+1 < len(str) && IsPonctuation(rune(str[indexClose+1])) {
				fmt.Println("if 1 :----------")
				newStr += " '" + strings.TrimSpace(str[indexOpen+1:indexClose]) + "'"
			} else if indexClose+1 < len(str) && !IsPonctuation(rune(str[indexClose+1])) {
				fmt.Println("if 2 :***********")
				newStr += " '" + strings.TrimSpace(str[indexOpen+1:indexClose]) + "' "
			} else {
				newStr += " '" + strings.TrimSpace(str[indexOpen+1:indexClose]) + "' "
			}
			// if indexClose+1 < len(str) {
			// 	fmt.Println("str[i+1] == ", string(str[indexClose+1]))
			// }
			fmt.Println("newword after = : ", newStr)
			checkClose = false
			checkOpen = false
			fmt.Println("i befor = ", i)
			i = indexClose
			fmt.Println("i after = ", i)

		} else {
			newStr += string(str[i])
		}
	}
	fmt.Println("newword = : ", newStr)

	newSlice = []string{newStr}
	fmt.Println("newSlice = : ", newSlice)

	return newSlice
}

// func ApplayQuotes(slice []string) []string {
// 	if len(slice) == 0 {
// 		return []string{}
// 	}
// 	str := strings.Join(slice, " ")
// 	if len(str) == 0 {
// 		return []string{}
// 	}
// 	checkCloseQuates := false
// 	s := ""
// 	slc := []string{}
// 	for i := 0; i < len(str); i++ {
// 		if i+1 < len(str) && i-1 >= 0 && IsQuotes(rune(str[i])) &&
// 			str[i+1] != ' ' &&
// 			str[i-1] != ' ' {
// 			s += string(str[i])
// 			continue
// 		}
// 		if IsQuotes(rune(str[i])) && !checkCloseQuates {
// 			stopIndex := i
// 			for j := i + 1; j < len(str); j++ {
// 				if IsQuotes(rune(str[j])) &&
// 					j+1 < len(str) && j-1 >= 0 && IsQuotes(rune(str[j])) &&

// 					str[j-1] == ' ' {
// 					checkCloseQuates = true
// 					// fmt.Println("insid if j = ", j)
// 					stopIndex = j
// 					break
// 				}
// 			}
// 			if IsQuotes(rune(str[i])) && checkCloseQuates {
// 				s += "'" + strings.TrimSpace(str[i+1:stopIndex]) + "'" + " "
// 				// fmt.Println("str[i : stopIndex+1] = ", str[i+1 : stopIndex])
// 				// fmt.Println("s = ", s)
// 				checkCloseQuates = false
// 				i = stopIndex
// 				continue
// 			}

// 		}
// 		s += string(str[i])
// 	}
// 	slc = append(slc, s)
// 	return slc
// }

func IsQuotes(r rune) bool {
	if r == '\'' {
		return true
	}
	return false
}
