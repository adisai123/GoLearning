package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var ao int = 10

var regx *regexp.Regexp = regexp.MustCompile(`[^ a-z | A-Z | 0-9]+`) //costly operation
func main() {

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		str := regx.ReplaceAllString(in.Text(), "")
		fmt.Println(str)
	}
	if err := in.Err(); err != nil {
		fmt.Println(err)
	}

}
