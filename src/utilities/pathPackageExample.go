package utilities

import (
	"fmt"
	"path"
)

func pathExample() {
	var dir, file string
	dir, file = path.Split("c/css/abc.css")
	fmt.Println("dir", dir)
	fmt.Println("file", file)
}
