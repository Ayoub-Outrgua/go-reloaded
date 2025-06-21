package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"goreload/functions"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("the args are not exist!")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	if !strings.HasSuffix(outputFile, ".txt") || !strings.HasSuffix(inputFile, ".txt") {
		fmt.Println("hdkfj")
		return
	}

	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("error in read file")
		return
	}

	slice := strings.Fields(string(file))
	fmt.Println(string(file))
	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		if slice[i] == "(up)" {
			if i != 0 {
				slice[i-1] = strings.ToUpper(slice[i-1])
			}
			slice[i] = ""
			slice = functions.CLeanSlice(slice)
			i--
		} else if slice[i] == "(low)" {
			if i != 0 {
				slice[i-1] = strings.ToLower(slice[i-1])
			}
			slice[i] = ""
			slice = functions.CLeanSlice(slice)
			i--
		} else if slice[i] == "(cap)" {
			if i != 0 {
				slice[i-1] = functions.Capitalized(slice[i-1])
			}
			slice[i] = ""
			slice = functions.CLeanSlice(slice)
			i--
		} else if slice[i] == "(bin)" {
			if i != 0 {
				check := false
				num, err := strconv.ParseInt(slice[i-1], 2, 64)
				if err != nil {
					// fmt.Println("Error: this is not valid bin")
					check = true
					// return
				}
				if !check {
					slice[i-1] = strconv.Itoa(int(num))
				}
			}
			slice[i] = ""
			slice = functions.CLeanSlice(slice)
			i--
		} else if slice[i] == "(hex)" {
			if i != 0 {
				check := false
				num, err := strconv.ParseInt(slice[i-1], 16, 64)
				if err != nil {
					// fmt.Println("Error: this is not valid hex")
					check = true
					// return
				}
				if !check {
					slice[i-1] = strconv.Itoa(int(num))
				}
			}
			slice[i] = ""
			slice = functions.CLeanSlice(slice)
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
					continue
					// return
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
					slice = functions.CLeanSlice(slice)
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
					slice = functions.CLeanSlice(slice)
					i--
				} else if slice[i] == "(cap," && !check {
					if number >= len(slice[:i]) {
						for j := len(slice[:i]); j > 0; j-- {
							slice[j-1] = functions.Capitalized(slice[j-1])
						}
					} else {
						for j := len(slice[:i]); j > 0; j-- {
							if number > 0 {
								slice[j-1] = functions.Capitalized(slice[j-1])
								number--
							} else {
								break
							}
						}
					}
					slice[i] = ""
					slice[i+1] = ""
					slice = functions.CLeanSlice(slice)
					i--
				}
			}
		}
	}
	slice = functions.CLeanSlice(slice)
	fmt.Println(slice)
	combined := strings.Join(slice, " ")
	byteSlice := []byte(combined)
	erre := os.WriteFile(outputFile, byteSlice, 0o644)
	if erre != nil {
		fmt.Println("error in write file")
		return
	}
}
