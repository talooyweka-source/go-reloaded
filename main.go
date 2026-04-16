package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input> <output>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(data)

	result := processText(text)

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
func processText(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		if words[i] == "(hex)" && i > 0 {
			words[i-1] = hexToDecimal(words[i-1])
			words[i] = ""
		}

		if words[i] == "(bin)" && i > 0 {
			words[i-1] = binToDecimal(words[i-1])
			words[i] = ""
		}
	}

	return strings.Join(clean(words), " ")
}
func hexToDecimal(s string) string {
	var n int
	fmt.Sscanf(s, "%x", &n)
	return fmt.Sprintf("%d", n)
}

func binToDecimal(s string) string {
	var n int
	fmt.Sscanf(s, "%b", &n)
	return fmt.Sprintf("%d", n)
}

func clean(words []string) []string {
	var res []string
	for _, w := range words {
		if w != "" {
			res = append(res, w)
		}
	}
	return res
}
