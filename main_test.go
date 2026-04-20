 package main

import "testing"

func TestProcess(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		// ===== HEX =====
		{"1E (hex)", "30"},
		{"FF (hex)", "255"},

		// ===== BIN =====
		{"10 (bin)", "2"},
		{"1010 (bin)", "10"},

		// ===== BASIC CASE =====
		{"hello (up)", "HELLO"},
		{"HELLO (low)", "hello"},
		{"world (cap)", "World"},

		// ===== MULTI CASE =====
		{"this is amazing (up, 2)", "this IS AMAZING"},
		{"THIS IS BORING (low, 3)", "this is boring"},
		{"welcome to the jungle (cap, 4)", "Welcome To The Jungle"},

		// ===== PUNCTUATION =====
		{"hello , world !", "hello, world!"},
		{"wait ... what ?", "wait... what?"},
		{"no way !?", "no way!?"},
		{"stop ; now :", "stop; now:"},

		// ===== QUOTES =====
		{"' hello world '", "'hello world'"},
		{"I am ' amazing '", "I am 'amazing'"},

		// ===== ARTICLES =====
		{"a apple", "an apple"},
		{"a elephant", "an elephant"},
		{"a house", "a house"},
		{"A orange", "An orange"},

		// ===== MIXED =====
		{
			"I have 1E (hex) apples , and 10 (bin) oranges !",
			"I have 30 apples, and 2 oranges!",
		},
		{
			"this is so cool (up, 2) , right ?",
			"this is SO COOL, right?",
		},
		{
			"' hello world ' is a example",
			"'hello world' is an example",
		},
	}

	for _, test := range tests {
		result := process(test.input)
		if result != test.expected {
			t.Errorf("\nInput: %q\nExpected: %q\nGot: %q\n",
				test.input, test.expected, result)
		}
	}
}
