package user

type UserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Request struct {
	User UserReq `json:"user"`
}

type UserRes struct {
	Email    string      `json:"email"`
	Username string      `json:"username"`
	Bio      interface{} `json:"bio"`
	Image    string      `json:"image"`
	Token    string      `json:"token"`
}

type Response struct {
	User UserRes `json:"user"`
}

func NewRequest(email string, password string, userName string) *Request {
	return &Request{User: UserReq{Email: email, Password: password, Username: userName}}
}
