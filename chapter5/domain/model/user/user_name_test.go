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
		if !tt.wantError && got.value != tt.input {
			t.Errorf("%s: NewUserName(%s) = %v, want %v", tt.name, tt.input, got.value, tt.input)
		}
	}
}
