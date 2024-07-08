package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/na0chan-go/ddd-go/chapter6/domain/model/user"
)

// main メイン関数
func main() {
	postgres := fmt.Sprintf("postgres://%s/%s?sslmode=disable&user=%s&password=%s&port=%s&timezone=Asia/Tokyo",
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"))
	db, err := sql.Open("postgres", postgres)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to database")

	userRepository, err := user.NewUserRepository(db)
	if err != nil {
		log.Fatal(err)
	}
	userService, err := user.NewUserService(userRepository)
	if err != nil {
		log.Fatal(err)
	}
	userRegisterService, err := user.NewUserRegisterService(userRepository, *userService)
	if err != nil {
		log.Fatal(err)
	}

	// ユーザを作成する
	err = userRegisterService.Handle(user.UserRegisterCommand{
		Name:        "na0chan-go",
		MailAddress: "example@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}
}
