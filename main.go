package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input> <output>")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := processText(string(data))

	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

// ================= PROCESS =================

func processText(text string) string {
	lines := strings.Split(text, "\n")
	var out []string

	for _, line := range lines {
		words := strings.Fields(line)

		for i := 0; i < len(words); i++ {

			// HEX
			if words[i] == "(hex)" && i > 0 {
				words[i-1] = hexToDecimal(words[i-1])
				words[i] = ""
			}

			// BIN
			if words[i] == "(bin)" && i > 0 {
				words[i-1] = binToDecimal(words[i-1])
				words[i] = ""
			}

			// UP / LOW / CAP
			if i > 0 {
				switch words[i] {
				case "(up)":
					words[i-1] = strings.ToUpper(words[i-1])
					words[i] = ""

				case "(low)":
					words[i-1] = strings.ToLower(words[i-1])
					words[i] = ""

				case "(cap)":
					words[i-1] = capitalize(words[i-1])
					words[i] = ""
				}
			}

			// (up, N) / (low, N) / (cap, N)
			if (strings.HasPrefix(words[i], "(up,") ||
				strings.HasPrefix(words[i], "(low,") ||
				strings.HasPrefix(words[i], "(cap,")) && i+1 < len(words) {

				numStr := strings.Trim(words[i+1], ")")
				var n int
				fmt.Sscanf(numStr, "%d", &n)

				for j := 1; j <= n && i-j >= 0; j++ {
					switch {
					case strings.HasPrefix(words[i], "(up"):
						words[i-j] = strings.ToUpper(words[i-j])
					case strings.HasPrefix(words[i], "(low"):
						words[i-j] = strings.ToLower(words[i-j])
					case strings.HasPrefix(words[i], "(cap"):
						words[i-j] = capitalize(words[i-j])
					}
				}

				words[i], words[i+1] = "", ""
			}
		}

		line := strings.Join(clean(words), " ")

		line = fixPunctuation(line)
		line = fixQuotes(line)
		line = fixArticles(line)

		out = append(out, line)
	}

	return strings.Join(out, "\n")
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

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
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

		if w == "..." || w == "!?" {
			res = append(res, w)
			continue
		}

		if len(w) == 1 && strings.ContainsRune(punct, rune(w[0])) {
			if len(res) > 0 {
				res[len(res)-1] += w
			}
			continue
		}

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

// ================= QUOTES =================

func fixQuotes(s string) string {
	var res []rune
	in := false

	for _, r := range s {
		if r == '\'' {
			if !in {
				res = append(res, r)
				in = true
			} else {
				res = append(res, r)
				in = false
			}
			continue
		}

		if in && unicode.IsSpace(r) {
			continue
		}

		res = append(res, r)
	}

	return string(res)
}

// ================= ARTICLE RULE =================

func fixArticles(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" {
			next := strings.ToLower(words[i+1])
			if len(next) > 0 && (strings.ContainsRune("aeiouh", rune(next[0]))) {
				words[i] = "an"
			}
		}
	}

	return strings.Join(words, " ")
}
