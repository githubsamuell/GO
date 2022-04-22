package main

import (
	"fmt"
)

type Anyinterface struct {
	A, B float64
}

func (c Anyinterface) sum() float64 {
	return c.A + c.B
}

	func main() {
		obj := Anyinterface{5.5, 10.0}

		fmt.Println(obj.sum())

	}