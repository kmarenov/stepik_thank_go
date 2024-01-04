package main

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

// начало решения

// prettify возвращает отформатированное
// строковое представление карты
func prettify(m map[string]int) string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	b := strings.Builder{}
	b.WriteRune('{')
	if len(m) == 1 {
		for _, k := range keys {
			b.WriteRune(' ')
			b.WriteString(k)
			b.WriteString(": ")
			b.WriteString(strconv.Itoa(m[k]))
			b.WriteRune(' ')
		}
	}
	if len(m) > 1 {
		b.WriteString("\n")
		for _, k := range keys {
			b.WriteString("    ")
			b.WriteString(k)
			b.WriteString(": ")
			b.WriteString(strconv.Itoa(m[k]))
			b.WriteString(",\n")
		}
	}
	b.WriteRune('}')
	return b.String()
}

// конец решения

func Test(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	const want = "{\n    one: 1,\n    three: 3,\n    two: 2,\n}"
	got := prettify(m)
	if got != want {
		t.Errorf("%v\ngot:\n%v\n\nwant:\n%v", m, got, want)
	}
}
