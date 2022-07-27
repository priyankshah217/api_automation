package user

func NewRequest(u *User) *Request {
	return &Request{*u}
}
