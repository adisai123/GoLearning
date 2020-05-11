package utilities

import (
	"fmt"
	"strconv"
	"time"
)

type duration int64

func definedTypeEx() {
	var d1 duration = 20
	d1++
	var i1 int64 = 20
	i1++
	if i1 == int64(d1) {
		fmt.Println("defined types are same")
	}

	d, _ := time.ParseDuration("4h30m")
	fmt.Println(d.Minutes())
	type (
		Gram     int
		Kilogram Gram
		Ton      Kilogram
	)
	var (
		salt  Gram     = 10
		apple Kilogram = 1
		truck Ton      = 1
	)
	fmt.Println(salt, apple, truck)
	const min = 10
	const yy = 3.14 * min
	fmt.Println(yy)
	if n, err := strconv.Atoi("123123"); err == nil {
		fmt.Println(n)
	}

}
