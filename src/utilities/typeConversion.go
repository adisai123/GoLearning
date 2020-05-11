package utilities

import "fmt"

func typeTesting() {
	speed := 100
	force := 1.9
	//speed = speed * int(force)
	speed = int(float64(speed) * force)
	fmt.Println(speed)
	orange := string(speed)
	fmt.Println("type conversion ", orange)
}
