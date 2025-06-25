package functions

func ApplayModifications(slice []string) []string {
	slice = ApplayFlag(slice)
	slice = ApplayPunctuations(slice)

	return slice
}
