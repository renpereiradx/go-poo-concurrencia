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
	endData string
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

type PrintInfo interface {
	getInfo() string
}

func getInfo(p PrintInfo) {
	fmt.Println(p.getInfo())
}

func (ftEmployee FullTimeEmployee) getInfo() string {
	return fmt.Sprintf("Full time %s", ftEmployee.name)
}

func (tEmployee TemporaryEmployee) getInfo() string {
	return fmt.Sprintf("Temporary %s", tEmployee.name)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "John"
	ftEmployee.age = 30
	ftEmployee.id = 1
	ftEmployee.endData = "2020-01-01"
	getInfo(ftEmployee)
	tEmployee := TemporaryEmployee{}
	tEmployee.name = "Jane"
	tEmployee.age = 25
	tEmployee.id = 2
	tEmployee.taxRate = 10
	getInfo(tEmployee)
}
