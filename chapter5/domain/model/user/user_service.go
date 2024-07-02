package user

// UserService ユーザサービス
type UserService struct {
	userRepository UserRepositoryInterface
}

// NewUserService ユーザサービスを生成する
func NewUserService(userRepository UserRepositoryInterface) (*UserService, error) {
	return &UserService{userRepository: userRepository}, nil
}

// Exists ユーザが存在するかどうかを確認する
func (s *UserService) Exists(user *User) (bool, error) {
	user, err := s.userRepository.FindByUserName(user.Name())
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
