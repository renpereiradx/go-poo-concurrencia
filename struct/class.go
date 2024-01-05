package main

import "fmt"

type Person struct {
	id   int16
	name string
}

// Receivers function / values and pointers reiceivers
func (p Person) String() string {
	return fmt.Sprintf("Person %d: %s", p.id, p.name)
}

func (p *Person) setID(id int16) {
	p.id = id
}

func (p *Person) setName(name string) {
	p.name = name
}

func (p *Person) getID() int16 {
	return p.id
}

func (p *Person) getName() string {
	return p.name
}

func main() {
	p := new(Person)
	p2 := Person{}
	p2.setID(2)
	p2.setName("Jane")
	fmt.Println(p)
	p.setID(1)
	p.setName("John")
	fmt.Println(p)
	fmt.Println(p.getID())
	fmt.Println(p.getName())
	fmt.Println(p2)
}
