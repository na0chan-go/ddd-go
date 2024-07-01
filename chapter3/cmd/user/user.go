package main

import (
	"fmt"
	"log"

	"github.com/na0chan-go/ddd-go/chapter3/domain/model/user"
)

func main() {
	userId, err := user.NewUserId("1")
	if err != nil {
		log.Fatal(err)
	}

	newUser, err := user.NewUser(*userId, "na0chan")
	if err != nil {
		log.Fatal(err)
	}
	newUserA := *newUser

	// ユーザ名を変更する
	if err := newUser.ChangeUserName("na0chan-go"); err != nil {
		log.Fatal(err)
	}
	newUserB := *newUser
	fmt.Println(newUserA)
	fmt.Println(newUserB)

	// 名前が変更されているが、UserIdが同じなのでtrue
	fmt.Println(newUserA.Equals(&newUserB))

}
