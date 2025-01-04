package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// TokenReader начитывает токены из источника
type TokenReader interface {
	// ReadToken считывает очередной токен
	// Если токенов больше нет, возвращает ошибку io.EOF
	ReadToken() (string, error)
}

// TokenWriter записывает токены в приемник
type TokenWriter interface {
	// WriteToken записывает очередной токен
	WriteToken(s string) error
}

// начало решения

// FilterTokens читает все токены из src и записывает в dst тех,
// кто проходит проверку predicate
func FilterTokens(dst TokenWriter, src TokenReader, predicate func(s string) bool) (n int, err error) {
	for {
		token, err := src.ReadToken()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return n, err
			} else {
				break
			}
		}
		if predicate(token) {
			err = dst.WriteToken(token)
			if err != nil {
				return n, err
			}
			n++
		}
	}
	return n, nil
}

// конец решения

type WordReader struct {
	scanner *bufio.Scanner
}

func (r *WordReader) ReadToken() (string, error) {
	if r.scanner.Scan() {
		return r.scanner.Text(), r.scanner.Err()
	}
	return "", io.EOF
}

func NewWordReader(text string) *WordReader {
	reader := strings.NewReader(text) // (1)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	return &WordReader{scanner: scanner}
}

type WordWriter struct {
	words int
}

func (w *WordWriter) WriteToken(s string) error {
	w.words++
	return nil
}

func (w *WordWriter) Words() int {
	return w.words
}

func NewWordWriter() *WordWriter {
	return &WordWriter{}
}

func main() {
	// Для проверки придется создать конкретные типы,
	// которые реализуют интерфейсы TokenReader и TokenWriter.

	// Ниже для примера используются NewWordReader и NewWordWriter,
	// но вы можете сделать любые на свое усмотрение.

	r := NewWordReader("go is awesome")
	w := NewWordWriter()
	predicate := func(s string) bool {
		return s != "is"
	}
	n, err := FilterTokens(w, r, predicate)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d tokens: %v\n", n, w.Words())
	// 2 tokens: [go awesome]
}
