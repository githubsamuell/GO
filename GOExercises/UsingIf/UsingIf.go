package main

import (
	"fmt"
	"math"
)

func pow(a, b, x float64) float64 {
	y := math.Pow(a, b);
	if y < x {
		return x
	}

	return y;
}

func main() {
	fmt.Println(pow(5, 2, 30));
	fmt.Println(pow(5, 2, 20));
}