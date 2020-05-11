package utilities

import "fmt"

const (
	monday    = iota + 2
	tuesday   = iota + 2
	wednesday = iota
)

func iotaTesting() {
	fmt.Println("wed", wednesday)
}
