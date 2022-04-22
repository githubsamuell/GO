package main

import ("fmt")

func main() {

var s interface{} = 5

i, ok := s.(int)
fmt.Println(i, ok)

j, ok := s.(string)
fmt.Println(j, ok)

var k interface{} = "Samuel"

l, ok := k.(string)
fmt.Println(l, ok)

h, ok := k.(int)
fmt.Println(h, ok)

}
