package main

import "fmt"

type structPerson struct {
	name string
	age  int64
	sex  string
}

func showStruct() {
	var p1 structPerson
	var p2 structPerson

	p1.name = "Tom"
	p1.age = 17
	p1.sex = "Man"

	p2 = structPerson{"Jack", 22, "SuperMan"}

	addAge(p1)
	fmt.Println("p1", p1)
	addAge1(&p2)
	fmt.Println("p2", p2)
}

func addAge(person structPerson) {
	person.age++
}

func addAge1(personPtr *structPerson) {
	personPtr.age++
}
