package main

import (
	"errors"
	"fmt"
	"math"
)

const pi = math.Pi

func main() {
	PrintAreaAndPerimeter(2, 1)

}
func PrintAreaAndPerimeter(a, b int) {
	Area, err := CalculateAreaAndPerimeter(a, b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("Area of the Recatangle use this formula: S=A*B")
	fmt.Printf("%d=A and %d=B ", a, b)
	fmt.Println("\nAREA S=:", Area)
	fmt.Printf("-------------------------------------------------------------")

	Perimeter, err := CalculatePerimeter(a, b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("------------")
	fmt.Println("Perimeter of the Recatangle use this formula: P=2(A+B)")
	fmt.Printf("%d=A and %d=B ", a, b)
	fmt.Println("\nAREA P=:", Perimeter)
}

func CalculateAreaAndPerimeter(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return int(0), errors.New("Один из катетов не должно быть ниже нуля")
	}

	return a * b, nil

}
func CalculatePerimeter(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return int(0), errors.New("Один из катетов не должно быть ниже нуля")
	}
	return 2 * (a + b), nil
}
