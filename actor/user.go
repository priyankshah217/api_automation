package actor

import (
	"api_automation/entity/article"
	"api_automation/helper"
	"fmt"
)

func (u *User) CreateArticle(a *article.Article) (*article.Response, error) {
	articleRequest := article.NewRequest(a)
	response, err := helper.CreateArticle(u.token, articleRequest)
	if err != nil {
		return nil, fmt.Errorf("error in creating article: %v", err)
	}
	return response, nil
}

func (u *User) UpdateArticle(articleToBeUpdatedRequest *article.Response) (*article.Response, error) {
	response, err := helper.UpdateArticle(u.token, articleToBeUpdatedRequest)
	if err != nil {
		return nil, fmt.Errorf("error in updating article: %v", err)
	}
	return response, nil
}

func (u *User) DeleteArticle(slug string) (*article.Response, error) {
	return helper.DeleteArticle(u.token, slug)
}

func (u *User) GetArticle(slug string) (*article.Response, error) {
	return helper.GetArticle(u.token, slug)
}
