package functions

func IsPonctuation(r rune) bool {
	if r == ',' || r == '.' || r == '!' || r == '?' || r == ';' || r == ':' {
		return true
	}
	return false
}
