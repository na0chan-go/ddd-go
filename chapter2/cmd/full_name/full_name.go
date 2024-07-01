package main

import (
	"log"

	"github.com/na0chan-go/ddd-go/chapter2/domain/model"
)

func main() {
	firstNameA, err := model.NewName("John")
	if err != nil {
		log.Fatal(err)
	}

	lastNameA, err := model.NewName("Doe")
	if err != nil {
		log.Fatal(err)
	}

	fullNameA := model.NewFullName(firstNameA, lastNameA)

	firstNameB, err := model.NewName("John")
	if err != nil {
		log.Fatal(err)
	}

	lastNameB, err := model.NewName("Renon")
	if err != nil {
		log.Fatal(err)
	}

	fullNameB := model.NewFullName(firstNameB, lastNameB)

	if fullNameA.Equals(fullNameB) {
		log.Println("同じ名前です")
	} else {
		log.Println("違う名前です")
	}

}
