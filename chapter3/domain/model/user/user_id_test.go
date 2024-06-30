package user

import (
	"testing"
)

// NewUserID関数のテスト関数
func TestNewUserID(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"Valid ID", "user-123", false},
		{"Empty ID", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewUserId(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
