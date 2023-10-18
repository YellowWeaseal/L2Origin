package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
		{"qwe\\\\5", "qwe\\\\\\\\\\"},
		{"abc3d2", "abcccdd"},
		{"a\\5b2c\\3", "a5bbc3"},
		{"\\5a", "5a"},
		{"\\", ""}, // Тест на обработку строки с одиночным символом "\\"
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			output, err := UnpackString(test.input)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if output != test.expected {
				t.Errorf("Input: %s, Expected: %s, Got: %s", test.input, test.expected, output)
			}
		})
	}

	// Тест на некорректную строку (с числом в начале)
	invalidInput := "123abc"
	_, err := UnpackString(invalidInput)
	if err == nil {
		t.Errorf("Expected an error for input: %s", invalidInput)
	}
}
