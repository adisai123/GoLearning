package main

import (
	"fmt"
	a "mygo/goCamp/src/utilities"
)

func main() {
	hello := "hello "
	fmt.Println("hello", hello)

	complex := 0
	fmt.Println(complex)

	var x = 10
	var y int
	y = x
	x = 221
	fmt.Println("previous vlu", y)
	fmt.Println("current vlu", x)

	a.ExecuteAll()
	for i:= 0 ; i<10 ; i++{
		fmt.Println(100/i-9)
	}
}
