package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(len(strings.Fields(os.Args[1])))
}
