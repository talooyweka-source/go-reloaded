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

	result := process(string(data))

	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

// ================= CORE =================

func process(text string) string {
	lines := strings.Split(text, "\n")
	var out []string

	for _, line := range lines {
		words := strings.Fields(line)

		for i := 0; i < len(words); i++ {

			// HEX
			if words[i] == "(hex)" && i > 0 {
				words[i-1] = toHex(words[i-1])
				words[i] = ""
			}

			// BIN
			if words[i] == "(bin)" && i > 0 {
				words[i-1] = toBin(words[i-1])
				words[i] = ""
			}

			// SIMPLE CASES
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

			// (up, n), (low, n), (cap, n)
			if i+1 < len(words) &&
				(strings.HasPrefix(words[i], "(up,") ||
					strings.HasPrefix(words[i], "(low,") ||
					strings.HasPrefix(words[i], "(cap,")) {

				var n int
				fmt.Sscanf(strings.Trim(words[i+1], ")"), "%d", &n)

				for j := 0; j < n && i-1-j >= 0; j++ {
					switch {
					case strings.HasPrefix(words[i], "(up"):
						words[i-1-j] = strings.ToUpper(words[i-1-j])
					case strings.HasPrefix(words[i], "(low"):
						words[i-1-j] = strings.ToLower(words[i-1-j])
					case strings.HasPrefix(words[i], "(cap"):
						words[i-1-j] = capitalize(words[i-1-j])
					}
				}

				words[i], words[i+1] = "", ""
			}
		}

		line := strings.Join(clean(words), " ")
		line = fixPunct(line)
		line = fixQuotes(line)
		line = fixArticles(line)

		out = append(out, line)
	}

	return strings.Join(out, "\n")
}

// ================= HELPERS =================

func toHex(s string) string {
	var n int
	fmt.Sscanf(s, "%x", &n)
	return fmt.Sprintf("%d", n)
}

func toBin(s string) string {
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

func clean(w []string) []string {
	var r []string
	for _, v := range w {
		if v != "" {
			r = append(r, v)
		}
	}
	return r
}

// ================= PUNCTUATION (FIXED) =================

func fixPunct(s string) string {
	// remove space before punctuation
	s = strings.ReplaceAll(s, " ,", ",")
	s = strings.ReplaceAll(s, " .", ".")
	s = strings.ReplaceAll(s, " !", "!")
	s = strings.ReplaceAll(s, " ?", "?")
	s = strings.ReplaceAll(s, " :", ":")
	s = strings.ReplaceAll(s, " ;", ";")

	// fix grouped punctuation
	s = strings.ReplaceAll(s, ". . .", "...")
	s = strings.ReplaceAll(s, "! ?", "!?")

	return s
}

// ================= QUOTES =================

func fixQuotes(s string) string {
	var result []rune
	inQuote := false
	spaceBuffer := false

	for _, r := range s {
		if r == '\'' {
			inQuote = !inQuote
			result = append(result, r)
			continue
		}

		if inQuote {
			if unicode.IsSpace(r) {
				spaceBuffer = true
				continue
			}
			if spaceBuffer && len(result) > 0 && result[len(result)-1] != '\'' {
				result = append(result, ' ')
			}
			spaceBuffer = false
		}

		result = append(result, r)
	}

	return string(result)
}

// ================= ARTICLES (FIXED) =================

func fixArticles(s string) string {
	w := strings.Fields(s)

	for i := 0; i < len(w)-1; i++ {
		if strings.ToLower(w[i]) == "a" {
			n := strings.ToLower(w[i+1])
			if len(n) > 0 && strings.ContainsRune("aeiouh", rune(n[0])) {
				if w[i] == "A" {
					w[i] = "An"
				} else {
					w[i] = "an"
				}
			}
		}
	}

	return strings.Join(w, " ")
}
