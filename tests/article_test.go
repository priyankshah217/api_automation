package tests

import (
	"api_automation/entity/article"
	"api_automation/entity/user"
	"api_automation/helper"
	"api_automation/testcontext"
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateArticle(t *testing.T) {
	ctx := context.Background()
	userRequest := user.NewRequest(faker.Email(), "123456", faker.FirstName())
	ctx = helper.CreateUser(ctx, userRequest)
	articleTitle := faker.Name()
	articleDesc := faker.Sentence()
	articleBody := faker.Sentence()
	articleRequest := article.NewRequest(article.Article(struct {
		TagList     []string
		Title       string
		Description string
		Body        string
	}{TagList: []string{"Java", "Python"}, Title: articleTitle, Description: articleDesc, Body: articleBody}))
	ctx = helper.CreateArticle(ctx, articleRequest)
	articleResponse := ctx.Value(testcontext.ArticleResponse{}).(article.Response)
	assert.Equal(t, articleTitle, articleResponse.Article.Title, "title does not match")
	assert.Equal(t, articleBody, articleResponse.Article.Body, "body does not match")
	assert.Equal(t, articleDesc, articleResponse.Article.Description, "description does not match")
}

func Test_UpdateArticle(t *testing.T) {
	ctx := context.Background()
	userRequest := user.NewRequest(faker.Email(), "123456", faker.FirstName())
	ctx = helper.CreateUser(ctx, userRequest)
	articleTitle := faker.Name()
	articleDesc := faker.Sentence()
	articleBody := faker.Sentence()
	articleRequest := article.NewRequest(article.Article(struct {
		TagList     []string
		Title       string
		Description string
		Body        string
	}{TagList: []string{"Java", "Python"}, Title: articleTitle, Description: articleDesc, Body: articleBody}))
	ctx = helper.CreateArticle(ctx, articleRequest)
	articleToBeUpdated := ctx.Value(testcontext.ArticleResponse{}).(article.Response)
	updatedArticleTitle := faker.Name()
	updatedArticleDesc := faker.Sentence()
	updatedArticleBody := faker.Sentence()
	articleToBeUpdated.Article.Title = updatedArticleTitle
	articleToBeUpdated.Article.Description = updatedArticleDesc
	articleToBeUpdated.Article.Body = updatedArticleBody
	ctx = helper.UpdateArticle(ctx, articleToBeUpdated)
	updatedArticleResponse := ctx.Value(testcontext.ArticleResponse{}).(article.Response)
	assert.Equal(t, updatedArticleTitle, updatedArticleResponse.Article.Title, "title was not updated")
	assert.Equal(t, updatedArticleBody, updatedArticleResponse.Article.Body, "body was not updated")
	assert.Equal(t, updatedArticleDesc, updatedArticleResponse.Article.Description, "description was not updated")
}

func Test_DeleteArticle(t *testing.T) {
	ctx := context.Background()
	userRequest := user.NewRequest(faker.Email(), "123456", faker.FirstName())
	ctx = helper.CreateUser(ctx, userRequest)
	articleTitle := faker.Name()
	articleDesc := faker.Sentence()
	articleBody := faker.Sentence()
	articleRequest := article.NewRequest(article.Article(struct {
		TagList     []string
		Title       string
		Description string
		Body        string
	}{TagList: []string{"Java", "Python"}, Title: articleTitle, Description: articleDesc, Body: articleBody}))
	ctx = helper.CreateArticle(ctx, articleRequest)
	articleToBeDeleted := ctx.Value(testcontext.ArticleResponse{}).(article.Response)
	ctx = helper.DeleteArticle(ctx, articleToBeDeleted.Article.Slug)
	ctx = helper.GetArticle(ctx, articleToBeDeleted.Article.Slug)
	articleResponse := ctx.Value(testcontext.ArticleResponse{}).(article.Response)
	fmt.Println(articleResponse.Article.Author)
	assert.Equal(t, "", articleResponse.Article.Author.Username, "article not deleted")
}
