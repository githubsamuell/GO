package main

import (
	"fmt"
)

type Samuel interface {
	Guilherme()
}

type Lopes struct {
	Silva string
}

func (s Lopes) Guilherme() {
	fmt.Println(s.Silva + ", Is improving")
}

func main() {
	v := Lopes{"Samuel Guilherme Lopes Silva"}
	v.Guilherme()
}