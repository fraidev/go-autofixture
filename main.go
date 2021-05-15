package main

import "fmt"

func main() {

	fixture := NewFixture()
	person := NewPerson()


	fixture.Create(person)


	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.BloodDonator)
	fmt.Println(person.B)
	fmt.Println(person.priva)
	// fmt.Println(person.Job)
	fmt.Println(person.Job.Level)
	fmt.Println(person.Job.Name)
	fmt.Println(person.JobWithoutPointer.Level)
	fmt.Println(person.JobWithoutPointer.Name)
	// fmt.Println(person.Dependents)
}
