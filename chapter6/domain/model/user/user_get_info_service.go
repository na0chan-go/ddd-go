package user

// UserGetInfoService ユーザ情報取得サービス
type UserGetInfoService struct {
	userRepository UserRepositoryInterface
}

// NewUserGetInfoService ユーザ情報取得サービスを生成する
func NewUserGetInfoService(_userRepository UserRepositoryInterface, _userService UserService) (*UserGetInfoService, error) {
	return &UserGetInfoService{
		userRepository: _userRepository,
	}, nil
}

// Handle ユーザを取得する
func (s *UserGetInfoService) Handle(command UserGetInfoCommand) (*UserData, error) {
	targetId, err := NewUserId(command.Id)
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
