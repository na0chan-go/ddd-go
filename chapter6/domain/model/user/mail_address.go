package user

// MailAddress メールアドレス
type MailAddress struct {
	value string
}

// NewMailAddress メールアドレスを生成する
func NewMailAddress(value string) (*MailAddress, error) {
	return &MailAddress{value: value}, nil
}

// Equals メールアドレスが同じかどうか
func (m *MailAddress) Equals(target *MailAddress) bool {
	return m.value == target.value
}
