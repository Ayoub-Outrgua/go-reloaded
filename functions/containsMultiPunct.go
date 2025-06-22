package functions

func ContainsMultiPunct(word string) bool {
	for _, v := range word {
		if v != '.' && v != ',' && v != '!' && v != '?' && v != ':' && v != ';' {
			return false
		}
	}
	return true
}
