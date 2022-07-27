package helper

import (
	"api_automation/constant"
	"api_automation/entity/user"
	"api_automation/http"
	"encoding/json"
	"fmt"
)

func CreateUser(userRequest *user.Request) (*user.Response, error) {
	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	response, err := http.PostRequest(constant.USER_ENDPOINT, headers, userRequest)
	if err != nil {
		return nil, fmt.Errorf("error in creating user: %v", err)
	}
	var userResponse *user.Response
	err = json.Unmarshal(response, &userResponse)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling response")
	}
	return userResponse, nil
}
