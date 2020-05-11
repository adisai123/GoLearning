package utilities

import "fmt"

func testPrintFormtting() {
	var abc string
	fmt.Printf("%q\n", abc)
	fmt.Printf("%T\n", abc)
	fmt.Printf("%d , \t %v\n", 123, 12)
}
