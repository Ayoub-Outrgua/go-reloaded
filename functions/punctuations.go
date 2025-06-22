package functions

func Punctuations(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		word := slice[i]
		if i-1 > 0 {
			check, index := Contains(word)
			if check {
				if index == 0 && len(word) > 1 {
					slice[i-1] = slice[i-1] + string(word[0])
					slice[i] = slice[i][1:]
				} else if index != 0 && index != len(word)-1 && len(word) > 1 {
					slice[i] = string(word[0])
					s := []string{string(word[index])}
					slice = append(slice[:index], append(s, slice[index:]...)...)
				} else if index == 0 && len(word) == 1 {
					slice[i-1] = slice[i-1] + string(word[0])
					slice[i] = ""
					slice = CLeanSlice(slice)
				}
			}
			// if slice[i] == "." || slice[i] == "," || slice[i] == "!" || slice[i] == "?" || slice[i] == ":" || slice[i] == ";" {

			// }
		}
	}
	return slice
}
