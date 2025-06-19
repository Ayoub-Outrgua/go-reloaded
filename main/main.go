package main

import (
	"fmt"
	"os"
	"strings"

	"goreload"
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

	slice := goreload.CleanStr(file)

	fmt.Println(string(file))
	fmt.Println(slice)

	erre := os.WriteFile(outputFile, file, 0o644)
	if erre != nil {
		fmt.Println("error in write file")
		return
	}
}
