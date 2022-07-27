package helper

import (
	"api_automation/constant"
	"api_automation/entity/article"
	"api_automation/http"
	"encoding/json"
	"fmt"
)

func CreateArticle(token string, articleRequest *article.Request) (*article.Response, error) {
	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	headers["authorization"] = "Token " + token
	response, err := http.PostRequest(constant.ARTICLE_ENDPOINT, headers, articleRequest)
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
	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	headers["authorization"] = "Token " + token
	updateArticleUrl := fmt.Sprintf("%s/%s", constant.ARTICLE_ENDPOINT, articleToBeUpdatedRequest.Article.Slug)
	response, err := http.PutRequest(updateArticleUrl, headers, articleToBeUpdatedRequest)
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
	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	headers["authorization"] = "Token " + token
	response, err := http.DeleteRequest(constant.ARTICLE_ENDPOINT+"/"+articleID, headers)
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
	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	headers["authorization"] = "Token " + token
	response, err := http.GetRequest(constant.ARTICLE_ENDPOINT+"/"+articleID, headers)
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
