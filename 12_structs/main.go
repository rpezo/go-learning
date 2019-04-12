package main

import (
	"fmt"
	"strconv"
)

//Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Rodrigo", lastName: "Banner", city: "Santiago", gender: "m", age: 39}
	person2 := Person{firstName: "Bruce", lastName: "Wayne", city: "Gotham", gender: "m", age: 43}
	person3 := Person{firstName: "Scarlet", lastName: "Witch", city: "Gotham", gender: "f", age: 36}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person2.age++
	fmt.Println(person2)
	fmt.Println(person1.greet())
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried(person3.lastName)
	fmt.Println(person1.greet())

}
