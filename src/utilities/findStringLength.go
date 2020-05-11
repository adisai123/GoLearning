package utilities

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func evalLength() {
	currentLen := len("français")
	fmt.Println("Current Length ", currentLen)
	currentLen = utf8.RuneCountInString("français")
	fmt.Println("current length ", currentLen)
	stringing := "aditya"
	l := len(stringing)
	s := stringing + strings.Repeat("!", l)
	fmt.Println(s)
	s = strings.ToUpper(s)
	fmt.Println("upper case", s)
}
