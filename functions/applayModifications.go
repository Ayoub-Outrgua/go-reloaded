package functions

func ApplayModifications(slice []string) []string {
	slice = ApplayFlag(slice)
	slice = ApplayPunctuations(slice)
	slice = ApplayQuotes(slice)
	return slice
}
