package main

// не удаляйте импорты, они используются при проверке
import (
    "fmt"
    "math/rand"
    "os"
    "strings"
    "testing"
)

type Words struct {
    c map[string]int
}

func MakeWords(s string) Words {
	words := strings.Fields(s)
    size := len(words) / 2
    if size > 10000 {
        size = 10000
    }
    c := make(map[string]int, size)

	for i, word := range words {
		if _, ok := c[word]; !ok {
			c[word] = i
		}
	}

    return Words{c}
}

func (w Words) Index(word string) int {
	i, ok := w.c[word]
	if ok {
		return i
	}

    return -1
}
