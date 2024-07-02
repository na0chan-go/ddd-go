package user

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	userId, err := NewUserId("123")
	if err != nil {
		t.Fatalf("Failed to create UserId: %v", err)
	}
	userName, err := NewUserName("JohnDoe")
	if err != nil {
		t.Fatalf("Failed to create UserName: %v", err)
	}

	user, err := NewUser(*userId, *userName)
	if err != nil {
		t.Fatalf("NewUser() error = %v", err)
	}

	if *user.Id() != *userId {
		t.Errorf("user.Id() = %v, want %v", *user.Id(), *userId)
	}

	if *user.Name() != *userName {
		t.Errorf("user.Name() = %v, want %v", *user.Name(), *userName)
	}
}
