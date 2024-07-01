package user

import (
	"testing"
)

func TestNewUserName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		{"Valid Name", "JohnDoe", false},
		{"Empty Name", "", true},
		{"Short Name", "Jo", true},
	}

	for _, tt := range tests {
		got, err := NewUserName(tt.input)
		if (err != nil) != tt.wantError {
			t.Errorf("%s: NewUserName(%s) error = %v, wantError %v", tt.name, tt.input, err, tt.wantError)
		}
		if !tt.wantError && got.Name() != tt.input {
			t.Errorf("%s: NewUserName(%s) = %v, want %v", tt.name, tt.input, got.Name(), tt.input)
		}
	}
}

func TestUserName_Name(t *testing.T) {
	name := "JohnDoe"
	userName, err := NewUserName(name)
	if err != nil {
		t.Fatalf("NewUserName(%s) failed: %v", name, err)
	}

	if got := userName.Name(); got != name {
		t.Errorf("UserName.Name() = %v, want %v", got, name)
	}
}

func TestUserName_Equals(t *testing.T) {
	userName1, _ := NewUserName("JohnDoe")
	userName2, _ := NewUserName("JohnDoe")
	userName3, _ := NewUserName("JaneDoe")

	if !userName1.Equals(userName2) {
		t.Errorf("Expected userName1 to equal userName2")
	}

	if userName1.Equals(userName3) {
		t.Errorf("Expected userName1 not to equal userName3")
	}
}
