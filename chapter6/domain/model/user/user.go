package user

// User ユーザ
type User struct {
	id          UserId
	name        UserName
	mailAddress MailAddress
}

// NewUser ユーザを生成する
func NewUser(_userId UserId, _userName UserName, _mailAddress MailAddress) (*User, error) {
	return &User{id: _userId, name: _userName, mailAddress: _mailAddress}, nil
}

// Id Idを取得する
func (u *User) Id() *UserId {
	return &u.id
}

// Name ユーザ名を取得する
func (u *User) Name() *UserName {
	return &u.name
}

// MailAddress メールアドレスを取得する
func (u *User) MailAddress() *MailAddress {
	return &u.mailAddress
}

// ChangeName ユーザ名を変更する
func (u *User) ChangeName(name UserName) {
	u.name = name
}

// ChangeMailAddress メールアドレスを変更する
func (u *User) ChangeMailAddress(mailAddress MailAddress) {
	u.mailAddress = mailAddress
}
