package utilities

import (
	"fmt"
	"sort"
	"strconv"
)

func stringEx() {
	var stringValue = `aditya
	\n,malpani`
	fmt.Println(stringValue)
	//concatanitation
	var stringconcanated = stringValue + "what a fun"
	fmt.Println(stringconcanated)
	stringSlice := []string{"asd", "asd", stringValue}
	mysorted := sort.StringSlice(stringSlice)
	fmt.Println(mysorted)
	eq := "1 + 2 ="
	sum := 1 + 1
	fmt.Println(eq + strconv.Itoa(sum))

}
