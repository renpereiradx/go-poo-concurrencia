package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func newEmployee(id int, name string) *Employee {
	return &Employee{id, name}
}

func newEmployee2(id int, name string) *Employee {
	return &Employee{id: id, name: name}
}

func main() {
	e := Employee{}
	e.id = 1
	e.name = "John"
	fmt.Println(e)
	e2 := Employee{2, "Jane"}
	fmt.Println(e2)
	e3 := Employee{id: 3, name: "Jack"}
	fmt.Println(e3)
	e4 := new(Employee)
	e4.id = 4
	e4.name = "Jill"
	fmt.Println(e4)
	e5 := newEmployee(5, "James")
	fmt.Println(e5)
	e6 := newEmployee2(6, "Jenny")
	fmt.Println(*e6)
}
