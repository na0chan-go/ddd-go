package user

import (
	"reflect"
)

// UserId ユーザID
type UserId struct {
	value string
}

// NewUserId ユーザIDを生成する
func NewUserId(value string) (*UserId, error) {
	return &UserId{value: value}, nil
}

// Equals ユーザIDが同じかどうか
func (u *UserId) Equals(target *UserId) bool {
	return reflect.DeepEqual(u.value, target.value)
}
