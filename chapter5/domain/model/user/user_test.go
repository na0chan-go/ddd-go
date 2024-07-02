package user

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewUser(t *testing.T) {
	userName, _ := NewUserName("JohnDoe")
	user, err := NewUser(*userName)
	if err != nil {
		t.Fatalf("NewUser() error = %v", err)
	}
	if user.userName != *userName {
		t.Errorf("NewUser().userName = %v, want %v", user.userName, userName)
	}
	if _, err := uuid.Parse(user.userId.Id()); err != nil {
		t.Errorf("NewUser().userId should be a valid UUID, got %v", user.userId.Id())
	}
}

func TestUser_UserId(t *testing.T) {
	userName, _ := NewUserName("JohnDoe")
	user, _ := NewUser(*userName)
	if got := user.UserId(); reflect.TypeOf(got).Kind() != reflect.String {
		t.Errorf("User.UserId() = %v, want a string", got)
	}
}

func TestUser_UserName(t *testing.T) {
	userName, _ := NewUserName("JohnDoe")
	user, _ := NewUser(*userName)
	if got := user.UserName(); got != "JohnDoe" {
		t.Errorf("User.UserName() = %v, want %v", got, "JohnDoe")
	}
}

func TestUser_Equals(t *testing.T) {
	userName1, _ := NewUserName("JohnDoe")
	user1, _ := NewUser(*userName1)
	user2, _ := NewUser(*userName1)
	userName3, _ := NewUserName("JaneDoe")
	user3, _ := NewUser(*userName3)

	if !user1.Equals(user1) {
		t.Errorf("Expected user1 to equal itself")
	}

	if user1.Equals(user2) {
		t.Errorf("Expected user1 not to equal user2 despite the same userName because userId should be different")
	}

	if user1.Equals(user3) {
		t.Errorf("Expected user1 not to equal user3")
	}
}
