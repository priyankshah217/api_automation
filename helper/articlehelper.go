package helper

import (
	"api_automation/api"
	"api_automation/constant"
	"api_automation/entity/article"
	"api_automation/header"
	"encoding/json"
	"fmt"
)

func CreateArticle(token string, articleRequest *article.Request) (*article.Response, error) {
	headers := header.GetJsonHeader()
	headers["authorization"] = "Token " + token
	response, err := api.NewClient().Post(constant.ARTICLE_ENDPOINT, headers, articleRequest)
	if err != nil {
		return nil, err
	}
	var articleResponse *article.Response
	err = json.Unmarshal(response, &articleResponse)
	if err != nil {
		return nil, err
	}
	return articleResponse, nil
}

func UpdateArticle(token string, articleToBeUpdatedRequest *article.Response) (*article.Response, error) {
	headers := header.GetJsonHeader()
	headers["authorization"] = "Token " + token
	updateArticleUrl := fmt.Sprintf("%s/%s", constant.ARTICLE_ENDPOINT, articleToBeUpdatedRequest.Article.Slug)
	response, err := api.NewClient().Put(updateArticleUrl, headers, articleToBeUpdatedRequest)
	if err != nil {
		return nil, err
	}
	var articleResponse *article.Response
	err = json.Unmarshal(response, &articleResponse)
	if err != nil {
		return nil, err
	}
	return articleResponse, nil
}

func DeleteArticle(token string, articleID string) (*article.Response, error) {
	headers := header.GetJsonHeader()
	headers["authorization"] = "Token " + token
	response, err := api.NewClient().Delete(constant.ARTICLE_ENDPOINT+"/"+articleID, headers)
	if err != nil {
		return nil, err
	}
	var articleResponse *article.Response
	err = json.Unmarshal(response, &articleResponse)
	if err != nil {
		return nil, err
	}
	return articleResponse, nil
}

func GetArticle(token string, articleID string) (*article.Response, error) {
	headers := header.GetJsonHeader()
	headers["authorization"] = "Token " + token
	response, err := api.NewClient().Get(constant.ARTICLE_ENDPOINT+"/"+articleID, headers)
	if err != nil {
		return nil, err
	}
	var articleResponse *article.Response
	err = json.Unmarshal(response, articleResponse)
	if err != nil {
		return nil, nil
	}
	return articleResponse, nil
}
