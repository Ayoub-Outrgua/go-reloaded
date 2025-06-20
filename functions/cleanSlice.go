package functions

import "strings"

func CLeanSlice(slice []string) []string {
	str := strings.Join(slice, " ")
	slice = strings.Fields(str)
	return slice
}
