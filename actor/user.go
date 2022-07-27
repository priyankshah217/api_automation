package actor

import (
	"api_automation/entity/article"
	"api_automation/helper"
)

func (u *User) CreateArticle(a *article.Article) (*article.Response, error) {
	articleRequest := article.NewRequest(a)
	return helper.CreateArticle(u.token, articleRequest)
}

func (u *User) UpdateArticle(articleToBeUpdatedRequest *article.Response) (*article.Response, error) {
	return helper.UpdateArticle(u.token, articleToBeUpdatedRequest)
}

func (u *User) DeleteArticle(slug string) (*article.Response, error) {
	return helper.DeleteArticle(u.token, slug)
}

func (u *User) GetArticle(slug string) (*article.Response, error) {
	return helper.GetArticle(u.token, slug)
}
