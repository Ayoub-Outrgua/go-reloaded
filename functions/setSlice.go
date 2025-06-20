package functions

import "strings"

func SetSlice(s string) []string {
	slice := strings.Fields(s)
	return slice
}
