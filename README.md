# # Go Reloaded

## Description

Go Reloaded is a command-line text processing tool written in Go. It reads a text file, applies a series of transformations, and writes the modified result into another file. This project focuses on string manipulation, file handling, and clean code practices.

## Features

The program supports several transformations. It can convert numbers using `(hex)` to transform a hexadecimal value into decimal (for example, `1E (hex)` becomes `30`) and `(bin)` to convert binary values into decimal (for example, `10 (bin)` becomes `2`). It also supports case transformations such as `(up)` to convert the previous word to uppercase, `(low)` to convert it to lowercase, and `(cap)` to capitalize it. These transformations can also be applied to multiple words using formats like `(up, n)`, `(low, n)`, or `(cap, n)`, where `n` represents the number of previous words affected (e.g., `this is amazing (up, 2)` becomes `this is AMAZING`).

In addition, the tool formats punctuation by removing spaces before punctuation marks and ensuring proper spacing after them, while correctly handling grouped punctuation such as `...` and `!?`. It also processes quotes by removing unnecessary spaces inside `' '` and attaching them properly to the enclosed words (e.g., `' hello world '` becomes `'hello world'`). Furthermore, it improves grammar by converting `a` to `an` when followed by a vowel or the letter `h`, as in `a apple` becoming `an apple`.

## Usage

The program is executed using the command:
`go run . input.txt output.txt`

## Examples

For example, the input `Simply add 42 (hex) and 10 (bin)` produces the output `Simply add 66 and 2`. Another example is `I was sitting over there , and then BAMM !!`, which becomes `I was sitting over there, and then BAMM!!`. Similarly, `This is so exciting (up, 2)` is transformed into `This is SO EXCITING`.

## Project Structure

The project consists of files such as `main.go` (entry point), processing and helper logic files, `sample.txt` for input examples, `result.txt` for outputs, and the `README.md` file.

## Concepts Used

This project demonstrates the use of file I/O operations (`os.ReadFile`, `os.WriteFile`), string manipulation with the `strings` and `unicode` packages, command-line argument handling, and modular programming practices.

## Notes

The program expects exactly two arguments. The input file must exist, and the output file will be created or overwritten if it already exists.

## Author

Talo Oyweka
