package main

import (
	"fmt"

	"github.com/na0chan-go/learn-ddd-go/chapter2/domein/model"
)

func main() {
	firstName := model.NewName("John")
	if firstName.IsNullOrEmpty() {
		fmt.Println("firstNameが空、または不正です。")
	}

	lastName := model.NewName("Doe")
	if lastName.IsNullOrEmpty() {
		fmt.Println("lastNameが空、または不正です。")
	}

	fullName := model.NewFullName(firstName, lastName)

	fmt.Println(fullName.FirstName()) // John
	fmt.Println(fullName.LastName())  // Doe

	myMoney := model.NewMoney(1000, "JPY")
	allowance := model.NewMoney(500, "JPY")
	// 金額を加算する
	result := myMoney.Add(allowance)
	if result == nil {
		fmt.Println("通貨が異なります。")
	}
	fmt.Println(result.Amount()) // 1500

	jpy := model.NewMoney(1000, "JPY")
	usd := model.NewMoney(100, "USD")
	// 金額を加算する
	result = jpy.Add(usd) // 通貨が異なるのでnilが返る
	if result == nil {
		fmt.Println("通貨が異なります。")
	}

}
