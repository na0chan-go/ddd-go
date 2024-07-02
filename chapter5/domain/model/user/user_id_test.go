package user

import (
	"testing"
)

func TestNewUserId(t *testing.T) {
	// 正常なIDを渡した場合
	userId, err := NewUserId("123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if userId.Id() != "123" {
		t.Errorf("Expected ID to be '123', got %v", userId.Id())
	}

	// 空のIDを渡した場合
	_, err = NewUserId("")
	if err == nil {
		t.Error("Expected error for empty ID, got none")
	}
}

func TestUserId_Equals(t *testing.T) {
	userId1, _ := NewUserId("123")
	userId2, _ := NewUserId("123")
	userId3, _ := NewUserId("456")

	if !userId1.Equals(userId2) {
		t.Error("Expected userId1 to equal userId2")
	}

	if userId1.Equals(userId3) {
		t.Error("Expected userId1 not to equal userId3")
	}
}
