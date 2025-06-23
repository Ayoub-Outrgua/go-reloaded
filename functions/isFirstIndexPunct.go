package functions

func IsFirstIndexPunct(word string) bool {
	if len(word) == 0 {
		return false
	}
	v := word[0]
	if v == '.' || v == ',' || v == '!' || v == '?' || v == ':' || v == ';' {
		return true
	} else {
		return false
	}
}
