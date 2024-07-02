package user

import (
	"fmt"
	"reflect"
)

// UserName ユーザ名
type UserName struct {
	value string
}

// NewUserName ユーザ名を生成する
func NewUserName(value string) (*UserName, error) {
	if len(value) < 3 {
		return nil, fmt.Errorf("name must be at least 3 characters long, got '%s'", value)
	}
	return &UserName{value: value}, nil
}

// Equals ユーザ名が同じかどうか
func (u *UserName) Equals(target *UserName) bool {
	return reflect.DeepEqual(u.value, target.value)
}
