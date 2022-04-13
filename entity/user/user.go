package user

type Req struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Request struct {
	User Req `json:"user"`
}

type Res struct {
	Email    string      `json:"email"`
	Username string      `json:"username"`
	Bio      interface{} `json:"bio"`
	Image    string      `json:"image"`
	Token    string      `json:"token"`
}

type Response struct {
	User Res `json:"user"`
}

func NewRequest(email string, password string, userName string) *Request {
	return &Request{User: Req{Email: email, Password: password, Username: userName}}
}
