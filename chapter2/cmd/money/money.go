package main

import (
	"log"

	"github.com/na0chan-go/ddd-go/chapter2/domain/model"
)

func main() {
	myMoney := model.NewMoney(1000, "JPY")
	allowance := model.NewMoney(500, "JPY")
	// 金額を加算する
	result, err := myMoney.Add(allowance)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result) // 1500

	jpy := model.NewMoney(1000, "JPY")
	usd := model.NewMoney(100, "USD")
	// 金額を加算する
	result, err = jpy.Add(usd) // 通貨が異なるのでnilが返る
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result) // 0
}
