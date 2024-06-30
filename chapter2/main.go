package main

import (
	"fmt"
)

func main() {
	fullName := newFullName("John", "Doe")
	fmt.Println(fullName.firstName) // John
	fmt.Println(fullName.lastName)  // Doe

	fmt.Println(fullName.FirstName()) // John
	fmt.Println(fullName.LastName())  // Doe
}

// fullName struct
type fullName struct {
	firstName string
	lastName  string
}

// newFullName constructor
func newFullName(firstName, lastName string) *fullName {
	return &fullName{
		firstName: firstName,
		lastName:  lastName,
	}
}

// FirstName getter
func (f *fullName) FirstName() string {
	return f.firstName
}

// LastName getter
func (f *fullName) LastName() string {
	return f.lastName
}
