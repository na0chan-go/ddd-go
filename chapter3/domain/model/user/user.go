package user

import (
	"fmt"
	"reflect"
)

// User ユーザ
type User struct {
	userId UserId
	name   string
}

// NewUser ユーザを生成する
func NewUser(userId UserId, name string) (*User, error) {
	user := new(User)
	user.userId = userId
	if err := user.ChangeUserName(name); err != nil {
		return nil, err
	}
	return user, nil
}

// ChangeUserName ユーザ名を変更する
func (u *User) ChangeUserName(name string) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if len(name) < 3 {
		return fmt.Errorf("name must be at least 3 characters long, got '%s'", name)
	}

	u.name = name
	return nil
}

// Equals ユーザが同じかどうか
func (u *User) Equals(target *User) bool {
	return reflect.DeepEqual(u.userId, target.userId)
}
