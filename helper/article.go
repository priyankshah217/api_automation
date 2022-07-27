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
	headers := header.GetJSONHeader()
	headers["authorization"] = "Token " + token
	response, err := api.NewClient().Post(constant.ArticleEndpoint, headers, articleRequest)
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
	headers := header.GetJSONHeader()
	headers["authorization"] = "Token " + token
	updateArticleURL := fmt.Sprintf("%s/%s", constant.ArticleEndpoint, articleToBeUpdatedRequest.Article.Slug)
	response, err := api.NewClient().Put(updateArticleURL, headers, articleToBeUpdatedRequest)
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
	headers := header.GetJSONHeader()
	headers["authorization"] = "Token " + token
	response, err := api.NewClient().Delete(constant.ArticleEndpoint+"/"+articleID, headers)
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
	headers := header.GetJSONHeader()
	headers["authorization"] = "Token " + token
	response, err := api.NewClient().Get(constant.ArticleEndpoint+"/"+articleID, headers)
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
