package utilities

import "fmt"

func slices() {
	myArr := [5]int{2, 3, 4, 5, 6}
	myslices := myArr[2:5]
	fmt.Println("myarray", myArr)
	fmt.Println("myslices", myslices)
	myslices[0] = 100
	defer fmt.Println("myarray", myArr)
	defer fmt.Println("myslices", myslices)
	fmt.Println(cap(myslices), &myslices)
	myslices = append(myslices, 2, 10, 15, 234, 234, 234, 234)
	copiedSlice := []int{}
	copy(copiedSlice, myslices)
	defer fmt.Println("copied array", copiedSlice)
	fmt.Println(cap(myslices), &myslices)
	myslices = append(myslices, 234)
	fmt.Println(myslices, cap(myslices))
}
