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

	// ユーザを作成する
	err = CreateUser(db, "1", "na0chan-go", "example@example.com")
	if err != nil {
		log.Fatal(err)
	}
}

// CreateUser ユーザを作成する
func CreateUser(db *sql.DB, id, name, email string) error {
	userName, err := user.NewUserName(name)
	if err != nil {
		return err
	}
	userId, err := user.NewUserId(id)
	if err != nil {
		return err
	}
	mailAddress, err := user.NewMailAddress(email)
	if err != nil {
		return err
	}

	newUser, err := user.NewUser(*userId, *userName, *mailAddress)
	if err != nil {
		return err
	}

	userRepository, err := user.NewUserRepository(db)
	if err != nil {
		return err
	}

	userService, err := user.NewUserService(userRepository)
	if err != nil {
		return err
	}

	isExists, err := userService.Exists(newUser)
	if err != nil {
		return err
	}

	// 既にユーザが存在する場合はエラーを返す
	if isExists {
		return fmt.Errorf(("failed to main.CreateUser(): user already exists"))
	}

	if err := userRepository.Save(newUser); err != nil {
		return err
	}

	log.Println("Successfully created user")
	return nil
}
