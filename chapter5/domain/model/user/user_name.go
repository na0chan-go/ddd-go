package user

import (
	"fmt"
	"reflect"
)

// UserName ユーザ名
type UserName struct {
	name string
}

// NewUserName ユーザ名を生成する
func NewUserName(name string) (*UserName, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if len(name) < 3 {
		return nil, fmt.Errorf("name must be at least 3 characters long, got '%s'", name)
	}
	return &UserName{name: name}, nil
}

// Name ユーザ名を取得する
func (u *UserName) Name() string {
	return u.name
}

// Equals ユーザ名が同じかどうか
func (u *UserName) Equals(target *UserName) bool {
	return reflect.DeepEqual(u.name, target.name)
}
