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
	// checkCloseQuates := false
	s := ""
	slc := []string{}
	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && i-1 >= 0 && IsQuotes(rune(str[i])) &&
			!IsQuotes(rune(str[i+1])) && str[i+1] != ' ' &&
			!IsQuotes(rune(str[i-1])) && str[i-1] != ' ' {
			s += string(str[i])
			continue
		}
		if IsQuotes(rune(str[i])) {
			stopIndex := i
			for j := i + 1; j < len(str); j++ {
				if IsQuotes(rune(str[j])) {
					// checkCloseQuates = true
					fmt.Println("insid if j = ", j)
					stopIndex = j
					break
				}
			}
			
			s += "'" + strings.TrimSpace(str[i+1 : stopIndex]) + "'"
			fmt.Println("str[i : stopIndex+1] = ", str[i+1 : stopIndex])
			fmt.Println("s = ", s)

			i = stopIndex + 1
		}
		s += string(str[i])
	}
	slc = append(slc, s)
	return slc
}

func IsQuotes(r rune) bool {
	if r == '\'' {
		return true
	}
	return false
}
