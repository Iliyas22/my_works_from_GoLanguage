package main

import (
	"fmt"
)

func main() {
	names := make(map[string]int)
	names["Abylai"] = 20
	names["Duman"] = 21
	names["Iliyas"] = 22

	for name, age := range names {
		fmt.Printf("Name:%s = %d \n", name, age)
	}
	//Запись значений  в мапу схожа с массивом, только вместо индекса мы указываем ключ.
	//Итерация по мапе аналогична со срезами и массивами
	fmt.Println("Пустой элемент ", names["Bakbergen"])
	//Обращение к несуществующему элементу мапы вполне корректно, и будет
	//содержать в себе нулевое значение типа.

	newname, exists := names["Арман"]

	if !exists {
		fmt.Println("Нет такой имени")
	} else {
		fmt.Println(newname)
	}
	//В данном случае exists == false , т.к. Армана нету в нашем списке. Но если мы его
	//добавим, то сработает условие else потому что в таком случае будет exists ==
	//true

	//Нельзя дублировать ключи в мапе, это приведет к ошибке компиляции.
	//но можно перезаписывать
	names["Iliyas"] = 21
	fmt.Println(names["Iliyas"])

	delete(names, "Duman")
	fmt.Println("После функции Delete")
	for name, age := range names {
		fmt.Printf("Name:%s = %d \n", name, age)
	}
	//Есть функция - с его помощью можно удалять элементы мапа,указав в качестве
	//аргумента саму мапу и ключ.
}
