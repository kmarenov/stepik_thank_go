package main

import (
	"regexp"
	"strings"
	"testing"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
func slugify(src string) string {
	re := regexp.MustCompile(`[A-Za-z0-9-]+`)
	indices := re.FindAllStringIndex(src, -1)
	parts := make([]string, 0, len(indices))
	for _, idx := range indices {
		parts = append(parts, strings.ToLower(src[idx[0]:idx[1]]))
	}
	return strings.Join(parts, "-")
}

// конец решения

func Test(t *testing.T) {
	const phrase = "Go Is Awesome!"
	const want = "go-is-awesome"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}
