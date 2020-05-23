package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	//res, _ := ioutil.ReadAll(resp.Body) 1 way
	//fmt.Println(string(res))

	//byte := make([]byte, 100000)  second way
	//resp.Body.Read(byte)
	//fmt.Println(string(byte))

	io.Copy(os.Stdout, resp.Body) //third way to print

}
