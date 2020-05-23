package main

import "fmt"

func main() {
	m := map[string]int{
		"aditya": 10,
		"saish":  11,
	}

	n := []int{1, 2, 3, 4, 5}
	print(m, "map")
	print(n, "slice")
	modifit(m, n)
	print(m, "map modified")
	print(n, "slice modified")
}

func modifit(m map[string]int, n []int) {
	n[0] = 100
	m["aditya"] = 9
}

func print(m interface{}, msg string) {
	fmt.Println("print", msg, ":", m)
}
