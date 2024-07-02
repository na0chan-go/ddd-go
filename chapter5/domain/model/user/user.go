package user

// User ユーザ
type User struct {
	id   UserId
	name UserName
}

// NewUser ユーザを生成する
func NewUser(userId UserId, userName UserName) (*User, error) {
	return &User{id: userId, name: userName}, nil
}

// Id Idを取得する
func (u *User) Id() *UserId {
	return &u.id
}

// Name ユーザ名を取得する
func (u *User) Name() *UserName {
	return &u.name
}
