package main

import (
	"fmt"
	"os"

	"github.com/na0chan-go/learn-ddd-go/chapter2/domein/model"
)

func main() {
	firstName := model.NewName("John")
	if firstName.IsNullOrEmpty() {
		fmt.Println("firstNameが空、または不正です。")
		os.Exit(1)
	}

	lastName := model.NewName("Doe")
	if lastName.IsNullOrEmpty() {
		fmt.Println("lastNameが空、または不正です。")
		os.Exit(1)
	}

	fullName := model.NewFullName(firstName, lastName)

	fmt.Println(fullName.FirstName()) // John
	fmt.Println(fullName.LastName())  // Doe
}
