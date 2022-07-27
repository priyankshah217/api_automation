package tests

import (
	"api_automation/actor"
	"api_automation/entity/article"
	"api_automation/entity/user"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

func TestCustomTestSuite(t *testing.T) {
	suite.Run(t, new(CustomTestSuite))
}

var (
	a *article.Article
	u *user.User
)

func (t *CustomTestSuite) BeforeTest(suiteName string, testName string) {
	log.Println("beforeTest: " + suiteName + ":::" + testName)
	u = user.
		NewBuilder().
		SetEmail(faker.Email()).
		SetPassword("123456").
		SetUsername(faker.FirstName()).
		Build()
	log.Printf("user: %v \n", u)

	a = article.
		NewBuilder().
		SetTagList([]string{"Java", "Python"}).
		SetTitle(faker.Name()).
		SetDescription(faker.Sentence()).
		SetBody(faker.Sentence()).
		Build()
	log.Printf("article: %v \n", a)
}

func (t *CustomTestSuite) AfterTest(suiteName string, testName string) {
	log.Println("afterTest: " + suiteName + ":::" + testName)
}

func (t *CustomTestSuite) Test_CreateArticle() {
	// Arrange
	admin := actor.NewAdmin()
	newUser, err := admin.CreateUser(u)
	t.Require().Nil(err, "error in creating user")

	// Act
	response, err := newUser.CreateArticle(a)

	// Assert
	t.Require().Nil(err, "error in creating article")
	t.Require().Equal(a.Title, response.Article.Title, "title does not match")
	t.Require().Equal(a.Body, response.Article.Body, "body does not match")
	t.Require().Equal(a.Description, response.Article.Description, "description does not match")
}

func (t *CustomTestSuite) Test_UpdateArticle() {
	// Arrange
	admin := actor.NewAdmin()
	newUser, err := admin.CreateUser(u)
	t.Require().Nil(err, "error in creating user")
	articleToBeUpdated, err := newUser.CreateArticle(a)
	t.Require().Nil(err, "error in creating article")

	// Act
	articleToBeUpdated.Article.Title = faker.Name()
	articleToBeUpdated.Article.Description = faker.Sentence()
	articleToBeUpdated.Article.Body = faker.Sentence()
	updateArticle, err := newUser.UpdateArticle(articleToBeUpdated)

	// Assert
	t.Require().Nil(err, "error in updating article")
	t.Require().NotEqual(a.Title, updateArticle.Article.Title, "title does not match")
	t.Require().NotEqual(a.Description, updateArticle.Article.Description, "description does not match")
	t.Require().NotEqual(a.Body, updateArticle.Article.Body, "body does not match")
}

func (t *CustomTestSuite) Test_DeleteArticle() {
	// Arrange
	admin := actor.NewAdmin()
	newUser, err := admin.CreateUser(u)
	t.Require().Nil(err, "error in creating user")
	articleToBeDeleted, err := newUser.CreateArticle(a)
	t.Require().Nil(err, "error in creating article")

	// Act
	deletedArticle, err := newUser.DeleteArticle(articleToBeDeleted.Article.Slug)

	// Assert
	t.Require().Nil(err, "error in deleting article")
	getArticle, err := newUser.GetArticle(deletedArticle.Article.Slug)
	t.Require().Nil(err, "error in getting article")
	t.Require().Equal("", getArticle.Article.Author.Username, "article should be deleted")
}
