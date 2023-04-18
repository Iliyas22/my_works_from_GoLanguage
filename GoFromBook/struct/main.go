package main

import "fmt"

type employee struct {
	name   string
	gender string
	age    int
	salary int
}

func newEmployee(name string, gender string, age int, salary int) employee {
	return employee{
		name:   name,
		gender: gender,
		age:    age,
		salary: salary,
	}
}

func (e employee) getInfo() string {
	return fmt.Sprintf("Сотрудник %s\n Пол %s \n age: %d \n salary : %d \n", e.name, e.gender, e.age, e.salary)

}
func (e *employee) setName(name string) {
	e.name = name
}
func main() {
	employee1 := newEmployee("Iliyas", "M", 22, 1000000)
	employee2 := newEmployee("Ilirron", "M", 23, 10000000)

	employee1.setName("Ino")

	fmt.Println("\n", employee1.getInfo())
	fmt.Printf("Name: %s \n Gender: %s\n Salary: %d\n  Age:%d\n", employee2.name, employee2.gender, employee2.salary, employee2.age)
	fmt.Println("\n", employee2.getInfo())

}
