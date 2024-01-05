package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	ID       int
	Position string
}

type FullTimeEmployee struct {
	Person
	Employee
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	return Person{}, nil
}

var GetEmployeeByID = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeByID(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee
	person, errP := GetPersonByDNI(dni)
	if errP != nil {
		return ftEmployee, errP
	}
	ftEmployee.Person = person
	employee, errE := GetEmployeeByID(id)
	if errE != nil {
		return ftEmployee, errE
	}
	ftEmployee.Employee = employee
	return ftEmployee, nil
}
