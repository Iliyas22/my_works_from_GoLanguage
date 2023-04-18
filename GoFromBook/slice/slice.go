package main

import "fmt"

func main() {
	todoList := []string{
		"Купить билеты на кино",
		"Купить попкорн на двоих",
		"Купить шоколад Казахстанскую",
	}
	for i := range todoList {
		fmt.Println(todoList[i])
	}
	newTodoList := append(todoList, "Купить цветы")
	fmt.Printf("Длина=%d , Емкость=%d , у старого среза \n", len(todoList), cap(todoList))

	fmt.Printf("Длина=%d , Емкость=%d , у нового среза", len(newTodoList), cap(newTodoList))
	//Функция append() не изменяет переданный в аргументах срез, а возвращает новый
	for i := range newTodoList {
		fmt.Println(newTodoList[i])
	}
}
