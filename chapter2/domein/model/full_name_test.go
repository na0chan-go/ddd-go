package model

import "testing"

func TestNewFullName(t *testing.T) {
	firstName := Name("Taro")
	lastName := Name("Yamada")
	fullName := NewFullName(firstName, lastName)

	if fullName.firstName != firstName || fullName.lastName != lastName {
		t.Errorf("NewFullName() failed, expected %v %v, got %v %v", firstName, lastName, fullName.firstName, fullName.lastName)
	}
}

func TestFullName_Equals(t *testing.T) {
	name1 := NewFullName(Name("Taro"), Name("Yamada"))
	name2 := NewFullName(Name("Taro"), Name("Yamada"))
	name3 := NewFullName(Name("Jiro"), Name("Yamada"))
	var name4 *fullName = nil

	tests := []struct {
		name     string
		a        *fullName
		b        *fullName
		expected bool
	}{
		{"SameNames", name1, name2, true},
		{"DifferentNames", name1, name3, false},
		{"CompareWithNil", name1, name4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := tt.a.Equals(tt.b); result != tt.expected {
				t.Errorf("Equals() = %v, want %v", result, tt.expected)
			}
		})
	}
}
