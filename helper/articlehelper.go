package helper

import (
	"api_automation/constant"
	"api_automation/entity/article"
	"api_automation/entity/user"
	"api_automation/http"
	"api_automation/testcontext"
	"context"
	"encoding/json"
)

func CreateArticle(ctx context.Context, articleRequest *article.Request) context.Context {
	userResponse := ctx.Value(testcontext.UserResponse{}).(user.Response)
	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	headers["authorization"] = "Token " + userResponse.User.Token
	response := http.PostRequest(constant.ARTICLE_ENDPOINT, headers, articleRequest)
	var articleResponse article.Response
	err := json.Unmarshal(response, &articleResponse)
	if err != nil {
		return nil
	}
	return context.WithValue(ctx, testcontext.ArticleResponse{}, articleResponse)
}

func UpdateArticle(ctx context.Context, articleToBeUpdatedRequest article.Response) context.Context {
	userResponse := ctx.Value(testcontext.UserResponse{}).(user.Response)
	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	headers["authorization"] = "Token " + userResponse.User.Token
	response := http.PutRequest(constant.ARTICLE_ENDPOINT+"/"+articleToBeUpdatedRequest.Article.Slug, headers, articleToBeUpdatedRequest)
	var articleResponse article.Response
	err := json.Unmarshal(response, &articleResponse)
	if err != nil {
		return nil
	}
	return context.WithValue(ctx, testcontext.ArticleResponse{}, articleResponse)
}

func DeleteArticle(ctx context.Context, articleID string) context.Context {
	userResponse := ctx.Value(testcontext.UserResponse{}).(user.Response)
	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	headers["authorization"] = "Token " + userResponse.User.Token
	response := http.DeleteRequest(constant.ARTICLE_ENDPOINT+"/"+articleID, headers)
	return context.WithValue(ctx, testcontext.ArticleResponse{}, response)
}

func GetArticle(ctx context.Context, articleID string) context.Context {
	userResponse := ctx.Value(testcontext.UserResponse{}).(user.Response)
	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	headers["authorization"] = "Token " + userResponse.User.Token
	response := http.GetRequest(constant.ARTICLE_ENDPOINT+"/"+articleID, headers)
	var articleResponse article.Response
	err := json.Unmarshal(response, &articleResponse)
	if err != nil {
		return nil
	}
	return context.WithValue(ctx, testcontext.ArticleResponse{}, articleResponse)
}
