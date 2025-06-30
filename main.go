package main

import (
	"fmt"
	"os"
	"strings"

	"goreload/functions"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: the arguments are not valid!")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if !strings.HasSuffix(outputFile, ".txt") || !strings.HasSuffix(inputFile, ".txt") {
		fmt.Println("Error: one of the arguments is not a text file!")
		return
	}

	file, readError := os.ReadFile(inputFile)
	if readError != nil {
		fmt.Println("File read error!", readError)
		return
	}

	str := functions.ApplayModifications(string(file))

	sliceOfByte := []byte(str)
	writeError := os.WriteFile(outputFile, sliceOfByte, 0o644)
	if writeError != nil {
		fmt.Println("File write error!", writeError)
		return
	}
}
