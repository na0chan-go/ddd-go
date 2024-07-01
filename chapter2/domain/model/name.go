package model

import (
	"errors"
	"regexp"
)

// Name 姓、名を表す値オブジェクト
type Name string

// NewName コンストラクタ
func NewName(name string) (Name, error) {
	// 姓、名が空文字かどうかを判定する
	if name == "" {
		return Name(""), errors.New("名前が空です。")
	}

	// 名前が正しいかどうかを判定する
	if !Name(name).ValidateName(name) {
		return Name(""), errors.New("名前が不正です。")
	}
	return Name(name), nil
}

// ValidateName 姓、名が正しいかどうかを判定する
func (n Name) ValidateName(value string) bool {
	// アルファベットのみにマッチする正規表現
	regex := regexp.MustCompile(`^[a-zA-Z]+$`)
	return regex.MatchString(value)
}
