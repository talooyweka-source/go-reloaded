package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
				words[i-1] = toBase(words[i-1], 16)
				words[i] = ""
			}

			// BIN
			if words[i] == "(bin)" && i > 0 {
				words[i-1] = toBase(words[i-1], 2)
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

			// (up, n) / (low, n) / (cap, n)
			if strings.HasPrefix(words[i], "(") && i+1 < len(words) {
				if strings.Contains(words[i], ",") {

					action := words[i]
					numStr := strings.Trim(words[i+1], ")")
					n, _ := strconv.Atoi(numStr)

					for j := 0; j < n && i-1-j >= 0; j++ {
						switch {
						case strings.HasPrefix(action, "(up"):
							words[i-1-j] = strings.ToUpper(words[i-1-j])
						case strings.HasPrefix(action, "(low"):
							words[i-1-j] = strings.ToLower(words[i-1-j])
						case strings.HasPrefix(action, "(cap"):
							words[i-1-j] = capitalize(words[i-1-j])
						}
					}

					words[i], words[i+1] = "", ""
				}
			}
		}

		cleaned := clean(words)

		line = strings.Join(cleaned, " ")
		line = fixPunct(line)
		line = fixQuotes(line)
		line = fixArticles(line)

		out = append(out, line)
	}

	return strings.Join(out, "\n")
}

// ================= HELPERS =================

func toBase(s string, base int) string {
	var n int64

	if base == 16 {
		fmt.Sscanf(s, "%x", &n)
	} else if base == 2 {
		fmt.Sscanf(s, "%b", &n)
	}

	return strconv.FormatInt(n, 10)
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

// ================= PUNCTUATION =================

func fixPunct(s string) string {
	punct := ".,!?:;"
	var res []rune

	for i := 0; i < len(s); i++ {
		c := rune(s[i])

		if strings.ContainsRune(punct, c) {

			// remove space before punctuation
			if len(res) > 0 && res[len(res)-1] == ' ' {
				res = res[:len(res)-1]
			}

			res = append(res, c)

			// always add space after punctuation (safe rule)
			res = append(res, ' ')
			continue
		}

		res = append(res, c)
	}

	return strings.TrimSpace(string(res))
}

// ================= QUOTES =================

func fixQuotes(s string) string {
	var res []rune
	in := false

	for _, r := range s {
		if r == '\'' {
			in = !in
			res = append(res, r)
			continue
		}

		if in && r == ' ' {
			continue
		}

		res = append(res, r)
	}

	return string(res)
}

// ================= ARTICLES =================

func fixArticles(s string) string {
	w := strings.Fields(s)

	for i := 0; i < len(w)-1; i++ {
		if strings.ToLower(w[i]) == "a" {
			next := strings.ToLower(w[i+1])
			if len(next) > 0 && strings.ContainsRune("aeiouh", rune(next[0])) {
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
