package user

type UserRegisterCommand struct {
	Name        string
	MailAddress string
}

func NewUserRegisterCommand(name string, mailAddress string) (*UserRegisterCommand, error) {
	return &UserRegisterCommand{
		Name:        name,
		MailAddress: mailAddress,
	}, nil
}
