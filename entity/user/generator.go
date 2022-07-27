package user

type IBuilder interface {
	SetEmail(email string) IBuilder
	SetPassword(password string) IBuilder
	SetUsername(username string) IBuilder
	Build() *User
}

type Builder struct {
	u User
}

func NewBuilder() IBuilder {
	return &Builder{}
}

func (b *Builder) SetEmail(email string) IBuilder {
	b.u.Email = email
	return b
}

func (b *Builder) SetPassword(password string) IBuilder {
	b.u.Password = password
	return b
}

func (b *Builder) SetUsername(username string) IBuilder {
	b.u.Username = username
	return b
}

func (b *Builder) Build() *User {
	return &b.u
}
