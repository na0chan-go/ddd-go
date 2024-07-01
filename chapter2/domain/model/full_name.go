package model

// fullName struct
type fullName struct {
	firstName Name
	lastName  Name
}

// NewFullName コンストラクタ
func NewFullName(firstName Name, lastName Name) *fullName {
	return &fullName{
		firstName: firstName,
		lastName:  lastName,
	}
}

// Equals フルネームを比較する
func (f *fullName) Equals(other *fullName) bool {
	if other == nil {
		return false
	}
	return f.firstName == other.firstName && f.lastName == other.lastName
}
