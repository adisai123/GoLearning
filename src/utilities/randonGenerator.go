package utilities

import (
	"fmt"
	"math/rand"
	"time"
)

func randonGenerator() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Printf("%3d",
			rand.Intn(10))
	}
	fmt.Println()

}