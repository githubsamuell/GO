package main

import (
	"fmt"
)

func main() {
	var a, b int = 10, 20;
	fmt.Println(a, b);
	var p, u = &a, &b;
	fmt.Println(*p, *u);
	*p = 200
	*u = 400
	fmt.Println(a, b);
	return ;

}