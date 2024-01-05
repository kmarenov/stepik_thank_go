package main

import (
	"fmt"
	"strings"
	"unicode"
)

// начало решения

func slugify(src string) string {
	length := len(src)
	builder := strings.Builder{}
	builder.Grow(length)
	add_sep := false
	is_first := true
	for i := 0; i < length; i++ {
		c := src[i]
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-' {
			if add_sep && !is_first {
				builder.WriteByte('-')
			}
			builder.WriteByte(c)
			is_first = false
			add_sep = false
		} else if c >= 'A' && c <= 'Z' {
			if add_sep && !is_first {
				builder.WriteByte('-')
			}
			builder.WriteByte(byte(unicode.ToLower(rune(c))))
			is_first = false
			add_sep = false
		} else {
			add_sep = true
		}
	}
	return builder.String()
}

// конец решения

func main() {
	const phrase = "A 100x Investment (2019)"
	//const phrase = "Go_Is_Awesome"
	slug := slugify(phrase)
	fmt.Println(slug)
	// a-100x-investment-2019
}
