// начало решения

// submitWords отправляет слова на подсчет

// countWords считает цифры в словах

// fillStats готовит итоговую статистику

// конец решения

package main

import (
	"fmt"
	"strings"
	"unicode"
)

// nextFunc возвращает следующее слово из генератора
type nextFunc func() string

// counter хранит количество цифр в каждом слове.
// ключ карты - слово, а значение - количество цифр в слове.
type counter map[string]int

// pair хранит слово и количество цифр в нем
type pair struct {
	word  string
	count int
}

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWords(next nextFunc) counter {
	pending := make(chan string)
	go submitWords(next, pending)

	counted := make(chan pair)
	go countWords(pending, counted)

	return fillStats(counted)
}

// начало решения

// submitWords отправляет слова на подсчет
func submitWords(next nextFunc, pending chan string) {
	// отправляет слова на подсчет
	go func() {
		// Пройдите по словам и отправьте их
		// в канал pending
		for {
			word := next()
			pending <- word
			if word == "" {
				break
			}
		}
	}()
}

// countWords считает цифры в словах
func countWords(pending chan string, counted chan pair) {
	// считает цифры в словах
	go func() {
		// Считайте слова из канала pending,
		// посчитайте количество цифр в каждом,
		// и запишите его в канал counted
		for {
			word := <-pending
			count := countDigits(word)
			counted <- pair{word, count}
			if word == "" {
				break
			}
		}
	}()
}

// fillStats готовит итоговую статистику
func fillStats(counted chan pair) counter {
	// Считайте значения из канала counted
	// и заполните stats.

	// В результате stats должна содержать слова
	// и количество цифр в каждом.
	stats := counter{}
	for {
		count := <-counted
		if count.word == "" {
			break
		}
		stats[count.word] = count.count
	}

	return stats
}

// конец решения

// countDigits возвращает количество цифр в строке
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats печатает слова и количество цифр в каждом
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// wordGenerator возвращает генератор, который выдает слова из фразы
func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}
