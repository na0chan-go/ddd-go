package user

import (
	"fmt"
)

// UserUpdateService ユーザアプリケーションサービス
type UserUpdateService struct {
	userRepository UserRepositoryInterface
	userService    UserService
}

// NewUserUpdateService ユーザアプリケーションサービスを生成する
func NewUserUpdateService(_userRepository UserRepositoryInterface, _userService UserService) (*UserUpdateService, error) {
	return &UserUpdateService{
		userRepository: _userRepository,
		userService:    _userService,
	}, nil
}

// Handle ユーザを更新する
func (s *UserUpdateService) Handle(command UserUpdateCommand) error {
	targetId, err := NewUserId(command.Id)
	if err != nil {
		return err
	}
	user, err := s.userRepository.FindByUserId(targetId)
	if err != nil {
		return err
	}

	if user == nil {
		return fmt.Errorf("user not found")
	}

	name := command.Name
	if name != "" {
		userName, err := NewUserName(name)
		if err != nil {
			return err
		}
		user.ChangeName(*userName)

		isExists, err := s.userService.Exists(user)
		if err != nil {
			return err
		}
		if isExists {
			return fmt.Errorf("user already exists")
		}
	}
	mailAddress := command.MailAddress
	if mailAddress != "" {
		newMailAddress, err := NewMailAddress(mailAddress)
		if err != nil {
			return err
		}
		user.ChangeMailAddress(*newMailAddress)
	}

	err = s.userRepository.Save(user)
	if err != nil {
		return err
	}

	return nil
}
