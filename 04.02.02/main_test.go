package main

import (
	"strings"
	"testing"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
func slugify(src string) string {
	slug := strings.Builder{}
	for _, c := range src {
		if strings.ContainsAny(string(c), "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-") {
			slug.WriteString(strings.ToLower(string(c)))
		} else {
			slug.WriteRune(' ')
		}
	}
	return strings.Join(strings.Fields(slug.String()), "-")
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
