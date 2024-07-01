package user

import (
	"database/sql"
	"fmt"
)

// UserService ユーザサービス
type UserService struct {
	db *sql.DB
}

// NewUserService ユーザサービスを生成する
func NewUserService(db *sql.DB) (*UserService, error) {
	return &UserService{db: db}, nil
}

// Exists ユーザが存在するかどうかを確認する
func (s *UserService) Exists(user *User) (bool, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return false, err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	// ここでユーザが存在するかどうかを確認するロジックを実装する
	rows, err := tx.Query("SELECT * FROM users WHERE username = $1", user.UserName())
	if err != nil {
		return false, fmt.Errorf("failed to userService.Exists(): %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, nil
}
