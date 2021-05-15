package main

type Person struct {
	Name string
	Age int
	BloodDonator bool
	Job *Job //pointer struct
	JobWithoutPointer Job //pointer struct
	// Dependents []Person
	B uint8
	priva uint8
}

type Job struct {
	Name string
	Level int
}

func NewPerson() *Person {
	return &Person{Job: &Job{}, JobWithoutPointer: Job{}};
}