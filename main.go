package main

import "fmt"

func main() {

	fixture := NewFixture()
	person := NewPerson()

	fixture.Create(person)

	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.BloodDonator)
}
