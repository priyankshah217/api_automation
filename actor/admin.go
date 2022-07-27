package actor

import (
	"api_automation/entity/user"
	"api_automation/helper"
)

func (a Admin) CreateUser(u *user.User) *User {
	request := user.NewRequest(u)
	response, err := helper.CreateUser(request)
	if err != nil {
		return nil
	}
	return &User{response, response.User.Token}
}
