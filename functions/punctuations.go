package functions

func Punctuations(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		// fmt.Println("slice [i] befor --- : ", slice[i])
		word := slice[i]
		check, index := Contains(word)
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

		if check {
			if index == 0 && len(word) > 1 {
				if i-1 > 0 {
					slice[i-1] = slice[i-1] + string(word[0])
					slice[i] = slice[i][1:]
					i--
				}
			} else if index == 0 && len(word) == 1 {
				if i-1 > 0 {
					slice[i-1] = slice[i-1] + string(word[0])
					slice[i] = ""
					slice = CLeanSlice(slice)
					i--
				}
			} else if index > 0 && index < len(word)-1 {
				slice[i] = word[:index+1]
				endSlice := slice[i+1:]
				addToSlice := []string{word[index+1:]}
				// fmt.Println("slice [i] after +++ : ", slice[i])
				// fmt.Println("endSlice : ", endSlice)
				// fmt.Println("slice[:i+1] : ", slice[:i+1])
				// fmt.Println("addToSlice : ", addToSlice)
				slice = append(slice[:i+1], append(addToSlice, endSlice...)...)
				i--
			}
		}
	}
	return slice
}
