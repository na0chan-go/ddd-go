package user

type Client struct {
	userRegisterService UserRegisterServiceInterface
}

func (c *Client) Register(name, mailAddress string) {
	command, _ := NewUserRegisterCommand(name, mailAddress)
	c.userRegisterService.Handle(*command)
}
