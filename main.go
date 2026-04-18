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
	lines := strings.Split(text, "\n")
	var result []string

	for _, line := range lines {
		words := strings.Fields(line)

		for i := 0; i < len(words); i++ {

			if words[i] == "(hex)" && i > 0 {
				words[i-1] = hexToDecimal(words[i-1])
				words[i] = ""
			}

			if words[i] == "(bin)" && i > 0 {
				words[i-1] = binToDecimal(words[i-1])
				words[i] = ""
			}

			if words[i] == "(up)" && i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words[i] = ""
			}

			if words[i] == "(low)" && i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words[i] = ""
			}

			if words[i] == "(cap)" && i > 0 {
				words[i-1] = capitalize(words[i-1])
				words[i] = ""
			}
		}

		cleaned := strings.Join(clean(words), " ")
		result = append(result, cleaned)
	}

	return strings.Join(result, "\n")
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

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
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
