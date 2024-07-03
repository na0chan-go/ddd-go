package user

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// UserRepositoryInterface ユーザリポジトリインターフェース
type UserRepositoryInterface interface {
	FindByUserName(name *UserName) (*User, error)
	FindByUserId(id *UserId) (*User, error)
	Save(user *User) error
}

// UserRepository ユーザリポジトリ
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository ユーザリポジトリを生成する
func NewUserRepository(db *sql.DB) (*UserRepository, error) {
	return &UserRepository{db: db}, nil
}

// FindByUserId ユーザIDからユーザを取得する
func (r *UserRepository) Save(user *User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	_, err = tx.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", user.id.value, user.name.value)
	if err != nil {
		return err
	}
	return nil
}

// FindByUserId ユーザIDからユーザを取得する
func (r *UserRepository) FindByUserName(name *UserName) (*User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	rows, err := tx.Query("SELECT * FROM users WHERE name = $1", name.value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userId := &UserId{}
	userName := &UserName{}
	mailAddress := &MailAddress{}
	var user *User
	// ユーザが存在する場合はユーザを返す
	if rows.Next() {
		err := rows.Scan(&userId.value, &userName.value, &mailAddress.value)
		if err != nil {
			return nil, err
		}
		user = &User{id: *userId, name: *userName, mailAddress: *mailAddress}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByUserId ユーザIDからユーザを取得する
func (r *UserRepository) FindByUserId(id *UserId) (*User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	rows, err := tx.Query("SELECT * FROM users WHERE id = $1", id.value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userId := &UserId{}
	userName := &UserName{}
	mailAddress := &MailAddress{}
	var user *User
	// ユーザが存在する場合はユーザを返す
	if rows.Next() {
		err := rows.Scan(&userId.value, &userName.value, &mailAddress.value)
		if err != nil {
			return nil, err
		}
		user = &User{id: *userId, name: *userName, mailAddress: *mailAddress}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}
