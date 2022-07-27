package user

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Request struct {
	User User `json:"user"`
}

type Response struct {
	User struct {
		Email    string      `json:"email"`
		Username string      `json:"username"`
		Bio      interface{} `json:"bio"`
		Image    string      `json:"image"`
		Token    string      `json:"token"`
	} `json:"user"`
}
