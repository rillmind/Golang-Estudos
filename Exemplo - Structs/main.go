package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var persona3 person
	persona1 := person{"Raul", "Lopes"}
	persona2 := person{
		firstName: "Alex",
		lastName:  "Montano",
	}

	persona3.firstName = "Roberta"
	persona3.lastName = "Carrilho"

	persona1.print()
	persona2.print()
	persona3.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
