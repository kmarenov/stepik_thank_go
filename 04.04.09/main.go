package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
)

// начало решения

type RandReader struct {
	b   []byte
	i   int
	max int
}

func (r *RandReader) Read(p []byte) (n int, err error) {
	if r.i >= r.max {
		return 0, io.EOF
	}

	res := make([]byte, 0, len(p))
	for {
		rand.Read(r.b)
		res = append(res, r.b...)
		if len(res) >= len(p) {
			res = res[:len(p)]
			break
		}
	}

	if r.i+len(p) > r.max {
		res = res[:r.max-r.i]
	}

	n = copy(p, res)
	r.i += n

	return n, nil
}

func NewReader(max int) *RandReader {
	return &RandReader{
		b:   make([]byte, 16),
		i:   0,
		max: max,
	}
}

// RandomReader создает читателя, который возвращает случайные байты,
// но не более max штук
func RandomReader(max int) io.Reader {
	return NewReader(max)
}

// конец решения

func main() {
	rnd := RandomReader(8888)
	rd := bufio.NewReaderSize(rnd, 20)
	i := 0
	for {
		b, err := rd.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", b)

		i++
		fmt.Println("")
		fmt.Println("")
		fmt.Println(i)
		fmt.Println("")
		fmt.Println("")
	}
	fmt.Println()
	// 1 148 253 194 250
	// (значения могут отличаться)
}
