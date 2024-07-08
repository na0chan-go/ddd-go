package user

import (
	"fmt"

	"github.com/google/uuid"
)

// UserRegisterServiceInterface ユーザ登録サービスインターフェース
type UserRegisterServiceInterface interface {
	Handle(command UserRegisterCommand) error
}

// UserRegisterService ユーザ登録サービス
type UserRegisterService struct {
	userRepository UserRepositoryInterface
	userService    UserService
}

// NewUserRegisterService ユーザ登録サービスを生成する
func NewUserRegisterService(_userRepository UserRepositoryInterface, _userService UserService) (*UserRegisterService, error) {
	return &UserRegisterService{
		userRepository: _userRepository,
		userService:    _userService,
	}, nil
}

// Handle ユーザを登録する
func (s *UserRegisterService) Handle(command UserRegisterCommand) error {
	userId, err := NewUserId(uuid.New().String())
	if err != nil {
		return err
	}

	userName, err := NewUserName(command.Name)
	if err != nil {
		return err
	}

	mailAddress, err := NewMailAddress(command.MailAddress)
	if err != nil {
		return err
	}

	user, err := NewUser(
		*userId, *userName, *mailAddress,
	)
	if err != nil {
		return err
	}

	isExists, err := s.userService.Exists(user)
	if err != nil {
		return err
	}
	if isExists {
		return fmt.Errorf("user already exists")
	}
	err = s.userRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}
