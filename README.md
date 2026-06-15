# Go Reloaded

> Zone 01 project — text transformation engine in Go

## 📚 Project Context
Peer-to-peer project focused on string manipulation, file I/O, and Go's standard library.

## 🎯 Learning Objectives
- Master Go's `strings`, `regexp`, and `strconv` packages
- Handle edge cases in text processing
- Write idiomatic Go without external dependencies

## 🛠️ Implementation
The program reads text, applies transformation rules (capitalization, punctuation spacing, article formatting), and outputs cleaned text.

## 🔍 Key Challenges Solved
| Challenge | My Solution |
|-----------|--------------|
| Handling nested transformations | Recursive parsing approach |
| Performance with large files | Buffered I/O |
| Edge cases (apostrophes, quotes) | State machine for quote tracking |

## 🧪 Run It
```bash
go run . sample.txt result.txt
