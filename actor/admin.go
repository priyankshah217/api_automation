package actor

import (
	"api_automation/entity/user"
	"api_automation/helper"
)

func (a *Admin) CreateUser(u *user.User) (*User, error) {
	request := user.NewRequest(u)
	response, err := helper.CreateUser(request)
	if err != nil {
		return nil, err
	}
	return &User{response, response.User.Token}, nil
}
