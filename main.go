package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	if len (os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output. txt")
	return
    }

    inputFile := os.Args[1]
    outputFile := os.Args[2]

    content, err := os.ReadFile(inputFile)
    if err != nil {
	fmt.Println(err)
	return 
    }

    result := processText(string(content))

    err = os.WriteFile(outputFile, []byte(result), 0644)
    if err != nil {
	    fmt.Println(err)
	    return
    }
}

func processText(text string) string {
	words := strings.Fields(text)

	words = handleHexBin(words)
	return strings.Join(words, " ")
}

