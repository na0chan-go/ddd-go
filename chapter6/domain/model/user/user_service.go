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
	// 重複のルールをユーザ名からメールアドレスに変更
	// user, err := s.userRepository.FindByUserName(user.Name())
	duplicateUser, err := s.userRepository.FindByMailAddress(user.MailAddress())
	if err != nil {
		return false, err
	}
	return duplicateUser != nil, nil
}
