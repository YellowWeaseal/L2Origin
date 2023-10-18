package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func UnpackString(s string) (r string, err error) {
	if _, err := strconv.Atoi(s); err == nil {
		return r, errors.New("invalid string")
	}

	var prev rune
	var escaped bool
	var b strings.Builder
	for _, char := range s {
		if unicode.IsDigit(char) && !escaped {
			m := int(char - '0')
			r := strings.Repeat(string(prev), m-1)
			b.WriteString(r)
		} else {
			escaped = string(char) == "\\" && string(prev) != "\\"
			if !escaped {
				b.WriteRune(char)
			}
			prev = char
		}
	}

	return b.String(), err
}
func main() {
	input1 := "qwe\\4\\5"
	input2 := "qwe\\45"
	input3 := "qwe\\\\5"
	input4 := "a4bc2d5e"
	input5 := "abcd"
	input6 := "45"
	input7 := ""

	fmt.Println(UnpackString(input1))
	fmt.Println(UnpackString(input2))
	fmt.Println(UnpackString(input3))
	fmt.Println(UnpackString(input4))
	fmt.Println(UnpackString(input5))
	fmt.Println(UnpackString(input6))
	fmt.Println(UnpackString(input7))
}
