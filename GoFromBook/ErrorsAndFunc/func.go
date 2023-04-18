package main

import "fmt"

const pi = 3.14

func main() {

	CircleArea(-2)

}
func CircleArea(radius int) {
	if radius <= 0 {
		fmt.Println("Raduis It shouldn't be less 0")
		return
	}

	Circle := CalculateArea(radius)
	fmt.Printf("Raduis is : %d \n", radius)
	fmt.Println("Use this formula S=Pi*R2 ")
	fmt.Println("Circle Area is a:", Circle)
}

func CalculateArea(raduis int) float32 {
	return float32(raduis) * float32(raduis) * pi

}
