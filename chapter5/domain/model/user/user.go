package user

import (
	"reflect"

	"github.com/google/uuid"
)

// User ユーザ
type User struct {
	userId   UserId
	userName UserName
}

// NewUser ユーザを生成する
func NewUser(userName UserName) (*User, error) {
	// ユーザIDはUUIDで生成する
	userId, err := NewUserId(uuid.New().String())
	if err != nil {
		return nil, err
	}
	return &User{userId: *userId, userName: userName}, nil
}

// UserId UserIdを取得する
func (u *User) UserId() string {
	return u.userId.Id()
}

// UserName ユーザ名を取得する
func (u *User) UserName() string {
	return u.userName.Name()
}

// Equals ユーザが同じかどうか
func (u *User) Equals(target *User) bool {
	return reflect.DeepEqual(u.userId, target.userId)
}
