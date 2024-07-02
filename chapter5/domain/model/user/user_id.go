package user

import (
	"fmt"
	"reflect"
)

// UserId ユーザID
type UserId struct {
	id string
}

// NewUserId ユーザIDを生成する
func NewUserId(id string) (*UserId, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	return &UserId{id: id}, nil
}

// Id IDを取得する
func (u *UserId) Id() string {
	return u.id
}

// Equals ユーザIDが同じかどうか
func (u *UserId) Equals(target *UserId) bool {
	return reflect.DeepEqual(u.id, target.id)
}
