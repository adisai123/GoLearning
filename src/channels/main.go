package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	site = []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
		"http://faceboosk2.com",
	}
)

func main() {
	c := make(chan string)
	for _, link := range site {
		go checkLink(link, c)

	}

	for i := 0; i < 1000; i++ {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(<-c, c)
		}()
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	fmt.Println("checking for ling", l)
	for i := 0; i < 10000000; i++ {
		fmt.Print(" ")
	}
	if err != nil {
		fmt.Println(l, "might be down")
		c <- l
		return
	}
	fmt.Println(l, "is up")
	c <- l
	fmt.Println("testing......", l)

	for i := 0; i < 1000000; i++ {
		fmt.Print(" ")
	}
}
