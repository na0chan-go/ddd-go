package user

// UserDeleteService ユーザアプリケーションサービス
type UserDeleteService struct {
	userRepository UserRepositoryInterface
}

// NewUserDeleteService ユーザアプリケーションサービスを生成する
func NewUserDeleteService(_userRepository UserRepositoryInterface) (*UserDeleteService, error) {
	return &UserDeleteService{
		userRepository: _userRepository,
	}, nil
}

// Handle ユーザを削除する
func (s *UserDeleteService) Handle(command UserDeleteCommand) error {
	targetId, err := NewUserId(command.Id)
	if err != nil {
		return err
	}
	user, err := s.userRepository.FindByUserId(targetId)
	if err != nil {
		return err
	}

	// ユーザが存在しない場合は退会成功とする
	if user == nil {
		return nil
	}

	err = s.userRepository.Delete(user)
	if err != nil {
		return err
	}

	return nil

}
