package utilities

import (
	"fmt"
	"strconv"
)

func stringEx() {
	var stringValue = `aditya
	\n,malpani`
	fmt.Println(stringValue)
	//concatanitation
	var stringconcanated = stringValue + "what a fun"
	fmt.Println(stringconcanated)

	eq := "1 + 2 ="
	sum := 1 + 1
	fmt.Println(eq + strconv.Itoa(sum))

}
