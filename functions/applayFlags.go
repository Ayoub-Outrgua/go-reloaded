package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func ApplayFlag(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[i] == "(up)" {
			if i != 0 {
				slice[i-1] = strings.ToUpper(slice[i-1])
			}
			slice[i] = ""
			slice = CLeanSlice(slice)
			i--
		} else if slice[i] == "(low)" {
			if i != 0 {
				slice[i-1] = strings.ToLower(slice[i-1])
			}
			slice[i] = ""
			slice = CLeanSlice(slice)
			i--
		} else if slice[i] == "(cap)" {
			if i != 0 {
				slice[i-1] = Capitalized(slice[i-1])
			}
			slice[i] = ""
			slice = CLeanSlice(slice)
			i--
		} else if slice[i] == "(bin)" {
			if i != 0 {
				checkError := false
				num, err := strconv.ParseInt(slice[i-1], 2, 64)
				if err != nil {
					fmt.Println("Error: this is invalide (bin) :", err)
					checkError = true
					// return
				}
				if !checkError {
					slice[i-1] = strconv.Itoa(int(num))
				}
			}
			slice[i] = ""
			slice = CLeanSlice(slice)
			i--
		} else if slice[i] == "(hex)" {
			if i != 0 {
				checkError := false
				num, err := strconv.ParseInt(slice[i-1], 16, 64)
				if err != nil {
					fmt.Println("Error: this is invalide (hex) :", err)
					checkError = true
					// return
				}
				if !checkError {
					slice[i-1] = strconv.Itoa(int(num))
				}
			}
			slice[i] = ""
			slice = CLeanSlice(slice)
			i--
		} else if slice[i] == "(up," || slice[i] == "(low," || slice[i] == "(cap," {
			check := false
			if i+1 < len(slice) {
				if strings.Count(slice[i+1], ")") != 1 {
					fmt.Println("Error: this is not valid flag : ", string(slice[i]), string(slice[i+1]))
					check = true
				}
				nb := strings.TrimRight(slice[i+1], ")")
				number, err := strconv.Atoi(string(nb))
				if err != nil {
					fmt.Println("Error: this is invalide number :", err)
					continue
				}
				if slice[i] == "(low," && !check {
					if number >= len(slice[:i]) {
						for j := len(slice[:i]); j > 0; j-- {
							slice[j-1] = strings.ToLower(slice[j-1])
						}
					} else {
						for j := len(slice[:i]); j > 0; j-- {
							if number > 0 {
								slice[j-1] = strings.ToLower(slice[j-1])
								number--
							} else {
								break
							}
						}
					}
					slice[i] = ""
					slice[i+1] = ""
					slice = CLeanSlice(slice)
					i--
				} else if slice[i] == "(up," && !check {
					if number >= len(slice[:i]) {
						for j := len(slice[:i]); j > 0; j-- {
							slice[j-1] = strings.ToUpper(slice[j-1])
						}
					} else {
						for j := len(slice[:i]); j > 0; j-- {
							if number > 0 {
								slice[j-1] = strings.ToUpper(slice[j-1])
								number--
							} else {
								break
							}
						}
					}
					slice[i] = ""
					slice[i+1] = ""
					slice = CLeanSlice(slice)
					i--
				} else if slice[i] == "(cap," && !check {
					if number >= len(slice[:i]) {
						for j := len(slice[:i]); j > 0; j-- {
							slice[j-1] = Capitalized(slice[j-1])
						}
					} else {
						for j := len(slice[:i]); j > 0; j-- {
							if number > 0 {
								slice[j-1] = Capitalized(slice[j-1])
								number--
							} else {
								break
							}
						}
					}
					slice[i] = ""
					slice[i+1] = ""
					slice = CLeanSlice(slice)
					i--
				}
			}
		}
	}
	return slice
}
