package main

import (
	"fmt"
	"strings"
)

func main() {
	fullName := "John Doe"
	token := strings.Split(fullName, " ")
	firstName := token[0]
	fmt.Println(firstName) // John
}
