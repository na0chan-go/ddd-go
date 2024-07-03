package user

// UserData ユーザデータ
type UserData struct {
	id          string
	name        string
	mailAddress string
}

// NewUserData ユーザデータを生成する
func NewUserData(source User) *UserData {
	return &UserData{
		id:          source.id.value,
		name:        source.name.value,
		mailAddress: source.mailAddress.value,
	}
}

// Id IDを取得する
func (u *UserData) Id() string {
	return u.id
}

// Name ユーザ名を取得する
func (u *UserData) Name() string {
	return u.name
}

// MailAddress メールアドレスを取得する
func (u *UserData) MailAddress() string {
	return u.mailAddress
}
