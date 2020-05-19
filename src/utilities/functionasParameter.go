package utilities

import "fmt"

type myfun func(integer int) bool

func isOdd(i int) bool {
	return i%2 != 1
}

func filterSlice(slice []int, f myfun) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	displaySlice(result, "odd numbers list")
	return result
}

func displaySlice(slice []int, msg string) {
	fmt.Println(msg, slice)
}
func ExecuteFunctionParameterExample() {
	filterSlice([]int{1, 2, 3, 4, 5}, isOdd)
}
