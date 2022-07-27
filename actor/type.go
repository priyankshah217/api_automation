package actor

import (
	"api_automation/entity/article"
	"api_automation/entity/user"
)

type IAdmin interface {
	CreateUser(*user.User) (*User, error)
}

type Admin struct {
}

func NewAdmin() IAdmin {
	return &Admin{}
}

type User struct {
	*user.Response
	token string
}

type IUser interface {
	CreateArticle(*article.Article) (*article.Response, error)
	UpdateArticle(*article.Response) (*article.Response, error)
	DeleteArticle(string) (*article.Response, error)
	GetArticle(string) (*article.Response, error)
}
