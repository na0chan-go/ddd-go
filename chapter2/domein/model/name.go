package model

import "regexp"

// Name 姓、名を表す値オブジェクト
type Name string

// NewName コンストラクタ
func NewName(name string) Name {
	if !Name(name).ValidateName(name) {
		return Name("")
	}
	return Name(name)
}

// Value 姓、名を取得する
func (n Name) Value() string {
	return string(n)
}

// IsNullOrEmpty 姓、名が空かどうかを判定する
func (n Name) IsNullOrEmpty() bool {
	return n == ""
}

// ValidateName 姓、名が正しいかどうかを判定する
func (n Name) ValidateName(value string) bool {
	// アルファベットのみにマッチする正規表現
	regex := regexp.MustCompile(`^[a-zA-Z]+$`)
	return regex.MatchString(value)
}
