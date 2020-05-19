package utilities

import "fmt"

func compareArray() {
	array1 := [...]int{1, 2, 3}
	array2 := [...]int{1, 2, 3}
	if array1 == array2 {
		fmt.Println("elemets")
	}
	mutiDimen := [3][4]int{
		{1, 2, 3, 4},
	}
	fmt.Println(mutiDimen)

	rate := [...]float64{
		9: 30, //  KEYED array: you can provide index while assigning value
	}

	fmt.Println(rate)
	const (
		ETH  = 9 - iota //9
		WAN             //8
		INCX            //7
	)
	fmt.Println(ETH, WAN, INCX)
	type (
		bookCase     [3]int
		cabinateCase [3]int
	)
	b := bookCase{2, 3, 4}
	c := cabinateCase{2, 3, 4}
	n := [3]int{2, 3, 4}
	//c==b not possible
	//b==n or c==n possile
	fmt.Println(b, c, n)
	fmt.Println(b == n)
	fmt.Println(c == n)
	fmt.Println(c == cabinateCase(b))

}
