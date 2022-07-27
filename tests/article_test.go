package tests

import (
	"api_automation/actor"
	"api_automation/entity/article"
	"api_automation/entity/user"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestCustomTestSuite(t *testing.T) {
	suite.Run(t, new(CustomTestSuite))
}

func (t *CustomTestSuite) Test_CreateArticle() {
	// Arrange
	u := user.
		NewBuilder().
		SetEmail(faker.Email()).
		SetPassword("123456").
		SetUsername(faker.FirstName()).
		Build()
	articleTitle := faker.Name()
	articleBody := faker.Sentence()
	articleDesc := faker.Sentence()
	a := article.
		NewBuilder().
		SetTagList([]string{"Java", "Python"}).
		SetTitle(articleTitle).
		SetDescription(articleDesc).
		SetBody(articleBody).
		Build()

	// Act
	admin := actor.NewAdmin()
	newUser := admin.CreateUser(u)
	response, err := newUser.CreateArticle(a)

	// Assert
	t.Require().Nil(err, "error in creating article")
	t.Require().Equal(articleTitle, response.Article.Title, "title does not match")
	t.Require().Equal(articleBody, response.Article.Body, "body does not match")
	t.Require().Equal(articleDesc, response.Article.Description, "description does not match")
}

func (t *CustomTestSuite) Test_UpdateArticle() {
	// Arrange
	u := user.
		NewBuilder().
		SetEmail(faker.Email()).
		SetPassword("123456").
		SetUsername(faker.FirstName()).
		Build()
	articleTitle := faker.Name()
	articleBody := faker.Sentence()
	articleDesc := faker.Sentence()
	a := article.
		NewBuilder().
		SetTagList([]string{"Java", "Python"}).
		SetTitle(articleTitle).
		SetDescription(articleDesc).
		SetBody(articleBody).
		Build()
	admin := actor.NewAdmin()
	newUser := admin.CreateUser(u)
	articleToBeUpdated, err := newUser.CreateArticle(a)
	t.Require().Nil(err, "error in creating article")

	// Act
	updatedArticleTitle := faker.Name()
	updatedArticleDesc := faker.Sentence()
	updatedArticleBody := faker.Sentence()
	articleToBeUpdated.Article.Title = updatedArticleTitle
	articleToBeUpdated.Article.Description = updatedArticleDesc
	articleToBeUpdated.Article.Body = updatedArticleBody
	updateArticle, err := newUser.UpdateArticle(articleToBeUpdated)

	// Assert
	t.Require().Nil(err, "error in updating article")
	t.Require().Equal(updatedArticleTitle, updateArticle.Article.Title, "title does not match")
	t.Require().Equal(updatedArticleDesc, updateArticle.Article.Description, "description does not match")
	t.Require().Equal(updatedArticleBody, updateArticle.Article.Body, "body does not match")
}

func (t *CustomTestSuite) Test_DeleteArticle() {
	// Arrange
	u := user.
		NewBuilder().
		SetEmail(faker.Email()).
		SetPassword("123456").
		SetUsername(faker.FirstName()).
		Build()
	articleTitle := faker.Name()
	articleBody := faker.Sentence()
	articleDesc := faker.Sentence()
	a := article.
		NewBuilder().
		SetTagList([]string{"Java", "Python"}).
		SetTitle(articleTitle).
		SetDescription(articleDesc).
		SetBody(articleBody).
		Build()
	admin := actor.NewAdmin()
	newUser := admin.CreateUser(u)
	articleToBeDeleted, err := newUser.CreateArticle(a)
	t.Require().Nil(err, "error in creating article")

	// Act
	deletedArticle, err := newUser.DeleteArticle(articleToBeDeleted.Article.Slug)
	t.Require().Nil(err, "error in deleting article")
	getArticle, err := newUser.GetArticle(deletedArticle.Article.Slug)
	t.Require().Nil(err, "error in getting article")
	t.Require().Equal("", getArticle.Article.Author.Username, "article should be deleted")
}
