package main

import (
	"fmt"
)

func FindZ(x float64) (z float64) {
	z = 1.0
	for i := 1; i < 10; i++ {
		//fmt.Println(i)
		z -= (z*z - x) / (2 * z)
		fmt.Println("Values from :", x, "~~", z)
	}
	return
}

func main() {
	for i := 1.0; i <= 4; i++ {
		fmt.Println("Value:", i)
		fmt.Println(FindZ(i))
	}
}
