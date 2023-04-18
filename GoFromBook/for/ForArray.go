package main

import "fmt"

func main() {

	var arr [3]int
	arr2 := ArrayCalculate(arr)

	fmt.Println("Array 1:", arr)
	fmt.Println("Array2:", arr2)
	for index, value := range arr2 {
		fmt.Printf(" Index: %d , value: %d \n", index, value)

	}
	// Пример второй
	var arrString = [3]string{
		"Apple",
		"Samsung",
	}
	arrString[2] = "Xiaomi"
	fmt.Printf("К-во элементов в массиве: %d \n", len(arrString))
	for item := range arrString {
		fmt.Printf("Index: %d value: %s \n", item, arrString[item])
	}
	fmt.Println(cap(arrString))
	// Конец примера
}
func ArrayCalculate(arr [3]int) [3]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}
