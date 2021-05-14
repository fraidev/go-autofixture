package main

type Person struct {
	Name string
	Age int
	BloodDonator bool

}

func NewPerson() *Person {
	return &Person{};
}