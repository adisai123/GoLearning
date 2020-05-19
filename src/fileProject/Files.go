package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("Provide a directory")
		return
	}
	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println("err", err)
		return
	}
	var bytes = make([]byte, 0, 300)
	for _, file := range files {
		name := file.Name()
		bytes = append(bytes, name...)
		bytes = append(bytes, '\n')
	}
	err = ioutil.WriteFile("myFile.name", bytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
