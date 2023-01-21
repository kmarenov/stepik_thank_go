package main

import (
	"fmt"
	"math/rand"
)

// начало решения

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
        defer close(out)
        for {
            select {
            case out <- randomWord(5):
            case <-cancel:
                return
            }
        }
	}()
	return out
}

func isUnique(word string) bool {
	if len(word) < 2 {
		return true
	}

	chars := make(map[rune]struct{})

	for _, r := range word {
		_, ok := chars[r]; if ok {
			return false
		}
		chars[r] = struct{}{}
	}

	return true
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string  {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
            case word, ok := <-in:
				if !ok {
                    return
                }
				if !isUnique(word) {
					continue
				}
				select {
				case out <- word:
				case <-cancel:
					return
				}
            case <-cancel:
                return
            }
		}
	}()
	return out
}

func rev(word string) string {
	w := ""
	for i := len(word)-1; i >= 0; i-- {
		w += string(word[i])
	}
	return w
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for word := range in {
			select {
            case out <- rev(word):
            case <-cancel:
                return
            }
		}
	}()
	return out
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan string) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        for c1 != nil || c2 != nil {
            select {
            case val1, ok := <-c1:
                if ok {
                    select {
					case out <- val1:
					case <-cancel:
						return
					}
                } else {
                    c1 = nil
                }

            case val2, ok := <-c2:
                if ok {
                    select {
					case out <- val2:
					case <-cancel:
						return
					}
                } else {
                    c2 = nil
                }
            
			case <-cancel:
				return
			}
        }
    }()
    return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan string, n int) {
	for i := 0; i < n; i++ {
		select {
		case word, ok := <-in:
			if !ok {
				return
			}
			fmt.Println(rev(word) + " -> " + word)
		case <-cancel:
			return
		}
	}
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
}
