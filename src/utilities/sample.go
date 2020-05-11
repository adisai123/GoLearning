package utilities

import (
	"fmt"
	"runtime"
	"time"
)

const (
	A_aditya = 10
)

//This returns number of cpus
func currentCpuCount() {
	var num = runtime.NumCPU()
	fmt.Println("number of cpu:", num)
	fmt.Println(1, 2, 3, 4, 5, true, 1.7, 8, A_aditya)
}

func getSpeedAndCurrentTime() (x int, y time.Time) {
	var (
		speed int
		now   time.Time
	)
	speed, now = 10, time.Now()
	fmt.Println("current speed and time", speed, now)
	var (
		i  int
		y1 bool
	)
	_, _ = i, y1
	return speed, now
}
