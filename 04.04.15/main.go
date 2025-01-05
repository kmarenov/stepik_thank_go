package main

import (
	"bufio"
	"fmt"
	mathrand "math/rand"
	"os"
	"path/filepath"
)

// алфавит планеты Нибиру
const alphabet = "aeiourtnsl"

// Census реализует перепись населения.
// Записи о рептилоидах хранятся в каталоге census, в отдельных файлах,
// по одному файлу на каждую букву алфавита.
// В каждом файле перечислены рептилоиды, чьи имена начинаются
// на соответствующую букву, по одному рептилоиду на строку.
type Census struct {
	dir     string
	count   int
	files   map[byte]*os.File
	writers map[byte]*bufio.Writer
}

// Count возвращает общее количество переписанных рептилоидов.
func (c *Census) Count() int {
	return c.count
}

// Add записывает сведения о рептилоиде.
func (c *Census) Add(name string) {
	first := name[0]
	writer := c.writers[first]
	_, err := writer.WriteString(name + "\n")
	if err != nil {
		panic(err)
	}
	c.count++
}

// Close закрывает файлы, использованные переписью.
func (c *Census) Close() {
	var anyErr error
	for _, w := range c.writers {
		if err := w.Flush(); anyErr == nil {
			anyErr = err
		}
	}
	for _, f := range c.files {
		if err := f.Close(); anyErr == nil {
			anyErr = err
		}
	}
	if anyErr != nil {
		panic(anyErr)
	}
}

// NewCensus создает новую перепись и пустые файлы
// для будущих записей о населении.
func NewCensus() *Census {
	makeDir := func(name string) {
		err := os.RemoveAll(name)
		if err != nil {
			panic(err)
		}
		err = os.Mkdir(name, 0755)
		if err != nil {
			panic(err)
		}
	}

	initWriters := func(dir string, letters []byte) (map[byte]*os.File, map[byte]*bufio.Writer) {
		files := make(map[byte]*os.File, len(letters))
		writers := make(map[byte]*bufio.Writer, len(letters))
		for _, let := range letters {
			path := filepath.Join(dir, fmt.Sprintf("%c.txt", let))
			file, err := os.Create(path)
			if err != nil {
				panic(err)
			}
			files[let] = file
			writers[let] = bufio.NewWriter(file)
		}
		return files, writers
	}

	const dir = "04.04.15/census"
	makeDir(dir)
	letters := []byte(alphabet)
	files, writers := initWriters(dir, letters)
	return &Census{dir: dir, files: files, writers: writers}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

var rand = mathrand.New(mathrand.NewSource(0))

// randomName возвращает имя очередного рептилоида.
func randomName(n int) string {
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(chars)
}

func main() {
	census := NewCensus()
	defer census.Close()
	for i := 0; i < 1024; i++ {
		reptoid := randomName(5)
		census.Add(reptoid)
	}
	fmt.Println(census.Count())
}
