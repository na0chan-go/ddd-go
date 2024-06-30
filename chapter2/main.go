package main

import (
	"fmt"

	"github.com/na0chan-go/learn-ddd-go/chapter2/domein/model"
)

func main() {
	fullName := model.NewFullName("John", "Doe")

	fmt.Println(fullName.FirstName()) // John
	fmt.Println(fullName.LastName())  // Doe
}
