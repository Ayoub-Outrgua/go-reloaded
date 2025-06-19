package main

import (
	"fmt"
	"os"
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

	for i := 0; i < len(slice); i++ {
		if slice[i] == "(up)" {
			if i != 0 {
				slice[i-1] = strings.ToUpper(slice[i-1])
			} else {
				slice[i] = ""
			}
			slice[i] = ""
		} else if slice[i] == "(low)" {
			if i != 0 {
				slice[i-1] = strings.ToLower(slice[i-1])
			} else {
				slice[i] = ""
			}
			slice[i] = ""
		} else if slice[i] == "(cap)" {
			if i != 0 {
				slice[i-1] = functions.Capitalized(slice[i-1])
			} else {
				slice[i] = ""
			}
			slice[i] = ""
		}
	}

	// resSlice := []string{}
	// for i := 0; i < len(slice); i++ {
	// 	if i+1 < len(slice) && (slice[i+1] == "(up)") {
	// 		continue
	// 	} else if slice[i] == "(up)" {
	// 		if i != 0 {
	// 			resSlice = append(resSlice, strings.ToUpper(slice[i-1]))
	// 		}
	// 	} else if i+1 < len(slice) && (slice[i+1] == "(low)") {
	// 		continue
	// 	} else if slice[i] == "(low)" {
	// 		if i != 0 {
	// 			resSlice = append(resSlice, strings.ToLower(slice[i-1]))
	// 		}
	// 	} else if i+1 < len(slice) && (slice[i+1] == "(cap)") {
	// 		continue
	// 	} else if slice[i] == "(cap)" {
	// 		if i != 0 {
	// 			resSlice = append(resSlice, functions.Capitalized(slice[i-1]))
	// 		}
	// 	} else {
	// 		resSlice = append(resSlice, slice[i])
	// 	}
	// }

	fmt.Println(string(file))
	fmt.Println(slice)
	gg := strings.Join(slice, " ")
	slice1 := strings.Fields(gg)
	fmt.Println(slice1)
	// fmt.Println(resSlice)

	erre := os.WriteFile(outputFile, file, 0o644)
	if erre != nil {
		fmt.Println("error in write file")
		return
	}
}
