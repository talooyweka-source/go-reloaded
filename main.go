
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

// ================= PROCESS =================

func processText(text string) string {
	lines := strings.Split(text, "\n")
	var result []string

	for _, line := range lines {
		words := strings.Fields(line)

		for i := 0; i < len(words); i++ {

			// HEX
			if words[i] == "(hex)" && i > 0 {
				words[i-1] = hexToDecimal(words[i-1])
				words[i] = ""
				continue
			}

			// BIN
			if words[i] == "(bin)" && i > 0 {
				words[i-1] = binToDecimal(words[i-1])
				words[i] = ""
				continue
			}

			// (up, N)
			if strings.HasPrefix(words[i], "(up,") && i+1 < len(words) {
				numStr := strings.TrimSuffix(words[i+1], ")")
				var n int
				fmt.Sscanf(numStr, "%d", &n)

				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}

				words[i] = ""
				words[i+1] = ""
				continue
			}

			// (low, N)
			if strings.HasPrefix(words[i], "(low,") && i+1 < len(words) {
				numStr := strings.TrimSuffix(words[i+1], ")")
				var n int
				fmt.Sscanf(numStr, "%d", &n)

				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}

				words[i] = ""
				words[i+1] = ""
				continue
			}

			// (cap, N)
			if strings.HasPrefix(words[i], "(cap,") && i+1 < len(words) {
				numStr := strings.TrimSuffix(words[i+1], ")")
				var n int
				fmt.Sscanf(numStr, "%d", &n)

				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = capitalize(words[i-j])
				}

				words[i] = ""
				words[i+1] = ""
				continue
			}

			// (up)
			if words[i] == "(up)" && i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words[i] = ""
				continue
			}

			// (low)
			if words[i] == "(low)" && i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words[i] = ""
				continue
			}

			// (cap)
			if words[i] == "(cap)" && i > 0 {
				words[i-1] = capitalize(words[i-1])
				words[i] = ""
				continue
			}
		}

		cleaned := strings.Join(clean(words), " ")
		result = append(result, fixPunctuation(cleaned))
	}

	return strings.Join(result, "\n")
}

// ================= CONVERSIONS =================

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

// ================= CLEAN =================

func clean(words []string) []string {
	var res []string
	for _, w := range words {
		if w != "" {
			res = append(res, w)
		}
	}
	return res
}

// ================= PUNCTUATION =================

func fixPunctuation(line string) string {
	punct := ".,!?:;"

	words := strings.Fields(line)
	var res []string

	for i := 0; i < len(words); i++ {
		w := words[i]

		// grouped punctuation
		if w == "..." || w == "!?" {
			res = append(res, w)
			continue
		}

		// single punctuation token
		if len(w) == 1 && strings.ContainsRune(punct, rune(w[0])) {
			if len(res) > 0 {
				res[len(res)-1] += w
			} else {
				res = append(res, w)
			}
			continue
		}

		// split trailing punctuation (hello,)
		last := w[len(w)-1]
		if strings.ContainsRune(punct, rune(last)) && len(w) > 1 {
			res = append(res, w[:len(w)-1])
			res = append(res, string(last))
			continue
		}

		res = append(res, w)
	}

	var final []string
	for i := 0; i < len(res); i++ {
		if i > 0 && len(res[i]) == 1 && strings.ContainsRune(punct, rune(res[i][0])) {
			final[len(final)-1] += res[i]
		} else {
			final = append(final, res[i])
		}
	}

	return strings.Join(final, " ")
}
