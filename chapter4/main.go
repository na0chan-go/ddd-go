package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/na0chan-go/ddd-go/chapter4/domain/model/user"
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
	err = CreateUser(db, "root")
	if err != nil {
		log.Fatal(err)
	}
}

// CreateUser ユーザを作成する
func CreateUser(db *sql.DB, name string) error {
	userName, err := user.NewUserName(name)
	if err != nil {
		return err
	}
	newUser, err := user.NewUser(*userName)
	if err != nil {
		return err
	}

	userService, err := user.NewUserService(db)
	if err != nil {
		return err
	}

	isExists, err := userService.Exists(newUser)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// トランザクションの終了処理
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	// 既にユーザが存在する場合はエラーを返す
	if isExists {
		return fmt.Errorf(("failed to main.CreateUser(): user already exists"))
	}

	// ここでユーザを登録するロジックを実装する
	_, err = tx.Exec("INSERT INTO users (id, username) VALUES ($1, $2)", newUser.UserId(), newUser.UserName())
	if err != nil {
		return fmt.Errorf("failed to main.CreateUser(): %w", err)
	}
	log.Println("Successfully created user")
	return nil
}
