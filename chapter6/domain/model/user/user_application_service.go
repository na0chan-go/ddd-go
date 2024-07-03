package user

import "fmt"

// UserApplicationService ユーザアプリケーションサービス
type UserApplicationService struct {
	userRepository UserRepositoryInterface
	userService    UserService
}

// NewUserApplicationService ユーザアプリケーションサービスを生成する
func NewUserApplicationService(_userRepository UserRepositoryInterface, _userService UserService) (*UserApplicationService, error) {
	return &UserApplicationService{
		userRepository: _userRepository,
		userService:    _userService,
	}, nil
}

// Register ユーザを登録する
func (s *UserApplicationService) Register(name string) error {

	userName, err := NewUserName(name)
	if err != nil {
		return err
	}
	userId, err := NewUserId("123")
	if err != nil {
		return err
	}
	mailAddress, err := NewMailAddress("example@example.com")
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

// Get ユーザを取得する
func (s *UserApplicationService) Get(userId string) (*UserData, error) {
	targetId, err := NewUserId(userId)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepository.FindByUserId(targetId)
	if err != nil {
		return nil, err
	}

	userData := NewUserData(*user)

	return userData, nil
}

// Update ユーザを更新する
func (s *UserApplicationService) Update(command UserUpdateCommand) error {
	targetId, err := NewUserId(command.UserId)
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

	name := command.UserName
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
		mailAddress, err := NewMailAddress(mailAddress)
		if err != nil {
			return err
		}
		user.ChangeMailAddress(*mailAddress)
	}

	err = s.userRepository.Save(user)
	if err != nil {
		return err
	}

	return nil
}
