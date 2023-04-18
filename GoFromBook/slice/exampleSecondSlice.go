package main

import "fmt"

func main() {
	todoList := [4]string{
		"Купить молоко",
		"Купить шоколадку",
		"Потренироваться",
		"Всегда быть на позитиве!",
	}
	tasks := []string{}
	fmt.Println(tasks == nil)

	tasks = todoList[1:4]
	for i := range tasks {
		fmt.Println(tasks[i])
	}
	// After Change List
	todoList[1] = "Быть оптимистом"
	fmt.Printf("\n Ater Change Elements\n")
	for i := range tasks {
		fmt.Println(tasks[i])
	}
}
