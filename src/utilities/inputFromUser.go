package utilities

import (
	"fmt"
	"os"
)

func main() {
	var args = os.Args
	s := "😆"
	fmt.Println(args[0])
	fmt.Println(s)
}
