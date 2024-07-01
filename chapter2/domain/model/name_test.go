package model

import (
	"testing"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		name      string
		wantValid bool
	}{
		{"John", true},
		{"Sara", true},
		{"123", false},
		{"John Doe", false},
		{"", false},
	}

	for _, tt := range tests {
		_, err := NewName(tt.name)
		if tt.wantValid && err != nil {
			t.Errorf("NewName() = %v, want %v for name %v", err, nil, tt.name)
		}
		if !tt.wantValid && err == nil {
			t.Errorf("NewName() = %v, want %v for name %v", err, nil, tt.name)
		}

	}
}

func TestValidateName(t *testing.T) {
	tests := []struct {
		name      string
		wantValid bool
	}{
		{"John", true},
		{"Sara", true},
		{"123", false},
		{"John Doe", false},
		{"", false},
	}

	for _, tt := range tests {
		n := Name(tt.name)
		if gotValid := n.ValidateName(tt.name); gotValid != tt.wantValid {
			t.Errorf("Name.ValidateName() = %v, want %v for name %v", gotValid, tt.wantValid, tt.name)
		}
	}
}
