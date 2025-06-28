package main

import (
	"fmt"
	"os"

	"goreload/functions"
)

func main() {
	// if len(os.Args) != 3 {
	// 	fmt.Println("the args are not exist!")
	// 	return
	// }

	// inputFile := os.Args[1]
	// outputFile := os.Args[2]
	inputFile := "sample.txt"
	outputFile := "result.txt"
	// if !strings.HasSuffix(outputFile, ".txt") || !strings.HasSuffix(inputFile, ".txt") {
	// 	fmt.Println("hdkfj")
	// 	return
	// }

	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("error in read file")
		return
	}

	// slice := strings.Fields(string(file))
	// fmt.Println(string(file))
	// fmt.Println(slice)
	str := functions.ApplayModifications(string(file))
	// slice = functions.CLeanSlice(slice)
	// fmt.Println(slice)
	// combined := strings.Join(slice, " ")
	byteSlice := []byte(str)
	erre := os.WriteFile(outputFile, byteSlice, 0o644)
	if erre != nil {
		fmt.Println("error in write file")
		return
	}

	// slice = functions.ApplayModifications(slice)
	// slc1 := strings.Join(functions.ApplayModifications(slice), " ")
	// slc2 := strings.Join(functions.ApplayModifications(slice), " ")

	// for slc1 != slc2 {
	// 	slice = functions.ApplayModifications(slice)
	// 	slc1 = strings.Join(functions.ApplayModifications(slice), " ")
	// 	slc2 = strings.Join(functions.ApplayModifications(slice), " ")
	// }

	// slice = functions.CLeanSlice(slice)
	// fmt.Println(slice)
	// combined := strings.Join(slice, " ")
	// byteSlice := []byte(combined)
	// erre := os.WriteFile(outputFile, byteSlice, 0o644)
	// if erre != nil {
	// 	fmt.Println("error in write file")
	// 	return
	// }
}
