package user

import "fmt"

type UserId struct {
	id string
}

func NewUserId(id string) (*UserId, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	return &UserId{id: id}, nil
}
