package main

import (
	"fmt"
)

func main() {
	var code, lang string
	fmt.Scan(&code)

	// определите полное название языка по его коду
	// и запишите его в переменную `lang`
	switch {
		case code == "en":
			lang = "English"
		case code == "fr":
			lang = "French"
		case code == "ru" || code == "rus":
			lang = "Russian"
		default:
			lang = "Unknown"
	}

	fmt.Println(lang)
}