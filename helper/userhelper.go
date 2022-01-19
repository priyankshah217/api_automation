package helper

import (
	"api_automation/constant"
	"api_automation/entity/user"
	"api_automation/http"
	"api_automation/testcontext"
	"context"
	"encoding/json"
)

func CreateUser(ctx context.Context, userRequest *user.Request) context.Context {
	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	response := http.PostRequest(constant.USER_ENDPOINT, headers, userRequest)
	var userResponse user.Response
	err := json.Unmarshal(response, &userResponse)
	if err != nil {
		return nil
	}
	return context.WithValue(ctx, testcontext.UserResponse{}, userResponse)
}
