package main

import (
	"fmt"
	"strings"
)

type book struct{
	name string
}
func (b book) String() string{
	str := []string{"sd","asdasd","asdsd"}
	var s  strings.Builder
	for _, str := range str {
		s.WriteString(str)
	}
	return fmt.Sprintf("book.name%s",s)
}
func main()  {
	b := book { "Ramayan"}
	fmt.Println("book",b)
}