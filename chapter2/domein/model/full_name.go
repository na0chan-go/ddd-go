package model

// fullName struct
type fullName struct {
	firstName string
	lastName  string
}

// NewFullName constructor
func NewFullName(firstName, lastName string) *fullName {
	return &fullName{
		firstName: firstName,
		lastName:  lastName,
	}
}

// FirstName getter
func (f *fullName) FirstName() string {
	return f.firstName
}

// LastName getter
func (f *fullName) LastName() string {
	return f.lastName
}

// Equals フルネームを比較する
func (f *fullName) Equals(other *fullName) bool {
	if other == nil {
		return false
	}
	return f.firstName == other.firstName && f.lastName == other.lastName
}
