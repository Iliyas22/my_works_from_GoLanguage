package main

import (
	"errors"
	"fmt"
	"math"
)

const pii = math.Pi

func main() {
	PrintArea(0)
	//PrintArea(-2)

}
func PrintArea(radius int) {
	CircleAreaa, err := CalculateAreaa(radius)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Radius: %d cm. \n", radius)
	fmt.Println("Circle Area is:", CircleAreaa)
}

func CalculateAreaa(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Radius shouldn't be less than 0! \n Error!")
	}

	return float32(radius) * float32(radius) * pii, nil
}
