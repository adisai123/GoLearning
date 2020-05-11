package utilities

import (
	"fmt"
	"strconv"
)

func exampleStrconv() {
	var mystring string = "11.1s"
	myconvertedNumber, err := strconv.ParseFloat(mystring, 64)
	fmt.Println("parsefloat", myconvertedNumber, err.Error())

}
