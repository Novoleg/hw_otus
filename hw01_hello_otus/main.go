package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	inputString := "Hello, OTUS!"
	reverseResult := stringutil.Reverse(inputString)
	fmt.Println(reverseResult)
}
