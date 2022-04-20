package main

import (
	"fmt"
)

func returnsStrings(a, b string) (string, string) {
	return a, b
}

func main() {
	x, y := returnsStrings("Samuel", "Guilherme")
	fmt.Println(x, y);
}