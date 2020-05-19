package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	args := os.Args[1:]

	if len(args) <= 0 {
		fmt.Println("usage go run . [your name]")
		return
	}
	moods := [...]string{
		"happy 😀", "good 👍", "awesome 😎",
		"sad 😞", "bad 👎", "terrible 😩",
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))
	fmt.Println("%s feels %s \n", args[0], moods[n])
}
