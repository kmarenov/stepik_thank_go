package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	words := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		word := ""
		for i, c := range text {
			if i == 0 {
				word += strings.ToUpper(string(c))
			} else {
				word += strings.ToLower(string(c))
			}
		}
		words = append(words, word)
	}
	fmt.Println(strings.Join(words, " "))
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
