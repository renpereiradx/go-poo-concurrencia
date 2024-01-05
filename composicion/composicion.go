package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	userName string
}

func (p Person) getName() string {
	return p.name
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "John"
	ftEmployee.age = 30
	ftEmployee.id = 1
	ftEmployee.userName = "johndoe"
	fmt.Println(ftEmployee)
	fmt.Println(ftEmployee.getName())
}
