package user

import (
	"testing"
)

// NewUser関数のテスト関数
func TestNewUser(t *testing.T) {
	userId, err := NewUserId("1")
	if err != nil {
		t.Fatal(err)
	}

	newUser, err := NewUser(*userId, "na0chan")
	if err != nil {
		t.Fatal(err)
	}

	if newUser.userId != *userId {
		t.Errorf("NewUser() userId = %v, want %v", newUser.userId, userId)
	}
	if newUser.name != "na0chan" {
		t.Errorf("NewUser() name = %v, want na0chan", newUser.name)
	}
}

// ChangeUserNameメソッドのテスト関数
func TestChangeUserName(t *testing.T) {
	tests := []struct {
		name    string
		newName string
		wantErr bool
	}{
		{"Valid name", "John Doe", false},
		{"Empty name", "", true},
		{"Short name", "Jo", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{}
			err := user.ChangeUserName(tt.newName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeUserName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Equalsメソッドのテスト関数
func TestEquals(t *testing.T) {
	userId, err := NewUserId("1")
	if err != nil {
		t.Fatal(err)
	}

	newUserA, err := NewUser(*userId, "na0chan")
	if err != nil {
		t.Fatal(err)
	}

	newUserB, err := NewUser(*userId, "na0chan-go")
	if err != nil {
		t.Fatal(err)
	}

	if !newUserA.Equals(newUserB) {
		t.Errorf("Equals() = %v, want %v", newUserA.Equals(newUserB), true)
	}
}
