package functions

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
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
				if !(strings.Contains(slice[i+1], ")")) {
					// fmt.Println("Error: this is not valid flag )")
					// return
					check = true
				}
				nb := strings.TrimRight(slice[i+1], ")")
				number, err := strconv.Atoi(string(nb))
				if err != nil {
					// fmt.Println("Error: this is not valid number")
					fmt.Println("Error: this is invalide number :", err)
					continue
					// return
				}
				if slice[i] == "(low," && !check {
					isLetter := false
					if number >= len(slice[:i]) {
						for j := len(slice[:i]); j > 0; j-- {
							slice[j-1] = strings.ToLower(slice[j-1])
						}
					} else {
						for j := len(slice[:i]); j > 0; j-- {
							if number > 0 {
								for _, v := range slice[j-1] {
									if unicode.IsLetter(v) {
										isLetter = true
										break
									}
								}
								if isLetter {
									slice[j-1] = strings.ToLower(slice[j-1])
									number--
									isLetter = false
								}
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
					isLetter := false
					if number >= len(slice[:i]) {
						for j := len(slice[:i]); j > 0; j-- {
							slice[j-1] = strings.ToUpper(slice[j-1])
						}
					} else {
						for j := len(slice[:i]); j > 0; j-- {
							if number > 0 {
								for _, v := range slice[j-1] {
									if unicode.IsLetter(v) {
										isLetter = true
										break
									}
								}
								if isLetter {
									slice[j-1] = strings.ToUpper(slice[j-1])
									number--
									isLetter = false
								}
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
					isLetter := false
					if number >= len(slice[:i]) {
						for j := len(slice[:i]); j > 0; j-- {
							slice[j-1] = Capitalized(slice[j-1])
						}
					} else {
						for j := len(slice[:i]); j > 0; j-- {
							if number > 0 {
								for _, v := range slice[j-1] {
									if unicode.IsLetter(v) {
										isLetter = true
										break
									}
								}
								if isLetter {
									slice[j-1] = Capitalized(slice[j-1])
									number--
									isLetter = false
								}
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
