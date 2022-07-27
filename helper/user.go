package helper

import (
	"api_automation/api"
	"api_automation/constant"
	"api_automation/entity/user"
	"api_automation/header"
	"encoding/json"
	"fmt"
)

func CreateUser(userRequest *user.Request) (*user.Response, error) {
	headers := header.GetJSONHeader()
	response, err := api.NewClient().Post(constant.UserEndpoint, headers, userRequest)
	if err != nil {
		return nil, fmt.Errorf("error in creating user: %w", err)
	}
	var userResponse *user.Response
	err = json.Unmarshal(response, &userResponse)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling response")
	}
	return userResponse, nil
}
